package main

import (
	"github.com/alphaaleph/sctrack/server"
	_ "github.com/alphaaleph/sctrack/server/docs"
	"github.com/alphaaleph/sctrack/server/src/database"
	"github.com/alphaaleph/sctrack/server/src/handlers"
	"os"
)

// main entry point to the application
//
//	@title Sctrack
//	@BasePath  /
//	@schemes http
func main() {
	os.Exit(execute())
}

// execute starts the project
func execute() int {
	server.Log.Info("Starting sctrack ... ")

	// initialize the database
	server.Db = database.DBInstance()
	if server.Db == nil {
		panic("Error connecting to database")
	}
	defer server.Db.Close()

	// start the web service
	handlers.StartYourEngines()

	return 0
}
