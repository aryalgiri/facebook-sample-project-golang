package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	"github.com/sample-fb/db"
	"github.com/sample-fb/models"
)

var (
	accounts []models.Account
)

// CreateAccount is used to create an account
func CreateAccount(w http.ResponseWriter, req *http.Request) {
	// decode json data to our data model
	var account models.Account

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err = json.Unmarshal(body, &account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if account.FirstName == "" || account.LastName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("FirstName and Lastname in the body is required"))
		return
	}

	// generate new id for the account

	account.ID = uuid.New().String()
	//Add status of an account
	account.Status="ACTIVE"
	
	//save account to db
	err = db.SaveAccount(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	logger.Infof("account %v created succesfully", account.ID)
	_, err:= db.Exec("Insert into accounts(id, firstname, middlename, lastname, email, phonenumber, password, gender, stats) ")VALUES($1, $2, $2, $3, $4, $5, $6, $7, $8, $9) returning uid:", account.ID, account.FirstName, account.MiddleName, account.Email, account.PhoneNumber, account.Password, account.Gender, account.Status")
}
