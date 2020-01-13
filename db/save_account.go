package db

import "github.com/sample-fb/models"

//SaveAccount is used to save the account details into DB
func SaveAccount(account models.Account) {
	logger.Infof("saving %v account details into db", account.ID)
	_, err = db.QueryRow("Insert Into accounts() values($1,$2,$3) returning uid:", "astaxie", "2012-12-09")
	return err
}
