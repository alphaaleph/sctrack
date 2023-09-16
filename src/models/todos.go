package models

import (
	"time"
)

// Todos is the todos table model
type Todos struct {
	UUID        string    `json:"uuid"`
	Created     time.Time `json:"created"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CarrierID   string    `json:"carrierID"`
	Action      `json:"action"`
}

// TodosData includes the carrier name in the response
type TodosData struct {
	CarrierID   string    `json:"carrierID"`
	CarrierName string    `json:"carrierName"`
	UUID        string    `json:"uuid"`
	Created     time.Time `json:"created"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Action      `json:"action"`
}

// TodosAdd used to add new todos
type TodosAdd struct {
	CarrierID   string `json:"carrierID"`
	Description string `json:"description"`
	Action      `json:"action"`
}
