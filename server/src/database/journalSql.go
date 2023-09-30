package database

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
