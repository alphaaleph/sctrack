package database

const (
	sqlGetCarriers        = "SELECT * FROM  carrier;"
	sqlGetCarrierDataByID = "SELECT id, name, telephone FROM carrier WHERE id = '%s';"
	sqlDeleteCarrierByID  = "DELETE FROM carrier WHERE id = '%s';"
	sqlAddCarrier         = "INSERT INTO carrier (id, name, telephone) VALUES ('%s', '%s', '%s');"
)
