package db

import (
	"fmt"
	"net/http"

	"github.com/sample-fb/models"
)

//ListAccounts is used to get the list of account details
func ListAccounts(w http.ResponseWriter, r *http.Request) {
	var id string
	value, ok := r.URL.Query()["id"]
	if ok {
		id = value[0]
	}
	accs, err := db.GetAccount(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	//Encote the response to json
	JSONResp
}

//GetAccount is used to retrieve account details from db
func GetAccount(id string) ([]models.Account, error) {
	logger.Info("retrieving account details from db")
	sqlQuery := "SELECT * from accounts"
	if id != "" {
		sqlQuery = fmt.Sprintf("%v WHERE id='%v'", sqlQuery, id)
	}
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	var accounts []models.Account
	for rows.Next() {
		var account models.Account
		err = rows.Scan(&account.ID, &account.FirstName, &account.MiddleName, &account.LastName, &account.Email, &account.PhoneNumber, &account.Password, &account.Gender, &account.Address, &account.Status)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
