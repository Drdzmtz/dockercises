package main

import (
	"encoding/xml"
	"net/http"
)

type people struct {
	XMLName       xml.Name  `xml:"people"`
	ListaPersonas []persona `xml:"person"`
}

type persona struct {
	ID          float64 `xml:"id" bson:"ID"`
	FirstName   string  `xml:"first_name" bson:"FirstName"`
	LastName    string  `xml:"last_name" bson:"LastName"`
	Company     string  `xml:"company" bson:"Company"`
	Email       string  `xml:"email" bson:"Email"`
	IPAddress   string  `xml:"ip_address" bson:"IPaddress"`
	PhoneNumber string  `xml:"phone_number" bson:"PhoneNumber"`
}

func (a persona) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
