package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sample-fb/models"
)

// loginHandler is used to login to the account
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var acc models.Account

	// read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// decode the body to acc variable
	if err = json.Unmarshal(body, &acc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// check if email and password exists in request body
	if acc.Email == "" || acc.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email and Password are required to login, please check the values that your entered"))
		return
	}

	// Get Account details from DB using Email
	accs, err := dbmgr.GetAccount("", acc.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// if length of accounts is zero, then account with id not found
	if len(accs) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No account found with given UserName"))
		return
	}

	if accs[0].Password != acc.Password {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid username/password"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
