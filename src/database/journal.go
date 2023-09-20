package database

import (
	"fmt"
	"github.com/alphaaleph/sctrack"
	"github.com/alphaaleph/sctrack/src/models"
	"golang.org/x/exp/slog"
	"time"
)

const (
	sqlGetJournals            = "SELECT uuid, index, event, todo_uuid FROM journal;"
	sqlGetJournalByUUID       = "SELECT uuid, index, event, todo_uuid FROM journal WHERE uuid = '%s';"
	sqlGetJournalsByCarrierID = "SELECT j.uuid AS journal_uuid, j.index AS journal_index, j.event AS journal_event, " +
		"j.todo_uuid as journal_todo_uuid FROM journal j INNER JOIN todos t ON j.todo_uuid = t.uuid " +
		"INNER JOIN carrier c ON t.carrier_id = c.id WHERE c.id = '%s';"
	sqlDeleteJournalByUUID      = "DELETE FROM journal WHERE journal.uuid = '%s';"
	sqlDeleteJournalByUUIDIndex = "DELETE FROM journal WHERE uuid = '%s' and index = '%d';"
	sqlAddJournal               = "INSERT INTO journal (uuid, event, todo_uuid) VALUES ('%s', '%s', '%s');"
)

// getJournalList gets a slice of journals
func getJournalList(query string) ([]models.Journal, error) {

	rows, err := sctrack.Db.Query(query)
	if err != nil {
		sctrack.Log.Warn("Failed to get all journals", slog.String("Error", err.Error()))
		return nil, err
	}
	defer rows.Close()

	// read all the journal entries
	var journals []models.Journal

	for rows.Next() {
		var journal models.Journal
		rows.Scan(&journal.UUID, &journal.Index, &journal.Event, &journal.TodoUUID)
		journals = append(journals, journal)
	}

	// return the journals
	return journals, nil
}

// GetJournals reads all the data from the journals table
func GetJournals() ([]models.Journal, error) {
	// TODO: remove query := "SELECT uuid, index, event, todo_uuid FROM journal;"
	return getJournalList(sqlGetActions)
}

// GetJournalByUUID returns the journal associated with the uuid
func GetJournalByUUID(uuid string) (*models.Journal, error) {

	// TODO: remove query := fmt.Sprintf("SELECT uuid, index, event, todo_uuid FROM journal WHERE uuid = '%s';", uuid)
	query := fmt.Sprintf(sqlGetJournalByUUID, uuid)
	row := sctrack.Db.QueryRow(query)

	var journal models.Journal
	if err := row.Scan(&journal.UUID, &journal.Index, &journal.Event, &journal.TodoUUID); err != nil {
		sctrack.Log.Error("Fail to get todo", slog.String("uuid", uuid), slog.String("Error", err.Error()))
		return nil, err
	}

	// return the information
	return &journal, nil
}

// GetJournalsByCarrierID returns the journals associated with a carrier id
func GetJournalsByCarrierID(carrierID string) ([]models.Journal, error) {

	/* TODO: reomvequery := fmt.Sprintf("SELECT j.uuid AS journal_uuid, j.index AS journal_index, j.event AS journal_event, "+
	"j.todo_uuid as journal_todo_uuid FROM journal j INNER JOIN todos t ON j.todo_uuid = t.uuid "+
	"INNER JOIN carrier c ON t.carrier_id = c.id WHERE c.id = '%s';", carrierID) */
	query := fmt.Sprintf(sqlGetJournalsByCarrierID, carrierID)
	return getJournalList(query)
}

// DeleteJournalByUUID deletes all journal entries based on uuid
func DeleteJournalByUUID(uuid string) error {

	// TODO: remove stmt := fmt.Sprintf("DELETE FROM journal WHERE journal.uuid = '%s';", uuid)
	stmt := fmt.Sprintf(sqlDeleteJournalByUUID, uuid)

	_, err := sctrack.Db.Exec(stmt)
	if err != nil {
		sctrack.Log.Warn("Failed to delete a journal", slog.String("uuid", uuid), slog.String("Error", err.Error()))
		return err
	}
	return nil
}

// DeleteJournalByUUIDIndex deletes a journal entry based on uuid and index
func DeleteJournalByUUIDIndex(index uint, uuid string) error {

	// TODO: remove stmt := fmt.Sprintf("DELETE FROM journal WHERE uuid = '%s' and index = '%d';", uuid, index)
	stmt := fmt.Sprintf(sqlDeleteJournalByUUIDIndex, uuid, index)

	_, err := sctrack.Db.Exec(stmt)
	if err != nil {
		sctrack.Log.Warn("Failed to delete a journal", slog.String("uuid", uuid), slog.Int("index", int(index)),
			slog.String("Error", err.Error()))
		return err
	}
	return nil
}

// AddJournal adds a new journal
func AddJournal(journal models.Journal) error {

	// add the journal
	stmt := fmt.Sprintf(sqlAddJournal, journal.UUID, journal.Event, journal.TodoUUID)
	_, err := sctrack.Db.Exec(stmt)
	if err != nil {
		sctrack.Log.Warn("Failed to add journal", slog.String("Error", err.Error()))
		return err
	}
	return nil
}

// newEvent creates a new event to track things like starting and finished a todos task, more things can be added to
// the event string to track additional info.  This is just a small program example, so keeping is simple. It returns
// an event
func newEvent(action models.Action) models.Event {
	var event models.Event
	event.TimeStamp = time.Now()
	event.Action = action
	return event
}
