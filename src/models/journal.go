package models

// Journal is the journal table model
type Journal struct {
	UUID      int    `json:"uuid"`
	Event     string `json:"event"`
	CarrierID string `json:"carrierID"`
}

// JournalData includes the carrier name in the response
type JournalData struct {
	CarrierID string `json:"carrierID"`
	Carrier   string `json:"carrier"`
	UUID      int    `json:"uuid"`
	Event     string `json:"event"`
}
