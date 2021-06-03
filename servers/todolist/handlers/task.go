package handlers

import (
	"encoding/json"
	"errors"
	"info441-final-project/servers/todolist/models/sessions"
	"info441-final-project/servers/todolist/models/tasks"
	"info441-final-project/servers/todolist/models/users"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

const SuccessDelete = "Delete is successful"

var ErrTodoListNotFound = errors.New("todo list not found")

func (ctx *HandlerContext) TasksHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, currentSession *sessions.SessionState) {
	// Handle request
	switch r.Method {
	case http.MethodGet:
		var todoList []*tasks.Task
		if currentSession.User != nil { // Get from user/task store
			requestedTodoList, err := ctx.TaskStore.GetByUserID(currentSession.User.ID)
			if err != nil {
				if err == tasks.ErrTaskNotFound {
					http.Error(w, err.Error(), http.StatusNotFound)
				} else {
					http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
				}
				return
			}
			todoList = requestedTodoList
		} else { // Get from session
			todoList = currentSession.TodoList
		}
		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(todoList); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		// Validate request
		if !strings.HasPrefix(r.Header.Get(ContentTypeHeader), ContentTypeJSON) {
			http.Error(w, ErrContentTypeNotJSON.Error(), http.StatusUnsupportedMediaType)
			return
		}

		// Decode request body
		newTask := &tasks.Task{}
		if err := json.NewDecoder(r.Body).Decode(newTask); err != nil {
			http.Error(w, ErrInvalidBody.Error(), http.StatusBadRequest)
			return
		}

		if currentSession.User != nil { // Create to user/task store
			// Add the newly created task to store
			insertedTask, err := ctx.TaskStore.Insert(newTask)
			if err != nil {
				if err == tasks.ErrTaskAlreadyExisted {
					http.Error(w, err.Error(), http.StatusBadRequest)
				} else {
					http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
				}
				return
			}
			newTask = insertedTask
		} else { // Create to session
			// Update current session's todo list
			currentSession.TodoList = append(currentSession.TodoList, newTask)
			if err := ctx.SessionStore.Save(sessionID, currentSession); err != nil {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
				return
			}
		}
		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newTask); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}

func (ctx *HandlerContext) SpecificTaskHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, currentSession *sessions.SessionState) {
	// Ensure user is logged in
	if currentSession.User != nil {
		http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
		return
	}

	// Extract id from path
	requestedTaskID, err := strconv.ParseInt(mux.Vars(r)["taskID"], 10, 64)
	if err != nil {
		http.Error(w, ErrInvalidResourcePath.Error(), http.StatusBadRequest)
		return
	}

	// Handle request
	switch r.Method {
	case http.MethodGet:
		var requestedTask *tasks.Task
		if currentSession.User != nil { // Get task from user/task store
			requestedTask, err = ctx.TaskStore.GetByID(requestedTaskID)
			if err != nil {
				if err == tasks.ErrTaskNotFound {
					http.Error(w, err.Error(), http.StatusNotFound)
				} else {
					http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
				}
				return
			}
		} else { // Get task from session
			requestedTask, err = searchTask(requestedTaskID, currentSession.TodoList)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
		}
		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		if err := json.NewEncoder(w).Encode(requestedTask); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPatch:
		// Validate request
		if !strings.HasPrefix(r.Header.Get(ContentTypeHeader), ContentTypeJSON) {
			http.Error(w, ErrContentTypeNotJSON.Error(), http.StatusUnsupportedMediaType)
			return
		}

		// Decode request body
		requestedUpdates := &tasks.Updates{}
		if err := json.NewDecoder(r.Body).Decode(requestedUpdates); err != nil {
			http.Error(w, ErrInvalidBody.Error(), http.StatusBadRequest)
			return
		}

		var updatedTask *tasks.Task
		if currentSession.User != nil { // Update task from user/task store
			// Apply updates store
			updatedTask, err = ctx.TaskStore.Update(requestedTaskID, requestedUpdates)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Update current session's todo list
			if currentSession.TodoList, err = ctx.TaskStore.GetByUserID(currentSession.User.ID); err != nil {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
				return
			}

			if err := ctx.SessionStore.Save(sessionID, currentSession); err != nil {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
				return
			}
		} else { // Update task from session
			updatedTask, err = searchTask(requestedTaskID, currentSession.TodoList)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			if err := updatedTask.ApplyUpdates(requestedUpdates); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := ctx.SessionStore.Save(sessionID, currentSession); err != nil {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
				return
			}
		}
		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		if currentSession.User != nil { // Delete task from user/task store
			if err := ctx.TaskStore.Delete(requestedTaskID); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
		} else { // Delete task from session
			// Update current session's todo list
			currentSession.TodoList = filterTodoList(requestedTaskID, currentSession.TodoList)
			if err := ctx.SessionStore.Save(sessionID, currentSession); err != nil {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
				return
			}
		}
		// Response to request
		w.Write([]byte(SuccessDelete))
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}

func (ctx *HandlerContext) ImportTasksHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, currentSession *sessions.SessionState) {
	// Handle request
	switch r.Method {
	case http.MethodGet:
		// Extract id from path
		r.Header.Set(sessions.HeaderAuthorization, sessions.SchemeBearer+mux.Vars(r)["sessionID"])
		requestedSession := &sessions.SessionState{}
		_, err := sessions.GetState(r, ctx.SigningKey, ctx.SessionStore, requestedSession)
		if err != nil {
			if err == sessions.ErrStateNotFound {
				http.Error(w, ErrTodoListNotFound.Error(), http.StatusNotFound)
			} else {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			}
		}

		var requestedTodoList []*tasks.Task
		if requestedSession.User != nil { // Get from user/task store
			requestedTodoList, err = ctx.TaskStore.GetByUserID(requestedSession.User.ID)
			if err == users.ErrUserNotFound {
				http.Error(w, ErrTodoListNotFound.Error(), http.StatusNotFound)
			} else {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			}
		} else { // Get from session
			requestedTodoList = requestedSession.TodoList
		}

		// Update requested todo list to current session
		currentSession.TodoList = append(currentSession.TodoList, requestedTodoList...)
		if err := ctx.SessionStore.Save(sessionID, currentSession); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		if err := json.NewEncoder(w).Encode(currentSession.TodoList); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}

func filterTodoList(taskIDToRemove int64, todolist []*tasks.Task) []*tasks.Task {
	newTodoList := []*tasks.Task{}
	for _, task := range todolist {
		if task.ID != taskIDToRemove {
			newTodoList = append(newTodoList, task)
		}
	}
	return newTodoList
}

func searchTask(taskID int64, todolist []*tasks.Task) (*tasks.Task, error) {
	for _, task := range todolist {
		if task.ID == taskID {
			return task, nil
		}
	}
	return nil, tasks.ErrTaskNotFound
}
