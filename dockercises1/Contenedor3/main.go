package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var people people
	xmlToPeople("people.xml", &people)

	client, err := getMongoClient()
	defer client.Disconnect(context.TODO())

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
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

	// fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func xmlToPeople(fileName string, people *people) {
	xmlFile, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Errror: %v", err)
	}
	defer xmlFile.Close()

	xmlBytes, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(xmlBytes, people)
}
