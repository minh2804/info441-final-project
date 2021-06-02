package users

import (
	"database/sql"
)

type MySQLStore struct {
	Client *sql.DB
}

// GetByID returns the User with the given id
func (ms *MySQLStore) GetByID(id int64) (*User, error) {
	return ms.selectUserWhere("ID", id)
}

// GetByUsername returns the User with the given username
func (ms *MySQLStore) GetByUsername(username string) (*User, error) {
	return ms.selectUserWhere("Username", username)
}

// Insert inserts the user into the database, and returns
// a newly-inserted User, complete with the DBMS-assigned ID
func (ms *MySQLStore) Insert(user *User) (*User, error) {
	query := "INSERT IGNORE INTO User (Username, PassHash, FirstName, LastName) VALUES (?, ?, ?, ?)"
	response, err := ms.Client.Exec(
		query,
		user.Username,
		user.PassHash,
		user.FirstName,
		user.LastName)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := response.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, ErrUserAlreadyExisted
	}

	id, err := response.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = id

	return user, nil
}

// Update applies updates to the user in the database,
// and returns a newly-updated User
func (ms *MySQLStore) Update(id int64, updates *Updates) (*User, error) {
	user, err := ms.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user.ApplyUpdates(updates); err != nil {
		return nil, err
	}

	query := "UPDATE User SET FirstName = ?, LastName = ? WHERE ID = ?"
	if _, err := ms.Client.Exec(
		query,
		user.FirstName,
		user.LastName,
		user.ID); err != nil {
		return nil, err
	}
	return user, nil
}

// Delete deletes the user from the database
func (ms *MySQLStore) Delete(id int64) error {
	query := "CALL DeleteChannel(?)"
	_, err := ms.Client.Exec(query, id)
	return err
}

// selectUserWhere executes the basic "select from where" sql statement, and returns the first row as object model
func (ms *MySQLStore) selectUserWhere(property string, value interface{}) (*User, error) {
	firstName := sql.NullString{}
	lastName := sql.NullString{}

	query := "SELECT ID, Username, PassHash, FirstName, LastName FROM User WHERE " + property + " = ?"
	row := ms.Client.QueryRow(query, value)
	user := &User{}
	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.PassHash,
		&firstName,
		&lastName); err != nil {
		return nil, ErrUserNotFound
	}

	user.FirstName = firstName.String
	user.LastName = lastName.String

	return user, nil
}
