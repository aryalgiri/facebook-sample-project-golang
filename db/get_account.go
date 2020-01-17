package db

import (
	"fmt"

	"github.com/google/logger"

	"github.com/sample-fb/models"
)

const (
	whereID    = "WHERE id="
	whereEmail = "WHERE email="
)

// GetAccount is used to retrieve account details from DB
func (c *dbClient) GetAccount(id, email string) ([]models.Account, error) {
	logger.Infof("retrieving account details from db with filter : %v, %v", id, email)

	sqlQuery := "SELECT * from accounts"
	if id != "" {
		sqlQuery = fmt.Sprintf("%v %v'%v'", sqlQuery, whereID, id)
	} else if email != "" {
		sqlQuery = fmt.Sprintf("%v %v'%v'", sqlQuery, whereEmail, email)
	}
	logger.Infof("get account details from db using query: %v", sqlQuery)

	rows, err := c.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	var accounts []models.Account

	for rows.Next() {
		var account models.Account
		err = rows.Scan(&account.ID, &account.Firstname, &account.MiddleName, &account.Lastname, &account.Email,
			&account.PhoneNumber, &account.Password, &account.Gender, &account.Address, &account.Status)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
