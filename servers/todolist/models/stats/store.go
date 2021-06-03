package stats

import (
	"errors"
	"info441-final-project/servers/todolist/models/tasks"
)

var ErrTaskNotFound = errors.New("task not found")

// Store represents a store for Stats
type Store interface {
	// Get all the tasks the user has created total
	GetAllByID(userID int64) ([]*tasks.Task, error)

	// Get all the tasks the user added this year
	GetAllWithinYear(userID int64) ([]*tasks.Task, error)

	// Get all the tasks the user added this month
	GetAllWithinMonth(userID int64) ([]*tasks.Task, error)

	// Get all the tasks the user added this week
	GetAllWithinWeek(userID int64) ([]*tasks.Task, error)

	// Get all the tasks the user had between two dates
	GetAllBetweenDates(userID int64, beginDate string, endDate string) ([]*tasks.Task, error)

	// Get all the tasks the user has completed total
	GetCompletedByID(userID int64) ([]*tasks.Task, error)

	// Get all the tasks the user completed this year
	GetCompletedWithinYear(userID int64) ([]*tasks.Task, error)

	// Get all the tasks the user completed this month
	GetCompletedWithinMonth(userID int64) ([]*tasks.Task, error)

	// Get all the tasks the user completed this week
	GetCompletedWithinWeek(userID int64) ([]*tasks.Task, error)

	// Get all the tasks the user completed between two dates
	GetCompletedBetweenDates(userID int64, beginDate string, endDate string) ([]*tasks.Task, error)
}
