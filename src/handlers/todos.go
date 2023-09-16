package handlers

import (
	"fmt"
	"github.com/alphaaleph/sctrack"
	"github.com/alphaaleph/sctrack/src/database"
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
// @Success 200 {array} models.Todos
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/todos/all [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {

	// get the todos
	var todos []models.Todos
	var err error

	if todos, err = database.GetTodos(); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Todos not found: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, todos)
}

// @Summary Get todos by carrier id
// @Description Get all the entries in the todo list that match the carrier id
// @Tags todos
// @Accept json
// @Produce json
// @Param carrier_id path string true "carrier_id"
// @Success 200 {array} models.Todos
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/todos/carrier/{carrier_id} [get]
func GetTodosByCarrierID(w http.ResponseWriter, r *http.Request) {

	carrierID := mux.Vars(r)["carrier_id"]

	// get the todos
	var todos []models.Todos
	var err error

	if todos, err = database.GetTodosByCarrierID(carrierID); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Todos not found: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, todos)
}

// @Summary Get todo by carrier uuid
// @Description Get all the entries in the todo list that match the uuid
// @Tags todos
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 200 {object} models.Todos
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/todos/{uuid} [get]
func GetTodoByUUID(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]

	// get the todos
	var todos *models.Todos
	var err error

	if todos, err = database.GetTodoByUUID(uuid); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Todos not found: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, todos)
}

// @Summary Delete todos by carrier_id
// @Description Delete entries in the todo list that match an carrier_id
// @Tags todos
// @Accept json
// @Produce json
// @Param carrier_id path string true "carrier_id"
// @Success 200 ""
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/todos/carrier/{carrier_id} [delete]
func DeleteTodosByCarrierID(w http.ResponseWriter, r *http.Request) {

	carrierID := mux.Vars(r)["carrier_id"]

	if err := database.DeleteTodosByCarrierID(carrierID); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Todo delete failed: %s", err.Error()))
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, "Success")
}

// @Summary Delete todos by uuid
// @Description Delete an entry in the todos list that match an uuid
// @Tags todos
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 200 ""
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/todos/{uuid} [delete]
func DeleteTodoByUUID(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]

	if err := database.DeleteTodosByUUID(uuid); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Todo delete failed: %s", err.Error()))
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, "Success")
}

// @Summary Add todos
// @Description Add a todos entry for a carrier
// @Tags todos
// @Accept json
// @Produce json
// @Param data body models.TodosAdd true "New Todos"
// @Success 200 ""
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/todos [post]
func AddTodo(w http.ResponseWriter, r *http.Request) {

	var td models.TodosAdd
	err := extractBodyJSON(r, &td)
	if err != nil {
		sctrack.Log.Warn("Failed parse todos add object", slog.String("Error", err.Error()))
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Failed todos add parse: %s", err.Error()))
		return
	}

	// add todos
	if err := database.AddTodo(td); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Failed todos insert: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, "Success")
}
