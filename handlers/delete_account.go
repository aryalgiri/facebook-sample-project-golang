package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sample-fb/models"
)

// deleteAccount is used to delete the account
func deleteAccount(w http.ResponseWriter, r *http.Request) {
	// capture URL Params
	urlParams := mux.Vars(r)

	// Get Account details from DB using ID
	accs, err := dbmgr.GetAccount(urlParams["id"], "")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// if length of accounts is zero, then account with id not found
	if len(accs) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No account found with given ID"))
		return
	}

	// update account status to deleted
	accs[0].Status = models.StatusDeleted

	// save account details to db
	if err = dbmgr.SaveAccount(&accs[0]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
