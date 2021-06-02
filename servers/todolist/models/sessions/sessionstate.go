package sessions

import (
	"info441-final-project/servers/todolist/models/tasks"
	"info441-final-project/servers/todolist/models/users"
	"time"
)

type SessionState struct {
	Time     time.Time     `json:"time,omitempty"`
	User     *users.User   `json:"user,omitempty"`
	TodoList []*tasks.Task `json:"todoList,omitempty"`
}

// NewSessionState returns a session state including the authenticated user and user's todo list
func NewSessionState(user *users.User, todoList []*tasks.Task) *SessionState {
	return &SessionState{
		Time:     time.Now(),
		User:     user,
		TodoList: todoList,
	}
}

// NewTemporarySessionState returns a session state with User value set to nil and an empty todo list
func NewTemporarySessionState() *SessionState {
	return &SessionState{
		Time:     time.Now(),
		TodoList: []*tasks.Task{},
	}
}
