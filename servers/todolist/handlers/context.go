package handlers

import (
	"info441-final-project/servers/todolist/models/sessions"
	"info441-final-project/servers/todolist/models/stats"
	"info441-final-project/servers/todolist/models/tasks"
	"info441-final-project/servers/todolist/models/users"
)

type HandlerContext struct {
	UserStore    users.Store
	TaskStore    tasks.Store
	StatStore    stats.Store
	SigningKey   string
	SessionStore *sessions.RedisStore
}

const ContentTypeHeader = "Content-Type"
const ContentTypeJSON = "application/json"
