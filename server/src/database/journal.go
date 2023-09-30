package database

import (
	"fmt"
	"github.com/alphaaleph/sctrack/server"
	models2 "github.com/alphaaleph/sctrack/server/src/models"
	"golang.org/x/exp/slog"
	"time"
)

// getJournalList gets a slice of journals
func getJournalList(query string) ([]models2.Journal, error) {

	rows, err := server.Db.Query(query)
	if err != nil {
		server.Log.Warn("getJournalList fail", slog.String("Error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	// read all the journal entries
	var journals []models2.Journal

	for rows.Next() {
		var journal models2.Journal
		rows.Scan(&journal.UUID, &journal.Index, &journal.Event, &journal.TodoUUID)
		journals = append(journals, journal)
	}

	// return the journals
	return journals, nil
}

// GetJournals reads all the data from the journals table
func GetJournals() ([]models2.Journal, error) {
	return getJournalList(sqlGetJournals)
}

// GetJournalByUUID returns the journal associated with the uuid
func GetJournalByUUID(uuid string) (*models2.Journal, error) {

	query := fmt.Sprintf(sqlGetJournalByUUID, uuid)
	row := server.Db.QueryRow(query)

	var journal models2.Journal
	if err := row.Scan(&journal.UUID, &journal.Index, &journal.Event, &journal.TodoUUID); err != nil {
		server.Log.Error("GetJournalByUUID fail", slog.String("uuid", uuid), slog.String("Error", err.Error()))
		return nil, err
	}

	// return the information
	return &journal, nil
}

// GetJournalsByCarrierID returns the journals associated with a carrier id
func GetJournalsByCarrierID(carrierID string) ([]models2.Journal, error) {

	query := fmt.Sprintf(sqlGetJournalsByCarrierID, carrierID)
	return getJournalList(query)
}

// DeleteJournalByUUID deletes all journal entries based on uuid
func DeleteJournalByUUID(uuid string) (int64, error) {

	stmt := fmt.Sprintf(sqlDeleteJournalByUUID, uuid)

	r, err := server.Db.Exec(stmt)
	if err != nil {
		server.Log.Warn("DeleteJournalByUUID fail", slog.String("uuid", uuid), slog.String("Error", err.Error()))
		return 0, err
	}

	// return affected rows
	rowsAffected, _ := r.RowsAffected()
	return rowsAffected, nil
}

// DeleteJournalByUUIDIndex deletes a journal entry based on uuid and index
func DeleteJournalByUUIDIndex(index uint, uuid string) (int64, error) {

	stmt := fmt.Sprintf(sqlDeleteJournalByUUIDIndex, uuid, index)

	r, err := server.Db.Exec(stmt)
	if err != nil {
		server.Log.Warn("DeleteJournalByUUIDIndex fail", slog.String("uuid", uuid), slog.Int("index", int(index)),
			slog.String("Error", err.Error()))
		return 0, err
	}

	// return affected rows
	rowsAffected, _ := r.RowsAffected()
	return rowsAffected, nil
}

// AddJournal adds a new journal
func AddJournal(journal models2.Journal) error {

	// add the journal
	stmt := fmt.Sprintf(sqlAddJournal, journal.UUID, journal.Event, journal.TodoUUID)
	_, err := server.Db.Exec(stmt)
	if err != nil {
		server.Log.Warn("AddJournal fail", slog.String("Error", err.Error()))
		return err
	}
	return nil
}

// newEvent creates a new event to track things like starting and finished a todos task, more things can be added to
// the event string to track additional info.  This is just a small program example, so keeping is simple. It returns
// an event
func newEvent(action models2.Action) string {
	ts := time.Now()
	event := fmt.Sprintf("{\"time\":\"%v\",\"action\":\"%s\"}", ts, action)
	return event
}
