package database

import (
	"fmt"
	"github.com/alphaaleph/sctrack/server"
	models2 "github.com/alphaaleph/sctrack/server/src/models"
	uuid2 "github.com/google/uuid"
	"golang.org/x/exp/slog"
	"time"
)

// get a slice of todos
func getTodosList(query string) ([]models2.Todos, error) {

	rows, err := server.Db.Query(query)
	if err != nil {
		server.Log.Warn("getTodosList fail", slog.String("Error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	// read all the todos entries
	var todos []models2.Todos

	for rows.Next() {
		var td models2.Todos
		rows.Scan(&td.UUID, &td.Created, &td.Description, &td.Completed, &td.CarrierID, &td.Action)
		todos = append(todos, td)
	}

	// return the actions
	return todos, nil
}

// GetTodos reads all the data from the todos table
func GetTodos() ([]models2.Todos, error) {
	return getTodosList(sqlGetTodos)
}

// GetTodosByCarrierID returns all the todos associated with a carrier
func GetTodosByCarrierID(carrierID string) ([]models2.Todos, error) {
	query := fmt.Sprintf(sqlGetTodosByCarrierID, carrierID)
	return getTodosList(query)
}

// GetTodoByUUID returns the todos associated with the uuid
func GetTodoByUUID(uuid string) (*models2.Todos, error) {

	query := fmt.Sprintf(sqlGetTodosByUUID, uuid)
	row := server.Db.QueryRow(query)

	var todos models2.Todos
	if err := row.Scan(&todos.UUID, &todos.Created, &todos.Description, &todos.Completed, &todos.CarrierID,
		&todos.Action); err != nil {
		server.Log.Error("GetTodoByUUID fail", slog.String("Error", err.Error()))
		return nil, err
	}

	// return the information
	return &todos, nil
}

// DeleteTodosByCarrierID deletes the todos associated with a carrier
func DeleteTodosByCarrierID(carrierID string) (int64, error) {

	//stmt := fmt.Sprintf("DELETE FROM todos WHERE carrier_id = '%s';", carrierID)
	stmt := fmt.Sprintf(sqlDeleteTodosByCarrierID, carrierID)

	r, err := server.Db.Exec(stmt)
	if err != nil {
		server.Log.Warn("DeleteTodosByCarrierID fail", slog.String("carrier_id", carrierID), slog.String("Error",
			err.Error()))
		return 0, err
	}

	// return affected rows
	rowsAffected, _ := r.RowsAffected()
	return rowsAffected, nil
}

// DeleteTodosByUUID deletes a todos matching the uuid
func DeleteTodosByUUID(uuid uuid2.UUID) (int64, error) {

	//stmt := fmt.Sprintf("DELETE FROM todos WHERE uuid = '%s';", uuid)
	stmt := fmt.Sprintf(sqlDeleteTodosByUUID, uuid)

	r, err := server.Db.Exec(stmt)
	if err != nil {
		server.Log.Warn("DeleteTodosByUUID fail", slog.String("uuid", (uuid).String()), slog.String("Error",
			err.Error()))
		return 0, err
	}

	// add a journal entry
	var journal models2.Journal
	journal.UUID = uuid2.New()
	journal.TodoUUID = uuid
	journal.Event = newEvent(models2.Completed)
	AddJournal(journal)

	//return affected rows
	rowsAffected, _ := r.RowsAffected()
	return rowsAffected, nil
}

// AddTodo adds a new todos
func AddTodo(td models2.TodosAdd) error {

	// create new values
	uuid := uuid2.New()
	created := time.Now().Format("2006-01-02 15:04:05.000")

	// add the todos
	stmt := fmt.Sprintf(sqlAddTodos, uuid, created, td.Description, false, td.CarrierID, td.Action)
	_, err := server.Db.Exec(stmt)
	if err != nil {
		server.Log.Warn("AddTodo fail", slog.String("Error", err.Error()))
		return err
	}

	// add a journal entry
	var journal models2.Journal
	journal.UUID = uuid2.New()
	journal.TodoUUID = uuid
	journal.Event = newEvent(td.Action)
	AddJournal(journal)

	return nil
}

// UpdateTodosCompleted updates a todos completed flag
func UpdateTodosCompleted(uuid uuid2.UUID, completed bool) error {

	// update the todos
	stmt := fmt.Sprintf(sqlPatchTodosCompleted, completed, uuid)
	_, err := server.Db.Exec(stmt)
	if err != nil {
		server.Log.Warn("UpdateTodosCompleted fail", slog.String("Error", err.Error()))
		return err
	}

	// add a journal entry
	var journal models2.Journal
	journal.UUID = uuid2.New()
	journal.TodoUUID = uuid
	journal.Event = newEvent(models2.Completed)
	AddJournal(journal)

	return nil
}
