package models

import (
	"encoding/xml"
	"net/http"
)

//People is a Struct for the XML as a whole; a group of persons
type People struct {
	XMLName       xml.Name  `xml:"people"`
	ListaPersonas []Persona `xml:"person"`
}

//Persona is a Struct for a single person
type Persona struct {
	ID          float64 `xml:"id" bson:"ID"`
	FirstName   string  `xml:"first_name" bson:"FirstName"`
	LastName    string  `xml:"last_name" bson:"LastName"`
	Company     string  `xml:"company" bson:"Company"`
	Email       string  `xml:"email" bson:"Email"`
	IPAddress   string  `xml:"ip_address" bson:"IPaddress"`
	PhoneNumber string  `xml:"phone_number" bson:"PhoneNumber"`
}

//Render is a function to satisfy the renderer interface
func (a Persona) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
