package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/logger"
	"github.com/google/uuid"

	"github.com/sample-fb/models"
)

// createAccount is used to create an account
func createAccount(w http.ResponseWriter, req *http.Request) {
	// decode json data to our data model
	var account models.Account

	// read the request body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// decode the request body to account variable
	if err = json.Unmarshal(body, &account); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// validation check
	if account.Firstname == "" || account.Lastname == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("FirstName and Lastname in the body is required"))
		return
	}

	// generate new id for the account
	account.ID = uuid.New().String()
	// Add status of an account
	account.Status = models.StatusActive

	// save account to DB
	if err = dbmgr.SaveAccount(&account); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	logger.Infof("account with ID: %v created successfully", account.ID)
	w.WriteHeader(http.StatusCreated)
}
