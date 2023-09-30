package database

import (
	"github.com/alphaaleph/sctrack/server"
	"github.com/alphaaleph/sctrack/server/src/models"
	"golang.org/x/exp/slog"
)

// GetActions reads all the data from the action_enum_table
func GetActions() ([]models.ActionTable, error) {

	rows, err := server.Db.Query(sqlGetActions)
	if err != nil {
		server.Log.Warn("Failed to get all actions", slog.String("Error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	// read all the action entries
	var actions []models.ActionTable

	for rows.Next() {
		var action models.ActionTable
		rows.Scan(&action.Action)
		actions = append(actions, action)
	}

	// return the actions
	return actions, nil
}
