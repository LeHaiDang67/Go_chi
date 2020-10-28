package router

import (
	"database/sql"
	"encoding/json"
	"go_chi/internal/feature/user"
	"net/http"
	"strconv"
)

func getUser(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			json.NewEncoder(w).Encode("Missing id parameter")
			return
		}
		user, err := user.GetUser(db, id)
		if err != nil {
			json.NewEncoder(w).Encode("Cannot fetch user")
			return
		}
		json.NewEncoder(w).Encode(user)
	})
}

func addUser(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestUser user.User
		json.NewDecoder(r.Body).Decode(&requestUser)
		errReq := user.AddUser(db, &requestUser)
		if errReq != nil {
			json.NewEncoder(w).Encode("Cannot insert user")
			return
		}

	})
}

func deleteUser(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			json.NewEncoder(w).Encode("Missing id parameter")
			return
		}
		errReq := user.DeleteUser(db, id)
		if errReq != nil {
			json.NewEncoder(w).Encode("Cannot delete user")
			return
		}

	})
}

func updateUser(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			json.NewEncoder(w).Encode("Missing id parameter")
			return
		}
		var requestUser user.User
		json.NewDecoder(r.Body).Decode(&requestUser)
		errRep := user.UpdateUser(db, &requestUser, id)
		if errRep != nil {
			json.NewEncoder(w).Encode("Cannot update user")
			return
		}
	})
}
