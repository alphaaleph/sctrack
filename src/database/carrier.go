package database

import (
	"fmt"
	"github.com/alphaaleph/sctrack"
	"github.com/alphaaleph/sctrack/src/models"
	"golang.org/x/exp/slog"
)

const (
	sqlGetCarriers        = "SELECT * FROM  carrier;"
	sqlGetCarrierDataByID = "SELECT id, name, telephone FROM carrier WHERE id = '%s';"
	sqlDeleteCarrierByID  = "DELETE FROM carrier WHERE id = '%s';"
	sqlAddCarrier         = "INSERT INTO carrier (id, name, telephone) VALUES ('%s', '%s', '%s');"
)

// GetCarriers reads all the data from the carrier table
func GetCarriers() ([]models.Carrier, error) {

	// TODO: remove query := `SELECT * FROM  carrier;`
	rows, err := sctrack.Db.Query(sqlGetCarriers)
	if err != nil {
		sctrack.Log.Warn("Failed to get all carriers", slog.String("Error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	// read all the carrier entries
	var carriers []models.Carrier

	for rows.Next() {
		var carrier models.Carrier
		rows.Scan(&carrier.ID, &carrier.Name, &carrier.Telephone)
		carriers = append(carriers, carrier)
	}

	// return the carriers
	return carriers, nil
}

// GetCarrierDataByID reads a selected carrier's data which includes the carriers info, the tasks, and the journal
// entries
func GetCarrierDataByID(carrierID string) (*models.CarrierData, error) {

	// read carrier data
	// TODO: remove  query := fmt.Sprintf("SELECT id, name, telephone FROM carrier WHERE carrier.id = '%s';", carrierID)
	query := fmt.Sprintf(sqlGetCarrierDataByID, carrierID)
	row := sctrack.Db.QueryRow(query)

	var err error
	var carrier models.Carrier

	if err := row.Scan(&carrier.ID, &carrier.Name, &carrier.Telephone); err != nil {
		sctrack.Log.Error("GetCarrierDataByID fail", slog.String("Error", err.Error()))
		return nil, err
	}

	// get the carrier's todos
	var todos []models.Todos
	if todos, err = GetTodosByCarrierID(carrierID); err != nil {
		sctrack.Log.Error("GetCarrierDataByID fail", slog.String("Error", err.Error()))
		return nil, err
	}

	// get the carrier's journals
	var journals []models.Journal
	if journals, err = GetJournalsByCarrierID(carrierID); err != nil {
		sctrack.Log.Error("GetCarrierDataByID fail", slog.String("Error", err.Error()))
		return nil, err
	}

	// return the response
	carrierData := models.CarrierData{
		Carrier: carrier,
		Todos:   todos,
		Journal: journals,
	}

	// return the carriers info
	return &carrierData, nil
}

// DeleteCarrierByID deletes a carrier based on the ID
func DeleteCarrierByID(carrierID string) error {

	// delete carrier and data
	// TODO: remove  stmt := fmt.Sprintf("DELETE FROM carrier WHERE id = '%s';", carrierID)
	stmt := fmt.Sprintf(sqlDeleteCarrierByID, carrierID)

	_, err := sctrack.Db.Exec(stmt)
	if err != nil {
		sctrack.Log.Warn("Failed to delete a todo by id", slog.String("carrier id", carrierID),
			slog.String("Error", err.Error()))
		return err
	}
	return nil
}

// AddCarrier adds a new carrier
func AddCarrier(carrier models.Carrier) error {

	// add the carrier
	stmt := fmt.Sprintf(sqlAddCarrier, carrier.ID, carrier.Name, carrier.Telephone)
	_, err := sctrack.Db.Exec(stmt)
	if err != nil {
		sctrack.Log.Warn("Failed to add carrier", slog.String("Error", err.Error()))
		return err
	}
	return nil
}
