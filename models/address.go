package models

//Address is used to hold the address details of the user
type Address struct {
	HouseNumber string `json:"housenumber,omitempty"`
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
	ZipCode     string `json:"address,omitempty"`
}
