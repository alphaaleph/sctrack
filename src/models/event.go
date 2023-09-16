package models

import "time"

type Event struct {
	TimeStamp time.Time `json:"time"`
	Action    `json:"action"`
}
