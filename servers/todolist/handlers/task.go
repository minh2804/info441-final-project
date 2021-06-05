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
		// Search for the requested object
		requestedTodoList, err := ctx.TaskStore.GetByUserID(currentSession.User.ID)
		if err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		if err := json.NewEncoder(w).Encode(requestedTodoList); err != nil {
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

		// Add the newly created task to store
		newTask.User = currentSession.User
		registeredTask, err := ctx.TaskStore.Insert(newTask)
		if err != nil {
			if err == tasks.ErrTaskAlreadyExisted {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			}
			return
		}

		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(registeredTask); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}

func (ctx *HandlerContext) SpecificTaskHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, currentSession *sessions.SessionState) {
	// Extract id from path
	requestedTaskID, err := strconv.ParseInt(mux.Vars(r)["taskID"], 10, 64)
	if err != nil {
		http.Error(w, tasks.ErrTaskNotFound.Error(), http.StatusNotFound)
		return
	}

	// Search for the requested object
	requestedTask, err := ctx.TaskStore.GetByID(requestedTaskID)
	if err != nil {
		if err == tasks.ErrTaskNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Handle request
	switch r.Method {
	case http.MethodGet:
		// Only the onwer of the task can view hidden task
		if requestedTask.IsHidden && (requestedTask.User.ID != currentSession.User.ID) {
			http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
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

		// Only the onwer of the task can update it
		if requestedTask.User.ID != currentSession.User.ID {
			http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		// Decode request body
		requestedUpdates := &tasks.Updates{}
		if err := json.NewDecoder(r.Body).Decode(requestedUpdates); err != nil {
			http.Error(w, ErrInvalidBody.Error(), http.StatusBadRequest)
			return
		}

		// Apply updates store
		updatedTask, err := ctx.TaskStore.Update(requestedTaskID, requestedUpdates)
		if err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		// Only the onwer of the task can delete it
		if requestedTask.User.ID != currentSession.User.ID {
			http.Error(w, ErrUnauthorized.Error(), http.StatusUnauthorized)
			return
		}

		// Delete the task
		if err := ctx.TaskStore.Delete(requestedTaskID); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
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
		requestedUserID, err := strconv.ParseInt(mux.Vars(r)["userID"], 10, 64)
		if err != nil {
			http.Error(w, users.ErrUserNotFound.Error(), http.StatusNotFound)
			return
		}

		// Get the todo list to import
		requestedTodoList, err := ctx.TaskStore.GetByUserID(requestedUserID)
		if err != nil {
			if err == tasks.ErrTaskNotFound {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			}
			return
		}

		// Append the new list to current list
		for _, task := range requestedTodoList {
			if !task.IsHidden {
				task.User = currentSession.User
				if _, err := ctx.TaskStore.Insert(task); err != nil {
					http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
					return
				}
			}
		}

		updatedTodoList, err := ctx.TaskStore.GetByUserID(currentSession.User.ID)
		if err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}

		// Response to request
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(updatedTodoList); err != nil {
			http.Error(w, ErrInternal.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, ErrRequestMethodNotAllowed.Error(), http.StatusMethodNotAllowed)
	}
}
