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
func GetUser(db *sql.DB, id int) (User, error) {
	// TODO:
	user := User{}
	r, err := db.Query("SELECT * From users where id = $1", id)
	if err != nil {
		return user, err
	}
	for r.Next() {
		err = r.Scan(&user.ID, &user.Address, &user.Birthday, &user.Name)
		if err != nil {
			panic(err)
		}
	}
	return user, nil

}

//AddUser insert new user
func AddUser(db *sql.DB, user *User) error {

	result, err := db.Exec("INSERT INTO users (address, birthday, name) VALUES ($1 , $2, $3)", user.Address, user.Birthday, user.Name)
	if err != nil {

		return err
	}
	result.RowsAffected()
	return nil
}

//DeleteUser delete user
func DeleteUser(db *sql.DB, id int) error {

	_, err := db.Exec("DELETE FROM users where id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

//UpdateUser edit user
func UpdateUser(db *sql.DB, user *User, id int) error {
	result, err := db.Exec("UPDATE users SET address = $1, birthday=$2, name = $3  where id = $4", user.Address, user.Birthday, user.Name, id)

	if err != nil {
		return err
	}
	result.RowsAffected()
	return nil
}
