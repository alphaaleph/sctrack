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

// @Summary Get journals
// @Description Get all entries from the journal
// @Tags journal
// @Accept json
// @Produce json
// @Success 200 {array} models.JournalData
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/journal/all [get]
func GetJournal(w http.ResponseWriter, r *http.Request) {

	query := `SELECT carrier.id, carrier.name, journal.uuid, journal.event
				FROM carrier
				INNER JOIN journal ON journal.carrier_id = carrier.id
				ORDER BY carrier.id;`

	rows, err := sctrack.Db.Query(query)
	if err != nil {
		sctrack.Log.Warn("Failed to get all journal entries", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Journal entries not found")
		return
	}
	defer rows.Close()

	// read all the journal entries
	journal := []models.JournalData{}

	for rows.Next() {
		var j models.JournalData
		rows.Scan(&j.CarrierID, &j.Carrier, &j.UUID, j.Event)
		journal = append(journal, j)
	}

	response := json.NewEncoder(w).Encode(journal)
	app.WriteJSONResponse(w, http.StatusOK, response)
}

// @Summary Get a journal
// @Description Get a journal entry that matches the uuid
// @Tags journal
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 200 {object} models.Journal
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/journal/{uuid} [get]
func GetJournalEntry(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]

	query := `SELECT * FROM journal WHERE journal.uuid = $1;`
	row := sctrack.Db.QueryRow(query, uuid)

	var journal models.Journal
	if err := row.Scan(&journal.UUID, &journal.Event, &journal.CarrierID); err != nil {
		sctrack.Log.Error("GetJournalEntry fail", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNotFound, "Journal read failed")
		return
	}

	app.WriteJSONResponse(w, http.StatusOK, journal)
}

// @Summary Delete journal
// @Description Delete an entry in the journal
// @Tags journal
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 200 ""
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/journal/{uuid} [delete]
func DeleteJournalByUUID(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]

	stmt := `DELETE FROM journal WHERE journal.uuid = $1;`

	_, err := sctrack.Db.Exec(stmt, uuid)
	if err != nil {
		sctrack.Log.Warn("Failed to delete a journal", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Journal delete failed")
		return
	}

	app.WriteJSONResponse(w, http.StatusOK, "Success")
}

// @Summary Add a journal
// @Description Add a new journal entry
// @Tags journal
// @Accept json
// @Produce json
// @Param data body models.Journal true "The Journal Inout"
// @Success 200 ""
// @Failure 400 {object} models.Error
// @Failure 404 {object} models.Error
// @Router /api/journal [post]
func AddJournal(w http.ResponseWriter, r *http.Request) {

	var journal models.Journal
	err := app.ExtractBodyJSON(r, &journal)
	if err != nil {
		sctrack.Log.Warn("Failed parse journal object", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusNoContent, "Failed journal parse")
		return
	}

	// add journal
	stmt := `INSERT INTO journal (uuid, event, carrier_id) VALUES ($1, $2, $3);`
	_, err = sctrack.Db.Exec(stmt, journal.UUID, journal.Event, journal.CarrierID)
	if err != nil {
		sctrack.Log.Warn("Failed to add journal", slog.String("Error", err.Error()))
		app.WriteErrorResponse(w, http.StatusBadRequest, "Failed journal insert")
		return
	}

	// return information
	app.WriteJSONResponse(w, http.StatusOK, "Success")
}
