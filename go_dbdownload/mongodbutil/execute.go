package mongodbutil

import (
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ExecuteMongoUtil() {
	// Connect to MongoDB
	if err := ConnectToMongoDB("mongodb://localhost:27017/", 10*time.Second); err != nil {
		log.Fatal(err)
	}
	defer DisconnectFromMongoDB()

	// Set the database and collection
	SetDatabaseAndCollection("yourdbname", "users")

	// Perform operations
	userID := "userIdtoBeUpdated"

	result, err := FetchDocumentByID(userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found document: %v\n", result)

	update := bson.M{
		"$set": bson.M{
			"field.toUpdate": []interface{}{},
		},
	}

	modifiedCount, err := UpdateDocument(userID, update)
	if err != nil {
		log.Fatal(err)
	}

	if modifiedCount == 0 {
		fmt.Println("No document found with that ID")
	} else {
		fmt.Printf("Successfully updated %v document(s)\n", modifiedCount)
	}
}
