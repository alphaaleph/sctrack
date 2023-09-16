package handlers

import (
	"fmt"
	"github.com/alphaaleph/sctrack"
	"github.com/alphaaleph/sctrack/src/database"
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
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/carrier/all [get]
func GetCarriers(w http.ResponseWriter, r *http.Request) {

	// get the carriers
	var carriers []models.Carrier
	var err error

	if carriers, err = database.GetCarriers(); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Carriers not found: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, carriers)
}

// @Summary Get carrier's data
// @Description Get carrier's data details by ID
// @Tags carriers
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.CarrierData
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/carrier/{id} [get]
func GetCarrierDataByID(w http.ResponseWriter, r *http.Request) {

	ID := mux.Vars(r)["id"]

	// get carrier
	var carrierData *models.CarrierData
	var err error

	if carrierData, err = database.GetCarrierDataByID(ID); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Carrier data not found: %s", err.Error()))
		return
	}

	// return the carrier info
	writeJSONResponse(w, http.StatusOK, carrierData)
}

// @Summary Delete carrier
// @Description Delete a carrier
// @Tags carriers
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 ""
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/carrier/{id} [delete]
func DeleteCarrierByID(w http.ResponseWriter, r *http.Request) {

	ID := mux.Vars(r)["id"]

	if err := database.DeleteCarrierByID(ID); err != nil {
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Carrier delete failed: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, "Success")
}

// @Summary Add carrier
// @Description Add a new carrier
// @Tags carriers
// @Accept json
// @Produce json
// @Param data body models.Carrier true "The Carrier Inout"
// @Success 200 ""
// @Failure 400 {object} models.DBError
// @Failure 404 {object} models.DBError
// @Router /api/carrier [post]
func AddCarrier(w http.ResponseWriter, r *http.Request) {

	var carrier models.Carrier
	err := extractBodyJSON(r, &carrier)
	if err != nil {
		sctrack.Log.Warn("Failed parse carrier object", slog.String("Error", err.Error()))
		writeErrorResponse(w, http.StatusNoContent, fmt.Sprintf("Failed carrier parse: %s", err.Error()))
		return
	}

	// add the carrier
	if err = database.AddCarrier(carrier); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Failed carrier insert: %s", err.Error()))
		return
	}

	// return the response
	writeJSONResponse(w, http.StatusOK, "Success")
}
