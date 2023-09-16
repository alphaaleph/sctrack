package handlers

import (
	"github.com/alphaaleph/sctrack/src/database"
	"github.com/alphaaleph/sctrack/src/models"
	"net/http"
)

// @Summary Get all actions
// @Description Get all action entries
// @Tags actions
// @Accept json
// @Produce json
// @Success 200 {array} models.Action
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/action/all [get]
func GetActions(w http.ResponseWriter, r *http.Request) {

	// get the actions
	var actions []models.ActionTable
	var err error

	if actions, err = database.GetActions(); err != nil {
		writeErrorResponse(w, http.StatusNoContent, "Actions not found")
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, actions)
}
