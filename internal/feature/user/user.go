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
	// return User{
	// 	ID:       "1",
	// 	Address:  "VN",
	// 	Birthday: time.Now(),
	// 	Name:     "Join",
	// }, nil

}

// func AddUser(db *sql.DB, user User) (User,error)
// {

// 	err := db.Query("INSERT INTO users (id, address, birthday,name) VALUES ($1,$2,$3,$4)",
// 	+ user.ID,user.Address,user.Birthday,user.Name)

// }
