package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Data map[string]interface{}

func JsonToDB() {
	const (
		mongoURI     = "mongodb://localhost:27017"
		databaseName = "odeserverlocaltestdb"
	)

	client, database := ConnectToMongoDB(mongoURI, databaseName)
	defer client.Disconnect(context.TODO())

	homeDir := GetHomeDirectory()
	dir := filepath.Join(homeDir, "Downloads", "mongodb_backup")

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".json" {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			var data []Data
			decoder := json.NewDecoder(file)
			if err = decoder.Decode(&data); err != nil {
				return err
			}

			collectionName := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
			collection := database.Collection(collectionName)

			for _, item := range data {
				if _, err = collection.InsertOne(context.TODO(), item); err != nil {
					return err
				}
			}

			fmt.Printf("Data from %s successfully inserted into %s collection\n", path, collectionName)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
