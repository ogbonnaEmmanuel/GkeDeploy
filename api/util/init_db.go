package util

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckAndCreateDatabase(client *mongo.Client, dbName string, collections []string) error {
	// Check if the database exists
	db := client.Database(dbName)
	for _, collectionName := range collections {
		// Check if the collection exists
		collectionsList, err := db.ListCollectionNames(context.TODO(), bson.M{"name": collectionName})
		if err != nil {
			return err
		}

		// Create the collection if it does not exist
		if len(collectionsList) == 0 {
			log.Printf("Creating collection: %s", collectionName)
			err := db.CreateCollection(context.TODO(), collectionName)
			if err != nil {
				return err
			}
		} else {
			log.Printf("Collection %s already exists", collectionName)
		}
	}
	return nil
}