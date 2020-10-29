package transaction

import (
	"database/sql"
	"log"
	"time"
)

//Transaction is an object representing a transaction
type Transaction struct {
	ID              int64   `json:"id"`
	UserID          int64   `json:"user_id"`
	AccountID       int64   `json:"account_id"`
	Amount          float64 `json:"amount"`
	Bank            string  `json:"bank"`
	TransactionType string  `json:"transaction_type"`
	CreatedAt       string  `json:"created_at"`
}

//GetTransaction get a transaction by userID, accountID
func GetTransaction(db *sql.DB, userID, accountID int) ([]Transaction, error) {

	transactions := []Transaction{}
	var rs *sql.Rows
	var err error
	if accountID == -1 {
		rs, err = db.Query("SELECT * FROM transactions where user_id = $1 order by id", userID)
	} else {
		rs, err = db.Query("SELECT * FROM transactions where user_id = $1 and account_id = $2 order by id", userID, accountID)
	}
	if err != nil {
		return transactions, err
	}

	for rs.Next() {
		transaction := Transaction{}
		err := rs.Scan(&transaction.ID, &transaction.UserID, &transaction.AccountID, &transaction.Amount, &transaction.Bank, &transaction.TransactionType, &transaction.CreatedAt)

		if err != nil {
			log.Println("[InsertTransaction] scan error: ", err.Error())
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

//AddTransaction create a new transaction
func AddTransaction(db *sql.DB, transaction *Transaction, userID int) error {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, transaction.CreatedAt)
	result, err := db.Exec("INSERT INTO transactions (user_id,account_id,amount,bank,transaction_type,created_at) VALUES ($1,$2,$3,$4,$5,$6)",
		userID, transaction.AccountID, transaction.Amount, transaction.Bank, transaction.TransactionType, t.Format("2006-01-02T15:04:05-0700"))
	if err != nil {
		return err
	}
	result.RowsAffected()
	return nil
}

//DeleteTransaction delete a transaction
func DeleteTransaction(db *sql.DB, userID, accountID int) error {
	_, err := db.Exec("DELETE FROM transactions WHERE user_id = $1 AND account_id = $2", userID, accountID)
	if err != nil {
		return err
	}
	return nil
}

//UpdateTrasaction edit a transaction by id
func UpdateTrasaction(db *sql.DB, transaction *Transaction, userID, id int) error {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, transaction.CreatedAt)
	result, err := db.Exec("UPDATE transactions SET user_id=$1 ,account_id=$2 , amount=$3 , bank=$4 , transaction_type=$5 , created_at=$6 where id=$7",
		userID, transaction.AccountID, transaction.Amount, transaction.Bank, transaction.TransactionType, t.Format("2006-01-02T15:04:05-0700"), id)
	if err != nil {
		return err
	}
	result.RowsAffected()
	return nil

}
