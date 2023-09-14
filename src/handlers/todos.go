package handlers

import (
	"encoding/json"
	"github.com/alphaaleph/sctrack"
	"github.com/alphaaleph/sctrack/src/app"
	"github.com/alphaaleph/sctrack/src/models"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"net/http"
)

// @Summary Get all todos
// @Description Get all the entries in the todo list
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} models.TodosData
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/todos/all [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {

	query := `SELECT carrier.id, carrier.name, todos.uuid, todos.title, todos.completed
				FROM carrier
				INNER JOIN todos ON todos.carrier_id = carrier.id
				ORDER BY carrier.id;`

	rows, err := sctrack.Db.Query(query)
	if err != nil {
		sctrack.Log.Warn("Failed to get all todos", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Todos entries not found")
		return
	}
	defer rows.Close()

	// read all the todo entries
	todos := []models.TodosData{}

	for rows.Next() {
		var td models.TodosData
		rows.Scan(&td.CarrierID, &td.Carrier, &td.UUID, &td.Title, &td.Completed)
		todos = append(todos, td)
	}

	response := json.NewEncoder(w).Encode(todos)
	app.WriteJSONResponse(w, http.StatusOK, response)
}

// @Summary Get todo by carrier uuid
// @Description Get all the entries in the todo list that match the uuid
// @Tags todos
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 200 {array} models.Todos
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/todos/{uuid} [get]
func GetTodosEntry(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]

	query := `SELECT * FROM todos WHERE todos.uuid = $1;`
	row := sctrack.Db.QueryRow(query, uuid)

	var todos models.Todos
	if err := row.Scan(&todos.UUID, &todos.Title, &todos.Completed, &todos.CarrierID); err != nil {
		sctrack.Log.Error("GetTodosEntry fail", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNotFound, "Todos read failed")
		return
	}

	app.WriteJSONResponse(w, http.StatusOK, todos)
}

// @Summary Delete todo by id
// @Description Delete an entry in the todo list that match an id
// @Tags todos
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 200 ""
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/todos/{uuid} [delete]
func DeleteTodoByUUID(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]

	stmt := `DELETE FROM todos WHERE uuid = $1;`

	_, err := sctrack.Db.Exec(stmt, uuid)
	if err != nil {
		sctrack.Log.Warn("Failed to delete a todo", slog.String("uuid", uuid), slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Todo delete failed")
		return
	}

	app.WriteJSONResponse(w, http.StatusOK, "Success")
}

// @Summary Add todo
// @Description Add a todo entry for a carrier
// @Tags todos
// @Accept json
// @Produce json
// @Param data body models.Todos true "The Todos Inout"
// @Success 200 ""
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/todos [post]
func AddTodos(w http.ResponseWriter, r *http.Request) {

	var todos models.Todos
	err := app.ExtractBodyJSON(r, &todos)
	if err != nil {
		sctrack.Log.Warn("Failed to delete a todo by id", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Failed todos parse")
		return
	}

	// add todos
	stmt := `INSERT INTO todos (uuid, title, completed, carrier_id) VALUES ($1, $2, $3, $4);`
	_, err = sctrack.Db.Exec(stmt, todos.UUID, todos.Title, todos.Completed, todos.CarrierID)
	if err != nil {
		sctrack.Log.Warn("Failed to delete a todo by id", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusBadRequest, "Failed todos insert")
		return
	}

	// return information
	app.WriteJSONResponse(w, http.StatusOK, "Success")
}
