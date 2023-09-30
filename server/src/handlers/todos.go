package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/alphaaleph/sctrack/server"
	"github.com/alphaaleph/sctrack/server/src/database"
	"github.com/alphaaleph/sctrack/server/src/models"
	uuid2 "github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"net/http"
)

// @Summary Get all todos
// @Description Get all the entries in the todos list
// @Tags todos
// @Accept json
// @Produce json
// @Router /api/todos/all [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {

	// get the todos
	var todos []models.Todos
	var err error

	if todos, err = database.GetTodos(); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Todos error: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, todos)
}

// @Summary Get todos by carrier id
// @Description Get all the entries in the todos list that match the carrier id
// @Tags todos
// @Accept json
// @Produce json
// @Param carrier_id path string true "carrier_id"
// @Router /api/todos/carrier/{carrier_id} [get]
func GetTodosByCarrierID(w http.ResponseWriter, r *http.Request) {

	carrierID := mux.Vars(r)["carrier_id"]

	// get the todos
	var todos []models.Todos
	var err error

	if todos, err = database.GetTodosByCarrierID(carrierID); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Todos error: %s", err.Error()))
		return
	}

	// return the response
	if todos == nil {
		writeErrorResponse(w, http.StatusNotFound, fmt.Sprintf("Todos data not found"))
		return
	}
	writeJSONResponse(w, http.StatusOK, todos)
}

// @Summary Get todos by carrier uuid
// @Description Get all the entries in the todos list that match the uuid
// @Tags todos
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Router /api/todos/{uuid} [get]
func GetTodosByUUID(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]

	// get the todos
	var todos *models.Todos
	var err error

	if todos, err = database.GetTodoByUUID(uuid); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeErrorResponse(w, http.StatusNotFound, fmt.Sprintf("Todos data not found: %s", err.Error()))
			return
		}
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Todos error: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, todos)
}

// @Summary Delete todos by carrier_id
// @Description Delete entries in the todos list that match an carrier_id
// @Tags todos
// @Accept json
// @Produce json
// @Param carrier_id path string true "carrier_id"
// @Router /api/todos/carrier/{carrier_id} [delete]
func DeleteTodosByCarrierID(w http.ResponseWriter, r *http.Request) {

	carrierID := mux.Vars(r)["carrier_id"]

	var (
		rowsAffected int64
		err          error
	)
	if rowsAffected, err = database.DeleteTodosByCarrierID(carrierID); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Todos error: %s", err.Error()))
	}

	// return the response
	if rowsAffected == 0 {
		writeErrorResponse(w, http.StatusNotFound, fmt.Sprintf("Todos data not found"))
		return
	}
	writeJSONResponse(w, http.StatusOK, "Success")
}

// @Summary Delete todos by uuid
// @Description Delete an entry in the todos list that match an uuid
// @Tags todos
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Router /api/todos/{uuid} [delete]
func DeleteTodosByUUID(w http.ResponseWriter, r *http.Request) {

	uuid, _ := uuid2.Parse(mux.Vars(r)["uuid"])

	var (
		rowsAffected int64
		err          error
	)
	if rowsAffected, err = database.DeleteTodosByUUID(uuid); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Todos error: %s", err.Error()))
	}

	// return the response
	if rowsAffected == 0 {
		writeErrorResponse(w, http.StatusNotFound, fmt.Sprintf("Todos data not found"))
		return
	}
	writeJSONResponse(w, http.StatusOK, "Success")
}

// @Summary Add todos
// @Description Add a todos entry for a carrier
// @Tags todos
// @Accept json
// @Produce json
// @Param data body models.TodosAdd true "New Todos"
// @Router /api/todos [post]
func AddTodos(w http.ResponseWriter, r *http.Request) {

	var td models.TodosAdd
	err := extractBodyJSON(r, &td)
	if err != nil {
		server.Log.Warn("Failed parse todos add object", slog.String("Error", err.Error()))
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Todos error: %s", err.Error()))
		return
	}

	// add todos
	if err := database.AddTodo(td); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Todos error: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, "Success")
}

// @Summary Update the todos completed
// @Description Update a todos completed flag for a carrier
// @Tags todos
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Param data body models.TodosStatus true "Update Completed"
// @Router /api/todos/{uuid}/completed [patch]
func PatchTodosCompleted(w http.ResponseWriter, r *http.Request) {

	// get the uuid
	uuid, _ := uuid2.Parse(mux.Vars(r)["uuid"])

	// read the body
	var ts models.TodosStatus
	err := extractBodyJSON(r, &ts)
	if err != nil {
		server.Log.Warn("Failed parse todos update object", slog.String("Error", err.Error()))
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Todos error: %s", err.Error()))
		return
	}

	// update todos
	if err := database.UpdateTodosCompleted(uuid, ts.Completed); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Todos error: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, "Success")
}
