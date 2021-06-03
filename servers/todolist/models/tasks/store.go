package tasks

import (
	"errors"
)

// ErrTaskAlreadyExisted is returned from Store.Insert() when the given task
// already existed in the database
var ErrTaskAlreadyExisted = errors.New("task already existed")

// ErrTaskNotFound is returned when a Store's getter function could not
// find the requested task
var ErrTaskNotFound = errors.New("task not found")

// Store represents a store for Tasks
type Store interface {
	// GetByID returns the Task with the given id
	GetByID(id int64) (*Task, error)

	// GetByUserID returns a user's todo list as an array of tasks
	GetByUserID(userID int64) ([]*Task, error)

	// Insert inserts the task into the database, and returns
	// a newly-inserted Task, complete with the DBMS-assigned ID
	Insert(task *Task) (*Task, error)

	// Update applies updates to the task in the database,
	// and returns a newly-updated Task
	Update(id int64, updates *Updates) (*Task, error)

	// Delete deletes the task from the database
	Delete(id int64) error
}
