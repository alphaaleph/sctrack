package database

import (
	"fmt"
	"github.com/alphaaleph/sctrack"
	"github.com/alphaaleph/sctrack/src/models"
	uuid2 "github.com/google/uuid"
	"golang.org/x/exp/slog"
	"time"
)

// get a slice of todos
func getTodosList(query string) ([]models.Todos, error) {

	rows, err := sctrack.Db.Query(query)
	if err != nil {
		sctrack.Log.Warn("Failed to get all todos", slog.String("Error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	// read all the todos entries
	var todos []models.Todos

	for rows.Next() {
		var td models.Todos
		rows.Scan(&td.UUID, &td.Created, &td.Description, &td.Completed, &td.CarrierID, &td.Action)
		todos = append(todos, td)
	}

	// return the actions
	return todos, nil
}

// GetTodos reads all the data from the todos table
func GetTodos() ([]models.Todos, error) {
	query := "SELECT uuid, created, description, completed, carrier_id, action FROM todos;"
	return getTodosList(query)
}

// GetTodosByCarrierID returns all the todos associated with a carrier
func GetTodosByCarrierID(carrierID string) ([]models.Todos, error) {

	query := fmt.Sprintf("SELECT uuid, created, description, completed, carrier_id, action FROM todos "+
		"WHERE carrier_id = '%s';", carrierID)
	return getTodosList(query)
}

// GetTodoByUUID returns the todos associated with the uuid
func GetTodoByUUID(uuid string) (*models.Todos, error) {

	query := fmt.Sprintf("SELECT uuid, created, description, completed, carrier_id, action FROM todos "+
		"WHERE uuid = '%s';", uuid)
	row := sctrack.Db.QueryRow(query)

	var todos models.Todos
	if err := row.Scan(&todos.UUID, &todos.Created, &todos.Description, &todos.Completed, &todos.CarrierID,
		&todos.Action); err != nil {
		sctrack.Log.Error("Fail to get todo", slog.String("Error", err.Error()))
		return nil, err
	}

	// return the information
	return &todos, nil
}

// DeleteTodosByCarrierID deletes the todos associated with a carrier
func DeleteTodosByCarrierID(carrierID string) error {

	stmt := fmt.Sprintf("DELETE FROM todos WHERE carrier_id = '%s';", carrierID)

	_, err := sctrack.Db.Exec(stmt)
	if err != nil {
		sctrack.Log.Warn("Failed to delete todos", slog.String("carrier_id", carrierID), slog.String("Error",
			err.Error()))
		return err
	}
	return nil
}

// DeleteTodosByUUID deletes a todos matching the uuid
func DeleteTodosByUUID(uuid string) error {

	stmt := fmt.Sprintf("DELETE FROM todos WHERE uuid = '%s';", uuid)

	_, err := sctrack.Db.Exec(stmt)
	if err != nil {
		sctrack.Log.Warn("Failed to delete todos", slog.String("uuid", uuid), slog.String("Error",
			err.Error()))
		return err
	}

	// add a journal entry
	var journal models.Journal
	journal.UUID = uuid2.NewString()
	journal.TodoUUID = uuid
	journal.Event = newEvent(models.Completed)
	AddJournal(journal)
	return nil
}

// AddTodo adds a new todos
func AddTodo(td models.TodosAdd) error {

	// create new values
	uuid := uuid2.NewString()
	created := time.Now().Format("2006-01-02 15:04:05.000")

	// add the todos
	stmt := fmt.Sprintf("INSERT INTO todos (uuid, created, description, completed, carrier_id, action) "+
		"VALUES ('%s', '%s', '%s', '%v', '%s', '%s');", uuid, created, td.Description, false, td.CarrierID, td.Action)
	_, err := sctrack.Db.Exec(stmt)
	if err != nil {
		sctrack.Log.Warn("Failed to add todos", slog.String("Error", err.Error()))
		return err
	}

	// add a journal entry
	var journal models.Journal
	journal.UUID = uuid2.NewString()
	journal.TodoUUID = uuid
	journal.Event = newEvent(td.Action)
	AddJournal(journal)

	return nil
}
