package main

import (
	"github.com/alphaaleph/sctrack"
	_ "github.com/alphaaleph/sctrack/docs"
	"github.com/alphaaleph/sctrack/src/database"
	"github.com/alphaaleph/sctrack/src/handlers"
	"os"
)

// @host   localhost
// @BasePath  /
// @schemes http
func main() {
	os.Exit(execute())
}

// execute starts the project
func execute() int {
	sctrack.Log.Info("Starting sctrack ... ")

	// initialize the database
	sctrack.Db = database.DBInstance()
	if sctrack.Db == nil {
		panic("Error connecting to database")
	}
	defer sctrack.Db.Close()

	// start the web service
	handlers.StartYourEngines()

	return 0
}
