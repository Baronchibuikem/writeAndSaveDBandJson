package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/user"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToMongoDB connects to the MongoDB instance and returns the client and database.
func ConnectToMongoDB(uri, databaseName string) (*mongo.Client, *mongo.Database) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, client.Database(databaseName)
}

// GetHomeDirectory returns the current user's home directory.
func GetHomeDirectory() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

// CreateDirectory creates a directory if it doesn't exist.
func CreateDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}
