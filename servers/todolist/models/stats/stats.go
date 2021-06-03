package stats

import "info441-final-project/servers/todolist/models/tasks"

type QueryResults struct {
	Completed []*tasks.Task `json:"completed"`
	Created   []*tasks.Task `json:"created"`
}
