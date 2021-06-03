package tasks

import (
	"database/sql"
	"info441-final-project/servers/todolist/models/users"
)

type MySQLStore struct {
	Client    *sql.DB
	UserStore users.Store
}

// GetByID returns the Task with the given id
func (ms *MySQLStore) GetByID(id int64) (*Task, error) {
	return ms.selectTaskWhere("ID", id)
}

// GetByUserID returns a user's todo list as an array of tasks
func (ms *MySQLStore) GetByUserID(userID int64) ([]*Task, error) {
	query := "SELECT ID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE UserID = ?"
	rows, err := ms.Client.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Extract each task
	tasks := []*Task{}
	for rows.Next() {
		description := sql.NullString{}
		task := &Task{}
		if err := rows.Scan(
			&task.ID,
			&task.Name,
			&description,
			&task.IsComplete,
			&task.IsHidden,
			&task.CreatedAt,
			&task.EditedAt); err != nil {
			return nil, err
		}
		task.Description = description.String
		tasks = append(tasks, task)
	}

	// Set all tasks to the same user
	user, err := ms.UserStore.GetByID(userID)
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		task.User = user
	}

	return tasks, nil
}

// Insert inserts the task into the database, and returns
// a newly-inserted Task, complete with the DBMS-assigned ID
func (ms *MySQLStore) Insert(task *Task) (*Task, error) {
	query := "INSERT IGNORE INTO TodoList (UserID, Name, Description, IsComplete, IsHidden) VALUES (?, ?, ?, ?, ?)"
	response, err := ms.Client.Exec(
		query,
		task.User.ID,
		task.Name,
		task.Description,
		task.IsComplete,
		task.IsHidden)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, ErrTaskAlreadyExisted
	}

	id, err := response.LastInsertId()
	if err != nil {
		return nil, err
	}
	return ms.GetByID(id) // The db handles the timestamp, so need to retrieve the task from it again.
}

// Update applies updates to the task in the database,
// and returns a newly-updated Task
func (ms *MySQLStore) Update(id int64, updates *Updates) (*Task, error) {
	task, err := ms.GetByID(id)
	if err != nil {
		return nil, err
	}

	if task.ApplyUpdates(updates); err != nil {
		return nil, err
	}

	query := "UPDATE TodoList SET Name = ?, Description = ?, IsComplete = ?, IsHidden = ? WHERE ID = ?"
	if _, err := ms.Client.Exec(
		query,
		task.Name,
		task.Description,
		task.IsComplete,
		task.IsHidden,
		task.ID); err != nil {
		return nil, err
	}
	return task, nil
}

// Delete deletes the task from the database
func (ms *MySQLStore) Delete(id int64) error {
	query := "DELETE FROM TodoList WHERE ID = ?"
	_, err := ms.Client.Exec(query, id)
	return err
}

// selectUserWhere executes the basic select from where sql statement, and returns the first row as object model
func (ms *MySQLStore) selectTaskWhere(property string, value interface{}) (*Task, error) {
	var userID int64
	description := sql.NullString{}

	query := "SELECT ID, UserID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE " + property + " = ?"
	row := ms.Client.QueryRow(query, value)

	task := &Task{}
	if err := row.Scan(
		&task.ID,
		&userID,
		&task.Name,
		&description,
		&task.IsComplete,
		&task.IsHidden,
		&task.CreatedAt,
		&task.EditedAt); err != nil {
		return nil, ErrTaskNotFound
	}
	task.Description = description.String

	user, err := ms.UserStore.GetByID(userID)
	if err != nil {
		return nil, err
	}
	task.User = user

	return task, nil
}
