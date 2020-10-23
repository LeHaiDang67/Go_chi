package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "test"
)

// type User struct {
// 	ID       int
// 	Address  string
// 	Birthday string
// 	Name     string
// }

type Storage struct {
	db *sql.DB
}

func SetUpStorage() (*Storage, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

// func ConnectDB() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatal("Failed to open a DB connection: ", err)
// 	}
// 	defer db.Close()

// 	// Create an empty user and make the sql query (using $1 for the parameter)
// 	var myUser User
// 	userSql := "SELECT * FROM users WHERE id = $1"

// 	err = db.QueryRow(userSql, 1).Scan(&myUser.ID, &myUser.Address, &myUser.Birthday, &myUser.Name)
// 	if err != nil {
// 		log.Fatal("Failed to execute query: ", err)
// 	}

// 	fmt.Printf("Hi %s, welcome back!\n", myUser.Name)
// }
