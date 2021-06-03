package stats

import (
	"database/sql"
	"info441-final-project/servers/todolist/models/users"
)

type MySQLStore struct {
	Client    *sql.DB
	UserStore users.Store
}

// Get all the tasks the user has created total
func (ms *MySQLStore) GetAllByID(userID int64) ([]*Task, error) {
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
		if err := rows.Scan(&task.ID, &task.Name, &description, &task.IsComplete,
			&task.IsHidden, &task.CreatedAt, &task.EditedAt); err != nil {
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

// Get all the tasks the user added this year
func (ms *MySQLStore) GetAllWithinYear(userID int64) ([]*Task, error) {
	query := "SELECT ID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE UserID = ? AND CreatedAt > DATEADD(year,-1,GETDATE())"
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
		if err := rows.Scan(&task.ID, &task.Name, &description, &task.IsComplete,
			&task.IsHidden, &task.CreatedAt, &task.EditedAt); err != nil {
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

// Get all the tasks the user added this month
func (ms *MySQLStore) GetAllWithinMonth(userID int64) ([]*Task, error) {
	query := "SELECT ID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE UserID = ? AND CreatedAt > DATEADD(month,-1,GETDATE())"
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
		if err := rows.Scan(&task.ID, &task.Name, &description, &task.IsComplete,
			&task.IsHidden, &task.CreatedAt, &task.EditedAt); err != nil {
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

// Get all the tasks the user added this week
func (ms *MySQLStore) GetAllWithinWeek(userID int64) ([]*Task, error) {
	query := "SELECT ID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE UserID = ? AND CreatedAt > DATEADD(week,-1,GETDATE())"
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
		if err := rows.Scan(&task.ID, &task.Name, &description, &task.IsComplete,
			&task.IsHidden, &task.CreatedAt, &task.EditedAt); err != nil {
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

// Get all the tasks the user had between two dates
func (ms *MySQLStore) GetAllBetweenDates(userID int64, beginDate string, endDate string) ([]*Task, error) {
	query := "SELECT ID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE UserID = ? AND CreatedAt between ? and ?"
	rows, err := ms.Client.Query(query, userID, beginDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Extract each task
	tasks := []*Task{}
	for rows.Next() {
		description := sql.NullString{}
		task := &Task{}
		if err := rows.Scan(&task.ID, &task.Name, &description, &task.IsComplete,
			&task.IsHidden, &task.CreatedAt, &task.EditedAt); err != nil {
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

// Get all the tasks the user has created total
func (ms *MySQLStore) GetCompletedByID(userID int64) ([]*Task, error) {
	query := "SELECT ID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE UserID = ? AND IsComplete"
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
		if err := rows.Scan(&task.ID, &task.Name, &description, &task.IsComplete,
			&task.IsHidden, &task.CreatedAt, &task.EditedAt); err != nil {
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

// Get all the tasks the user added this year
func (ms *MySQLStore) GetCompletedWithinYear(userID int64) ([]*Task, error) {
	query := "SELECT ID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE UserID = ? AND CreatedAt > DATEADD(year,-1,GETDATE()) AND IsComplete"
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
		if err := rows.Scan(&task.ID, &task.Name, &description, &task.IsComplete,
			&task.IsHidden, &task.CreatedAt, &task.EditedAt); err != nil {
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

// Get all the tasks the user added this month
func (ms *MySQLStore) GetCompletedWithinMonth(userID int64) ([]*Task, error) {
	query := "SELECT ID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE UserID = ? AND CreatedAt > DATEADD(month,-1,GETDATE()) AND IsComplete"
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
		if err := rows.Scan(&task.ID, &task.Name, &description, &task.IsComplete,
			&task.IsHidden, &task.CreatedAt, &task.EditedAt); err != nil {
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

// Get all the tasks the user added this week
func (ms *MySQLStore) GetCompletedWithinWeek(userID int64) ([]*Task, error) {
	query := "SELECT ID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE UserID = ? AND CreatedAt > DATEADD(week,-1,GETDATE()) AND IsComplete"
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
		if err := rows.Scan(&task.ID, &task.Name, &description, &task.IsComplete,
			&task.IsHidden, &task.CreatedAt, &task.EditedAt); err != nil {
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

// Get all the tasks the user had between two dates
func (ms *MySQLStore) GetCompletedBetweenDates(userID int64, beginDate string, endDate string) ([]*Task, error) {
	query := "SELECT ID, Name, Description, IsComplete, IsHidden, CreatedAt, EditedAt FROM TodoList WHERE UserID = ? AND CreatedAt between ? and ? AND IsComplete"
	rows, err := ms.Client.Query(query, userID, beginDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Extract each task
	tasks := []*Task{}
	for rows.Next() {
		description := sql.NullString{}
		task := &Task{}
		if err := rows.Scan(&task.ID, &task.Name, &description, &task.IsComplete,
			&task.IsHidden, &task.CreatedAt, &task.EditedAt); err != nil {
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
