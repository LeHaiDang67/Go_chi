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
