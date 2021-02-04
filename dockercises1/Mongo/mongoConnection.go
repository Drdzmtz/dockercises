package mongoconnection

import (
	"context"
	"fmt"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db string = os.Getenv("MONGO_INITDB_DATABASE")
var dbUserName string = os.Getenv("ME_CONFIG_MONGODB_ADMINUSERNAME")
var dbPassword string = os.Getenv("ME_CONFIG_MONGODB_ADMINPASSWORD")
var dbHost string = os.Getenv("MONGO_HOST")
var dbPort string = os.Getenv("MONGO_PORT")
var connectionString string = fmt.Sprintf("mongodb://%s:%s@%s:%s/?connect=direct", dbUserName, dbPassword, dbHost, dbPort)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

// GetMongoClient returns a client instance to the mongoDatabase
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(connectionString)
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
