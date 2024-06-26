package mongodbutil

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
)

// ConnectToMongoDB connects to the MongoDB server.
func ConnectToMongoDB(uri string, timeout time.Duration) error {
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		cancel()
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		client.Disconnect(ctx)
		cancel()
		return err
	}
	return nil
}

// SetDatabaseAndCollection sets the database and collection to operate on.
func SetDatabaseAndCollection(dbName, collName string) {
	database = client.Database(dbName)
	collection = database.Collection(collName)
}

// FetchDocumentByID fetches a document by its ID.
func FetchDocumentByID(userID string) (bson.M, error) {
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no document found with that ID")
		}
		return nil, err
	}
	return result, nil
}

// UpdateDocument updates a document in the collection.
func UpdateDocument(userID string, update bson.M) (int64, error) {
	filter := bson.M{"_id": userID}
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

// DisconnectFromMongoDB disconnects from the MongoDB server.
func DisconnectFromMongoDB() {
	client.Disconnect(ctx)
	cancel()
}
