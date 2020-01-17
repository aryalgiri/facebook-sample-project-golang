package db

import (
	"github.com/google/logger"

	"github.com/sample-fb/models"
)

// SaveAccount is used to save the account details into DB
func (c *dbClient) SaveAccount(account *models.Account) error {
	// check if account exists in the DB
	accs, err := c.GetAccount(account.ID, "")
	if err != nil {
		return err
	}

	logger.Infof("saving account details %v", account.ID)
	// save details into db
	// if account with ID already exists, just update the account
	if len(accs) == 0 {
		_, err = c.db.Exec("INSERT INTO accounts(id,firstname,middlename,lastname,email,phonenumber,password,gender,address,status) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);",
			account.ID, account.Firstname, account.MiddleName, account.Lastname, account.Email, account.PhoneNumber, account.Password, account.Gender, account.Address, account.Status)
	} else {
		_, err = c.db.Exec("UPDATE accounts SET firstname=$1,middlename=$2,lastname=$3,email=$4,phonenumber=$5,password=$6,gender=$7,address=$8,status=$9 WHERE id=$10;",
			account.Firstname, account.MiddleName, account.Lastname, account.Email, account.PhoneNumber, account.Password, account.Gender, account.Address, account.Status, account.ID)
	}
	return err
}
