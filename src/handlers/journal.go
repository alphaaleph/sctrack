package handlers

import (
	"fmt"
	"github.com/alphaaleph/sctrack/src/database"
	"github.com/alphaaleph/sctrack/src/models"
	"github.com/gorilla/mux"
	"net/http"
)

// @Summary Get journals
// @Description Get all entries from the journal
// @Tags journal
// @Accept json
// @Produce json
// @Success 200 {array} models.Journal
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/journal/all [get]
func GetJournals(w http.ResponseWriter, r *http.Request) {

	// get the journals
	var journals []models.Journal
	var err error

	if journals, err = database.GetJournals(); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Journals not found: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, journals)
}

// @Summary Get a journal
// @Description Get a journal entry that matches the uuid
// @Tags journal
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 200 {object} models.Journal
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/journal/{uuid} [get]
func GetJournalByUUID(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]

	// get the todos
	journal := new(models.Journal)
	var err error

	if journal, err = database.GetJournalByUUID(uuid); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Journal not found: %s", err.Error()))
		return
	}

	// return journal
	writeJSONResponse(w, http.StatusOK, journal)
}

// @Summary Delete journal
// @Description Delete an entry in the journal by UUID
// @Tags journal
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 200 ""
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/journal/{uuid} [delete]
func DeleteJournalByUUID(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]

	if err := database.DeleteJournalByUUID(uuid); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Journal delete failed: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, "Success")
}

// @Summary Delete journal
// @Description Delete an entry in the journal by UUID and Index
// @Tags journal
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Param index path int true "index"
// @Success 200 ""
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/journal/{uuid}/{index} [delete]
func DeleteJournalByUUIDIndex(w http.ResponseWriter, r *http.Request) {

	uuid := mux.Vars(r)["uuid"]
	index := nUint(mux.Vars(r)["index"])

	if err := database.DeleteJournalByUUIDIndex(index, uuid); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Journal delete failed: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, "Success")
}
