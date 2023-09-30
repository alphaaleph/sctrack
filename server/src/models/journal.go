package models

import uuid2 "github.com/google/uuid"

// Journal is the journal table model
type Journal struct {
	UUID     uuid2.UUID `json:"uuid"`
	Index    int        `json:"index"`
	Event    string     `json:"event"`
	TodoUUID uuid2.UUID `json:"todo_uuid"`
}

// JournalData includes the carrier name in the response
type JournalData struct {
	CarrierID   string     `json:"carrierID"`
	CarrierName string     `json:"carrierName"`
	TodoUUID    uuid2.UUID `json:"todo_uuid"`
	UUID        uuid2.UUID `json:"uuid"`
	Index       int        `json:"index"`
	Event       string     `json:"event"`
}
