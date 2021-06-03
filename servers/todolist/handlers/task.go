package handlers

import (
	"database/sql"
	"encoding/json"
	"info441-final-project/servers/todolist/models/sessions"
	"info441-final-project/servers/todolist/models/tasks"
	"io/ioutil"
	"net/http"
	"path"

	"strconv"

	"time"
)

type TaskItem struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	TaskName    string    `json:"taskName"`
	Description string    `json:"description"`
	IsComplete  bool      `json:"IsComplete"`
	IsHidden    bool      `json:"IsHidden"`
	CreatedAt   time.Time `json:"CreatedAt"`
	EditedAt    time.Time `json:"EditedAt"`
}

type MysqlStore struct {
	DB *sql.DB
}

func (sql *MysqlStore) TaskPostHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, sessionState *sessions.SessionState) {
	if sessionState.User == nil {
		hc := &HandlerContext{}
		hc.TodoList(w, r, sessionID, sessionState)
		task := &tasks.Task{}
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sessionState.TodoList[len(sessionState.TodoList)] = task
		w.Header().Add(ContentTypeHeader, ContentTypeJSON)
		json.NewEncoder(w).Encode(sessionState.TodoList)
	} else {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		} else {
			task := &TaskItem{}
			err := json.NewDecoder(r.Body).Decode(&task)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			insq := "INSERT INTO TodoList (UserID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt) VALUES (?, ?, ?, ?, ?, ?, ?)"
			res, err := sql.DB.Exec(insq, task.UserID, task.TaskName, task.Description, task.IsComplete,
				task.IsHidden, time.Now(), time.Now())
			if err != nil {
				w.Write([]byte(err.Error()))
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			} else {
				id, err := res.LastInsertId()
				if err != nil {
					w.Write([]byte(err.Error()))
					http.Error(w, err.Error(), http.StatusForbidden)
					return
				} else {
					task.ID = id
				}
			}

			encodedTask, err := json.Marshal(task)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			w.Write(encodedTask)
			w.Header().Set("Content-Type", "application/json")

			file, _ := json.MarshalIndent(task, "", " ")
			_ = ioutil.WriteFile("todoList.json", file, 0644)
		}
	}
}

func (sql *MysqlStore) TaskPatchHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, sessionState *sessions.SessionState) {
	if sessionState.User == nil {
		w.Write([]byte("Update is unsuccessful. User not logged in"))
	} else {
		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		} else {
			taskId := path.Base(r.URL.Path)
			task := &TaskItem{}
			err := json.NewDecoder(r.Body).Decode(&task)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			insq := "Update TodoList SET UserId=?, Name=?, Description=?, IsComplete=?,EditedAt=? WHERE TaskId=?"
			res, err := sql.DB.Exec(insq, task.UserID, task.TaskName, task.Description, task.IsComplete,
				time.Now(), taskId)
			if err != nil {
				w.Write([]byte(err.Error()))
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			} else {
				id, err := res.LastInsertId()
				if err != nil {
					w.Write([]byte(err.Error()))
					http.Error(w, err.Error(), http.StatusForbidden)
					return
				} else {
					task.ID = id
				}
			}

			encodedTask, err := json.Marshal(task)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			w.Write(encodedTask)
			w.Header().Set("Content-Type", "application/json")

			file, _ := json.MarshalIndent(task, "", " ")
			_ = ioutil.WriteFile("todoList.json", file, 0644)
		}
	}
}

func (sql *MysqlStore) GetTaskHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, sessionState *sessions.SessionState) {
	if sessionState.User == nil {
		w.Write([]byte("Get is unsuccessful. User not logged in"))
	} else {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		} else {
			itemID := path.Base(r.URL.Path)
			userID, err := strconv.ParseInt(itemID, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			sqlStatement := "select UserId, Name, Description, IsComplete,IsHidden, CreatedAt, EditedAt from TodoList where TaskId=?"
			rows, err := sql.DB.Query(sqlStatement, userID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			task := &TaskItem{}

			for rows.Next() {
				err = rows.Scan(task.UserID, task.TaskName, task.Description, task.IsComplete,
					task.IsHidden, task.CreatedAt, task.EditedAt)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
			}

			encodedTask, err := json.Marshal(task)
			if err != nil {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			w.Write(encodedTask)
			w.Header().Set("Content-Type", "application/json")
			file, _ := json.MarshalIndent(task, "", " ")
			_ = ioutil.WriteFile("todoList.json", file, 0644)
		}
	}
}

func (sql *MysqlStore) DeleteTaskHandler(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, sessionState *sessions.SessionState) {
	if sessionState.User == nil {
		w.Write([]byte("Delete is unsuccessful. User not logged in"))
	} else {
		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		} else {
			itemID := path.Base(r.URL.Path)
			userID, err := strconv.ParseInt(itemID, 10, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			sqlStatement := "DELETE FROM task WHERE TodoList = ?;"
			_, err = sql.DB.Query(sqlStatement, userID)
			if err != nil {
				w.Write([]byte(err.Error()))
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			w.Write([]byte("Delete is successful"))

		}
	}
}
