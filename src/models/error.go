package models

// Error is used for error responses
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
