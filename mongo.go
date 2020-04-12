package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func healthCheck() string {
	client, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return "Connected to MongoDB!"
}

func connectToMongoDB() (*mongo.Client, error) {
	/// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	return mongo.Connect(context.TODO(), clientOptions)
}

func getCollection(collection string) *mongo.Collection {
	client, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("world").Collection(collection)
}

func getAllCursor(collection string) *mongo.Cursor {
	col := getCollection(collection)

	cur, err := col.Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	return cur
}

func getByIDCursor(collection string, id string) *mongo.SingleResult {
	col := getCollection(collection)

	objID, _ := primitive.ObjectIDFromHex(id)
	return col.FindOne(context.TODO(), bson.M{"_id": objID})
}
