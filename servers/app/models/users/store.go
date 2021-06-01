package users

import (
	"errors"
)

var ErrUserAlreadyExisted = errors.New("user already existed")
var ErrUserNotFound = errors.New("user not found")

// Store represents a store for Users
type Store interface {
	// GetByID returns the User with the given id
	GetByID(id int64) (*User, error)

	// GetByUsername returns the User with the given username
	GetByUsername(username string) (*User, error)

	// Insert inserts the user into the database, and returns
	// a newly-inserted User, complete with the DBMS-assigned ID
	Insert(user *User) (*User, error)

	// Update applies updates to the user in the database,
	// and returns a newly-updated User
	Update(id int64, updates *Updates) (*User, error)

	// Delete deletes the user from the database
	Delete(id int64) error
}
