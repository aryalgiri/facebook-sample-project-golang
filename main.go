package main

import (
	"os"

	"github.com/google/logger"
	_ "github.com/lib/pq"

	"github.com/sample-fb/handlers"
)

func main() {
	// create a file to store logs
	logFile, err := os.Create("logs.txt")
	if err != nil {
		panic(err)
	}
	// Initialize logger
	logger.Init("sample-db", false, true, logFile)
	// Initialize the routes
	handlers.Init()
}
