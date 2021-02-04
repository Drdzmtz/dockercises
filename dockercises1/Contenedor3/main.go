package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	models "github.com/Drdzmtz/dockercises/models"
	mongoconnection "github.com/Drdzmtz/dockercises/mongoconnection"
)

func main() {
	var people models.People
	xmlToPeople("people.xml", &people)

	client, err := mongoconnection.GetMongoClient()
	defer client.Disconnect(context.TODO())

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	database := client.Database("dockercises1")
	peopleCollection := database.Collection("people")

	var insertSlice []interface{}
	for _, i := range people.ListaPersonas {
		insertSlice = append(insertSlice, i)

	}

	_, err = peopleCollection.InsertMany(context.TODO(), insertSlice)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func xmlToPeople(fileName string, people *models.People) {
	xmlFile, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Errror: %v", err)
	}
	defer xmlFile.Close()

	xmlBytes, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(xmlBytes, people)
}
