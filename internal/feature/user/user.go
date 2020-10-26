package user

import (
	"database/sql"
	"time"
)

// User is an object representing the user.
type User struct {
	ID       string    `json:"id"`
	Address  string    `json:"address"`
	Birthday time.Time `json:"birthday"`
	Name     string    `json:"name"`
}

// GetUser get a user, return error if fail
func GetUser(db *sql.DB, id int) (*User, error) {
	// TODO:
	return &User{
		ID:       "1",
		Address:  "VN",
		Birthday: time.Now(),
		Name:     "Join",
	}, nil
}
