package models

// Todos is the todos table model
type Todos struct {
	UUID      int    `json:"uuid"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CarrierID string `json:"carrierID"`
}

// TodosData includes the carrier name in the response
type TodosData struct {
	CarrierID string `json:"carrierID"`
	Carrier   string `json:"carrier"`
	UUID      int    `json:"uuid"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
