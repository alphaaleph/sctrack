package models

// Journal is the journal table model
type Journal struct {
	UUID     string `json:"uuid"`
	Index    int    `json:"index"`
	Event    `json:"event"`
	TodoUUID string `json:"todo_uuid"`
}

// JournalData includes the carrier name in the response
type JournalData struct {
	CarrierID   string `json:"carrierID"`
	CarrierName string `json:"carrierName"`
	TodoUUID    string `json:"todo_uuid"`
	UUID        string `json:"uuid"`
	Index       int    `json:"index"`
	Event       `json:"event"`
}
