package Users

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getHandler(s Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			json.NewEncoder(w).Encode("Something went wrong... ")
			return
		}
		u, err := s.ReadUser(id)
		if err != nil {
			json.NewEncoder(w).Encode("Something went wrong... ")
			return
		}
		json.NewEncoder(w).Encode(&u)
	}
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("You sent me a post req")
}
