package handlers

import (
	"encoding/json"
	"info441-final-project/servers/todolist/models/sessions"
	"info441-final-project/servers/todolist/models/tasks"
	"info441-final-project/servers/todolist/models/users"
	"net/http"
)

type HandlerContext struct {
	SigningKey   string
	SessionStore sessions.Store
	UserStore    users.Store
	TaskStore    tasks.Store
}

const ContentTypeHeader = "Content-Type"
const ContentTypeJSON = "application/json"

const AdminUserID = 1

// Example todo list, the data is hard coded
func (ctx *HandlerContext) TodoList(w http.ResponseWriter, r *http.Request, sessionID sessions.SessionID, currentSession *sessions.SessionState) {
	// Object todoList is an array of object Task
	todoList, err := ctx.TaskStore.GetByUserID(AdminUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentSession.TodoList = todoList

	w.Header().Add(ContentTypeHeader, ContentTypeJSON)
	json.NewEncoder(w).Encode(currentSession.TodoList)
}
