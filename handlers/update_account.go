package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sample-fb/models"
)

// UpdateAccount is used to update the account details of user
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	// read url params of the request
	urlParams := mux.Vars(r)

	var account models.Account
	var isUpdate bool

	// read the body of the request
	body, err := ioutil.ReadAll(r.Body)
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

	for i := range accounts {
		if urlParams["id"] == accounts[i].ID {
			if account.FirstName != "" {
				accounts[i].FirstName = account.FirstName
				isUpdate = true
			}
			if account.MiddleName != "" {
				accounts[i].MiddleName = account.MiddleName
				isUpdate = true
			}
			if account.LastName != "" {
				accounts[i].LastName = account.LastName
				isUpdate = true
			}
			if account.Email != "" {
				accounts[i].Email = account.Email
				isUpdate = true
			}
			if account.PhoneNumber != "" {
				accounts[i].PhoneNumber = account.PhoneNumber
				isUpdate = true
			}
			if account.Password != "" {
				if account.OldPassword == "" || account.OldPassword != accounts[i].Password {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("old password did not match"))
					return
				}

				accounts[i].Password = account.Password
				isUpdate = true
			}
			if account.Gender != "" {
				accounts[i].Gender = account.Gender
				isUpdate = true
			}
			if account.Address != nil {
				if account.Address.HouseNumber != "" {
					accounts[i].Address.HouseNumber = account.Address.HouseNumber
					isUpdate = true
				}
				if account.Address.Street != "" {
					accounts[i].Address.Street = account.Address.Street
					isUpdate = true
				}
				if account.Address.City != "" {
					accounts[i].Address.City = account.Address.City
					isUpdate = true
				}
				if account.Address.ZipCode != "" {
					accounts[i].Address.ZipCode = account.Address.ZipCode
					isUpdate = true

				}
			}
		}
	}

	if isUpdate == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No changes detected"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
