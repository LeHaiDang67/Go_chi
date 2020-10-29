package router

import (
	"database/sql"
	"encoding/json"
	"go_chi/internal/feature/transaction"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func getTransaction(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))

		if err != nil {
			json.NewEncoder(w).Encode("Missing user_id parameter")
			return
		}
		accountID, err := strconv.Atoi(r.URL.Query().Get("account_id"))
		if err != nil {
			json.NewEncoder(w).Encode("Missing account_id parameter")
			return
		}

		transactions, err := transaction.GetTransaction(db, userID, accountID)
		if err != nil {
			json.NewEncoder(w).Encode("Cannot get transaction")
			return
		}
		json.NewEncoder(w).Encode(transactions)
	})
}

func addTransaction(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
		if err != nil {
			json.NewEncoder(w).Encode("Missing user_id parameter")
			return
		}
		var requestTran transaction.Transaction
		json.NewDecoder(r.Body).Decode(&requestTran)
		errReq := transaction.AddTransaction(db, &requestTran, userID)
		if errReq != nil {
			json.NewEncoder(w).Encode("Cannot create a new transantion")
			return
		}
	})
}
func updateTransaction(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
		if err != nil {
			json.NewEncoder(w).Encode("Missing user_id parameter")
			return
		}

		id, err2 := strconv.Atoi(r.FormValue("id"))
		if err2 != nil {
			json.NewEncoder(w).Encode("Missing id parameter")
			return
		}
		var requestTran transaction.Transaction
		json.NewDecoder(r.Body).Decode(&requestTran)

		errReq := transaction.UpdateTrasaction(db, &requestTran, userID, id)
		if errReq != nil {
			json.NewEncoder(w).Encode("Cannot update transaction")
			return
		}
	})
}

func deleteTransaction(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err1 := strconv.Atoi(chi.URLParam(r, "user_id"))
		if err1 != nil {
			json.NewEncoder(w).Encode("Missing user_id parameter")
			return
		}
		accountID, err2 := strconv.Atoi(r.FormValue("account_id"))
		if err2 != nil {
			json.NewEncoder(w).Encode("Missing account_id parameter")
			return
		}
		errReq := transaction.DeleteTransaction(db, userID, accountID)
		if errReq != nil {
			json.NewEncoder(w).Encode("Cannot delete transaction ")
			return
		}
	})
}
