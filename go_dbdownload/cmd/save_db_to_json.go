package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DBToJson() {
	const (
		mongoURI     = "mongodb://localhost:27017"
		databaseName = "odeserverlocaltestdb"
	)

	client, db := ConnectToMongoDB(mongoURI, databaseName)
	defer client.Disconnect(context.TODO())

	collections, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	homeDir := GetHomeDirectory()
	outputDir := filepath.Join(homeDir, "Documents", "mongodb_backup")
	CreateDirectory(outputDir)

	for _, collectionName := range collections {
		collection := db.Collection(collectionName)
		cursor, err := collection.Find(context.TODO(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.TODO())

		var data []bson.M
		if err = cursor.All(context.TODO(), &data); err != nil {
			log.Fatal(err)
		}

		for _, item := range data {
			if id, ok := item["_id"].(primitive.ObjectID); ok {
				item["_id"] = id.Hex()
			}
		}

		filePath := filepath.Join(outputDir, fmt.Sprintf("%s.json", collectionName))
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if err := json.NewEncoder(file).Encode(data); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Data for collection '%s' saved to '%s'\n", collectionName, filePath)
	}

	fmt.Println("Backup completed successfully.")
}
