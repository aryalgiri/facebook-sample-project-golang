package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" //initizing driver
	"github.com/sample-fb/config"
)

var (
	db  *sql.DB
	err error
)

// NewClient is used to establish a connection between microservice and database
func NewClient() {
	var err error
	//dbable database connections
	db, err = sql.Open(config.DBName, fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUserName, config.DBPassword, config.DBName))
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully established DB connection")
}
