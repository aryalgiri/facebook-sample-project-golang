package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sample-fb/config"
	"github.com/sample-fb/db"
	"github.com/sample-fb/handlers"
)

func main() {
	//establish a database connection
	db.NewClient()

	r := mux.NewRouter()
	r.HandleFunc("/v1/account", handlers.CreateAccount).Methods(http.MethodPost)
	r.HandleFunc("/v1/account/login", handlers.LoginHandler).Methods(http.MethodPost)

	r.HandleFunc("/v1/account", handlers.ListAccounts).Methods(http.MethodGet)

	r.HandleFunc("/v1/account/{id}", handlers.DeleteAccount).Methods(http.MethodDelete)

	r.HandleFunc("/v1/account/{id}", handlers.UpdateAccount).Methods(http.MethodPatch)

	http.ListenAndServe(":"+config.Port, r)

}
