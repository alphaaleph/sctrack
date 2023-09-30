package database

import (
	"fmt"
	"github.com/alphaaleph/sctrack/server"
	models2 "github.com/alphaaleph/sctrack/server/src/models"
	"golang.org/x/exp/slog"
)

// GetCarriers reads all the data from the carrier table
func GetCarriers() ([]models2.Carrier, error) {

	rows, err := server.Db.Query(sqlGetCarriers)
	if err != nil {
		server.Log.Warn("GetCarriers fail", slog.String("Error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	// read all the carrier entries
	var carriers []models2.Carrier

	for rows.Next() {
		var carrier models2.Carrier
		rows.Scan(&carrier.ID, &carrier.Name, &carrier.Telephone)
		carriers = append(carriers, carrier)
	}

	// return the carriers
	return carriers, nil
}

// GetCarrierDataByID reads a selected carrier's data which includes the carriers info, the tasks, and the journal
// entries
func GetCarrierDataByID(carrierID string) (*models2.CarrierData, error) {

	// read carrier data
	query := fmt.Sprintf(sqlGetCarrierDataByID, carrierID)
	row := server.Db.QueryRow(query)

	var err error
	var carrier models2.Carrier

	if err := row.Scan(&carrier.ID, &carrier.Name, &carrier.Telephone); err != nil {
		server.Log.Error("GetCarrierDataByID fail", slog.String("Error", err.Error()))
		return nil, err
	}

	// get the carrier's todos
	var todos []models2.Todos
	if todos, err = GetTodosByCarrierID(carrierID); err != nil {
		server.Log.Error("GetCarrierDataByID fail", slog.String("Error", err.Error()))
		return nil, err
	}

	// get the carrier's journals
	var journals []models2.Journal
	if journals, err = GetJournalsByCarrierID(carrierID); err != nil {
		server.Log.Error("GetCarrierDataByID fail", slog.String("Error", err.Error()))
		return nil, err
	}

	// return the response
	carrierData := models2.CarrierData{
		Carrier: carrier,
		Todos:   todos,
		Journal: journals,
	}

	// return the carriers info
	return &carrierData, nil
}

// DeleteCarrierByID deletes a carrier based on the ID
func DeleteCarrierByID(carrierID string) (int64, error) {

	// delete carrier and data
	stmt := fmt.Sprintf(sqlDeleteCarrierByID, carrierID)

	// delete and fins how many rows were affected
	r, err := server.Db.Exec(stmt)
	if err != nil {
		server.Log.Warn("DeleteCarrierByID fail", slog.String("carrier id", carrierID),
			slog.String("Error", err.Error()))
		return 0, err
	}

	rowsAffected, _ := r.RowsAffected()
	return rowsAffected, nil
}

// AddCarrier adds a new carrier
func AddCarrier(carrier models2.Carrier) error {

	// add the carrier
	stmt := fmt.Sprintf(sqlAddCarrier, carrier.ID, carrier.Name, carrier.Telephone)
	_, err := server.Db.Exec(stmt)
	if err != nil {
		server.Log.Warn("AddCarrier fail", slog.String("Error", err.Error()))
		return err
	}
	return nil
}
