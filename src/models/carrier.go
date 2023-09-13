package models

// Carrier is the carrier table model
type Carrier struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// CarrierData is the response containing all the data pertaining a carrier
type CarrierData struct {
	Carrier Carrier   `json:"carrier"`
	Todos   []Todos   `json:"todos"`
	Journal []Journal `json:"journal"`
}
