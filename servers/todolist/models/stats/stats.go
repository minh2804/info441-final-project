package stats

import (
	"info441-final-project/servers/todolist/models/users"
	"time"
)

type Task struct {
	ID          int64       `json:"id"`
	User        *users.User `json:"user"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	IsComplete  bool        `json:"isComplete"`
	IsHidden    bool        `json:"isHidden"`
	CreatedAt   time.Time   `json:"createdAt"`
	EditedAt    time.Time   `json:"editedAt"`
}

type QueryResults struct {
	Completed *[]*Task `json: completed`
	Created   *[]*Task `json: created`
}
