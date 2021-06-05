package users

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// bcryptCost is the default bcrypt cost to use when hashing passwords
const bcryptCost = 13

var ErrInvalidPasswordLength = errors.New("password must be at least 6 characters")
var ErrInvalidUsername = errors.New("username must be non-zero length and may not contain spaces")
var ErrMismatchPassword = errors.New("password and passwordConf does not match")

// User represents a user account in the database
type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	PassHash    []byte `json:"-"` // never JSON encoded/decoded
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	IsTemporary bool   `json:"IsTemporary"`
}

// Credentials represents user sign-in credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Updates represents allowed updates to a user profile
type Updates struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

// NewUser represents a new user signing up for an account
type NewUser struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	IsTemporary  bool   `json:"IsTemporary"`
}

// Validate validates the new user and returns an error if
// any of the validation rules fail, or nil if its valid
func (nu *NewUser) Validate() error {
	if len(nu.Password) < 6 {
		return ErrInvalidPasswordLength
	}
	if nu.Password != nu.PasswordConf {
		return ErrMismatchPassword
	}
	if len(nu.Username) == 0 || strings.Contains(nu.Username, " ") {
		return ErrInvalidUsername
	}
	return nil
}

// ToUser converts the NewUser to a User, and set the PassHash field appropriately
func (nu *NewUser) ToUser() (*User, error) {
	if err := nu.Validate(); err != nil {
		return nil, err
	}
	user := &User{
		Username:    nu.Username,
		FirstName:   nu.FirstName,
		LastName:    nu.LastName,
		IsTemporary: nu.IsTemporary,
	}
	user.SetPassword(nu.Password)
	return user, nil
}

// SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	if len(password) < 6 {
		return ErrInvalidPasswordLength
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	u.PassHash = passHash
	return err
}

// Authenticate compares the plaintext password against the stored hash
// and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
}

// FullName returns the user's full name, in the form:
// "<FirstName> <LastName>"
// If either first or last name is an empty string, no
// space is put between the names. If both are missing,
// this returns an empty string
func (u *User) FullName() string {
	return strings.TrimSpace(u.FirstName + " " + u.LastName)
}

// ApplyUpdates applies the updates to the user. An error
// is returned if the updates are invalid
func (u *User) ApplyUpdates(updates *Updates) error {
	if updates.FirstName != nil {
		u.FirstName = *updates.FirstName
	}
	if updates.LastName != nil {
		u.LastName = *updates.LastName
	}
	return nil
}
