package handlers

import (
	"encoding/json"
	"github.com/alphaaleph/sctrack"
	"github.com/alphaaleph/sctrack/src/app"
	"github.com/alphaaleph/sctrack/src/models"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
	"net/http"
)

// @Summary Get all carriers
// @Description Get the information for all carriers
// @Tags carriers
// @Accept json
// @Produce json
// @Success 200 {array} models.Carrier
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/carrier/all [get]
func GetCarriers(w http.ResponseWriter, r *http.Request) {

	query := `SELECT * FROM  carrier;`
	rows, err := sctrack.Db.Query(query)
	if err != nil {
		sctrack.Log.Warn("Failed to get all carriers", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Carriers not found")
		return
	}
	defer rows.Close()

	// read all the carrier entries
	carriers := []models.Carrier{}

	for rows.Next() {
		var carrier models.Carrier
		rows.Scan(&carrier.ID, &carrier.Name, &carrier.Telephone)
		carriers = append(carriers, carrier)
	}

	response := json.NewEncoder(w).Encode(carriers)
	app.WriteJSONResponse(w, http.StatusOK, response)
}

// @Summary Get carrier's data
// @Description Get carrier's data details by ID
// @Tags carriers
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.CarrierData
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/carrier/{id} [get]
func GetCarrierDataByID(w http.ResponseWriter, r *http.Request) {

	ID := mux.Vars(r)["id"]

	// read carrier data
	query := `SELECT * FROM carrier WHERE carrier.id = $1;`
	row := sctrack.Db.QueryRow(query, ID)

	var carrier models.Carrier
	if err := row.Scan(&carrier.ID, &carrier.Name, &carrier.Telephone); err != nil {
		sctrack.Log.Error("GetCarrierDataByID fail", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNotFound, "Carrier read failed")
		return
	}

	// get the carrier's todos
	query = `SELECT * FROM todos WHERE todos.carrier_id = $1;`
	rows, err := sctrack.Db.Query(query, ID)
	if err != nil {
		sctrack.Log.Warn("Failed to get all the carrier's todos", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Failed todos")
		return
	}
	defer rows.Close()

	todos := []models.Todos{}

	for rows.Next() {
		var td models.Todos
		rows.Scan(&td.UUID, &td.Title, &td.Completed, &td.CarrierID)
		todos = append(todos, td)
	}

	// get the carrier's journal
	query = `SELECT * FROM journal WHERE journal.carrier_id = $1;`
	rows, err = sctrack.Db.Query(query, ID)
	if err != nil {
		sctrack.Log.Warn("Failed to get all the carrier's journal", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Failed journal")
		return
	}

	journal := []models.Journal{}

	for rows.Next() {
		var je models.Journal
		rows.Scan(&je.UUID, &je.Event, &je.CarrierID)
		journal = append(journal, je)
	}

	// send the response
	response := models.CarrierData{
		Carrier: carrier,
		Todos:   todos,
		Journal: journal,
	}

	app.WriteJSONResponse(w, http.StatusOK, response)
}

// @Summary Delete carrier
// @Description Delete a carrier
// @Tags carriers
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 ""
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/carrier/{id} [delete]
func DeleteCarrierByID(w http.ResponseWriter, r *http.Request) {

	ID := mux.Vars(r)["id"]

	// delete carrier and data
	stmt := `DELETE FROM carrier WHERE id = $1;`

	_, err := sctrack.Db.Exec(stmt, ID)
	if err != nil {
		sctrack.Log.Warn("Failed to delete a todo by id", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Carrier delete failed")
		return
	}

	app.WriteJSONResponse(w, http.StatusOK, "Success")
}

// @Summary Add carrier
// @Description Add a new carrier
// @Tags carriers
// @Accept json
// @Produce json
// @Param data body models.Carrier true "The Carrier Inout"
// @Success 200 ""
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/carrier [post]
func AddCarrier(w http.ResponseWriter, r *http.Request) {

	var carrier models.Carrier
	err := app.ExtractBodyJSON(r, &carrier)
	if err != nil {
		sctrack.Log.Warn("Failed parse carrier object", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Failed carrier parse")
		return
	}

	// add the carrier
	stmt := `INSERT INTO carrier (id, name, telephone) VALUES ($1, $2, $3);`
	_, err = sctrack.Db.Exec(stmt, carrier.ID, carrier.Name, carrier.Telephone)
	if err != nil {
		sctrack.Log.Warn("Failed to add carrier", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusBadRequest, "Failed carrier insert")
		return
	}

	// return information
	app.WriteJSONResponse(w, http.StatusOK, "Success")
}
