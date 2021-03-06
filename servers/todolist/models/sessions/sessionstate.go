package sessions

import (
	"info441-final-project/servers/todolist/models/users"
	"time"
)

type SessionState struct {
	Time time.Time   `json:"time,omitempty"`
	User *users.User `json:"user,omitempty"`
}

// NewSessionState returns a session state including the authenticated user and user's todo list
func NewSessionState(user *users.User) *SessionState {
	return &SessionState{
		Time: time.Now(),
		User: user,
	}
}

// NewTemporarySessionState returns a session state with User value set to nil and an empty todo list
func NewTemporarySessionState() *SessionState {
	return &SessionState{
		Time: time.Now(),
	}
}
