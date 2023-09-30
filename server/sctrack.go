package server

import (
	"database/sql"
	"golang.org/x/exp/slog"
	"os"
)

// Log will be used through the program to provide logging information
var Log *slog.Logger
var Db *sql.DB

// init used to get the program's logging started
func init() {
	// TODO: these options should be set with a configuration file
	opts := &slog.HandlerOptions{
		//AddSource: true,
		Level: slog.LevelDebug,
	}
	// set the logging handler
	Log = slog.New(slog.NewJSONHandler(os.Stdout, opts))
}
