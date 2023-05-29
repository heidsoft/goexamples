package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Define the mongodb client URL
	var uri = "mongodb://localhost:27017"

	// Establish the connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Create go routine to defer the closure
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("Employee").Collection("scoreCollection")
	docs := []interface{}{
		bson.D{{"name", "Alley"}, {"score", 7.5}},
		bson.D{{"name", "Bob"}, {"score", 8.5}},
		bson.D{{"name", "Carry"}, {"score", 6.8}},
		bson.D{{"name", "Daniel"}, {"score", 5.5}},
		bson.D{{"name", "Danish"}, {"score", 4.8}},
		bson.D{{"name", "Era"}, {"score", 9.2}},
		bson.D{{"name", "Hush"}, {"score", 10}},
		bson.D{{"name", "Halley"}, {"score", 3.6}},
		bson.D{{"name", "John"}, {"score", 7.5}},
	}

	// insertMany
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}
	// end insertMany

	// When you run this file, it should print:
	// Document inserted with ID: ObjectID("...")
	for _, id := range result.InsertedIDs {
		fmt.Printf("\t%s\n", id)
	}

	filter := bson.D{{Key: "score", Value: bson.D{{"$lte", 7.5}}}}
	update := bson.D{{Key: "$set", Value: bson.D{{"type", "bad"}}}}

	// update Many
	update_result, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	// end update Many

	// When you run this file for the first time, it should print:
	// Number of documents replaced:
	fmt.Printf("Documents updated: %v\n", update_result.ModifiedCount)

}
