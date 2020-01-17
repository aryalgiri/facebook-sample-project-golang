package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sample-fb/config"
	"github.com/sample-fb/db"
)

var (
	dbmgr db.Datastore
)

// Init will initialize the handlers
func Init() {
	// build db manager
	dbmgr = db.NewClient()
	// Initialize the router and handlers
	r := mux.NewRouter()
	r.HandleFunc("/", healthCheck).Methods(http.MethodGet)
	r.HandleFunc("/v1/account", createAccount).Methods(http.MethodPost)
	r.HandleFunc("/v1/account/login", loginHandler).Methods(http.MethodPost)
	r.HandleFunc("/v1/account", listAccounts).Methods(http.MethodGet)
	r.HandleFunc("/v1/account/{id}", updateAccount).Methods(http.MethodPatch)
	r.HandleFunc("/v1/account/{id}", deleteAccount).Methods(http.MethodDelete)

	http.ListenAndServe(":"+config.Port, r)
}
