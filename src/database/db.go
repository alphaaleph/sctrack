package database

import (
	"database/sql"
	"fmt"
	"github.com/alphaaleph/sctrack"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
	"os"
	"sync"
)

var (
	db       *sql.DB
	onceProd sync.Once
)

// DBInstance instantiates the database connection and sets a global connection to it.
func DBInstance() *sql.DB {
	onceProd.Do(func() {
		// create singleton
		var err error
		if db, err = connect(); err != nil {
			sctrack.Log.Debug("Database singleton failed.")
		}
		sctrack.Log.Debug("Database singleton created.")
	})

	// return the database instance
	return db
}

// connect starts a new connection with the database represented by the passed sql instance. All database connection
// information is read locally from the .env file, and from the environment variables in the cloud service.
func connect() (*sql.DB, error) {

	// create the database url connection
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=require",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	sctrack.Log.Debug("Connect string", slog.String("DB URL", url))

	// open a connection to the database
	db, err := sql.Open(os.Getenv("DB_DRIVER"), url)
	if err != nil {
		sctrack.Log.Error(fmt.Sprintf("Database connect failed: %s", err.Error()))
		return nil, err
	}

	sctrack.Log.Info(fmt.Sprintf("Connected to database"))
	return db, nil
}

// Close closes the database connection is a database instance is available.
func Close() {
	if sctrack.Db != nil {
		sctrack.Log.Debug("Closing the database")
		sctrack.Db.Close()
	}
}
