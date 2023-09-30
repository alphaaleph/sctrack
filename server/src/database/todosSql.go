package database

const (
	sqlGetTodos            = "SELECT uuid, created, description, completed, carrier_id, action FROM todos;"
	sqlGetTodosByCarrierID = "SELECT uuid, created, description, completed, carrier_id, action " +
		"FROM todos WHERE carrier_id = '%s';"
	sqlGetTodosByUUID         = "SELECT uuid, created, description, completed, carrier_id, action FROM todos WHERE uuid = '%s';"
	sqlDeleteTodosByCarrierID = "DELETE FROM todos WHERE carrier_id = '%s';"
	sqlDeleteTodosByUUID      = "DELETE FROM todos WHERE uuid = '%s';"
	sqlAddTodos               = "INSERT INTO todos (uuid, created, description, completed, carrier_id, action) " +
		"VALUES ('%s', '%s', '%s', '%v', '%s', '%s');"
	sqlPatchTodosCompleted = "UPDATE todos set completed = '%v' WHERE uuid = '%s';"
)
