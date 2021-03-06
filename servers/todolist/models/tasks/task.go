package tasks

import (
	"info441-final-project/servers/todolist/models/users"
	"time"
)

// Task represents a single task of a user's todo list
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

// Updates represents allowed updates to a task
type Updates struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	IsComplete  *bool   `json:"isComplete"`
	IsHidden    *bool   `json:"isHidden"`
}

// ApplyUpdates applies the updates to the task. An error
// is returned if the updates are invalid
func (t *Task) ApplyUpdates(updates *Updates) error {
	if updates.Name != nil {
		t.Name = *updates.Name
	}
	if updates.Description != nil {
		t.Description = *updates.Description
	}
	if updates.IsComplete != nil {
		t.IsComplete = *updates.IsComplete
	}
	if updates.IsHidden != nil {
		t.IsHidden = *updates.IsHidden
	}
	return nil
}
