package models

type ActionTable struct {
	Action `json:"action"`
}

type Action string

var (
	Other     Action = "other"
	Completed Action = "completed"
	Delivery  Action = "delivery"
	Pickup    Action = "pickup"
	Refuel    Action = "refuel"
)
