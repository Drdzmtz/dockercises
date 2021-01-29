package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var r *chi.Mux
var client *mongo.Client
var err error

var personIDKey = "Key"

func init() {
	r = chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	client, err = getMongoClient()

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func main() {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dockercise1; Contenedor 4"))
	})

	r.Route("/people", func(r chi.Router) {
		r.Get("/", getAllPeople)

		r.Route("/{PersonID}", func(r chi.Router) {
			r.Use(itemContext)

			r.Get("/", getPerson)
		})
	})

	http.ListenAndServe(":7777", r)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	personID := r.Context().Value(personIDKey).(int)

	database := client.Database("dockercises1")
	peopleCollection := database.Collection("people")

	queryCursor, _ := peopleCollection.Find(context.TODO(), bson.M{"ID": personID})

	var results []bson.M
	if err := queryCursor.All(context.TODO(), &results); err != nil || len(results) != 1 {
		if len(results) != 1 {
			render.Render(w, r, errorRenderer(fmt.Errorf("Not found in database")))
			return
		}
	}

	render.RenderList(w, r, peopleListResponse(results))
}

func getAllPeople(w http.ResponseWriter, r *http.Request) {
	database := client.Database("dockercises1")
	peopleCollection := database.Collection("people")

	queryCursor, err := peopleCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err := queryCursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	render.RenderList(w, r, peopleListResponse(results))
}

func peopleListResponse(p []bson.M) []render.Renderer {
	var ppl people
	list := []render.Renderer{}

	for _, i := range p {

		ppl.ListaPersonas = append(ppl.ListaPersonas, persona{
			ID:          i["ID"].(float64),
			FirstName:   i["FirstName"].(string),
			LastName:    i["LastName"].(string),
			Company:     i["Company"].(string),
			Email:       i["Email"].(string),
			IPAddress:   i["IPaddress"].(string),
			PhoneNumber: i["PhoneNumber"].(string),
		})
	}

	for _, person := range ppl.ListaPersonas {
		list = append(list, person)
	}

	return list
}

func itemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		PersonID := chi.URLParam(r, "PersonID")

		id, err := strconv.Atoi(PersonID)
		if err != nil {
			render.Render(w, r, errorRenderer(fmt.Errorf("Invalid item ID")))
			return
		}

		ctx := context.WithValue(r.Context(), personIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type errorResponse struct {
	Err        error  `json:"err"`
	StatusCode int    `json:"statuscode"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

func errorRenderer(err error) *errorResponse {
	return &errorResponse{
		Err:        err,
		StatusCode: 400,
		StatusText: "Bad Request",
		Message:    err.Error(),
	}
}

func (e *errorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}
