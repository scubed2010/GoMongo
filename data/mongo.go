package data

import (
	"context"
	"log"

	"../config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// HealthCheck - Checks if MongoDB is up
func HealthCheck() string {
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

// GetCollection - Gets the collection to be executed upon
func GetCollection(collection string) *mongo.Collection {
	configuration := config.GetConfiguration()

	client, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(configuration.MongoDBDatabase).Collection(collection)
}

// GetAllCursor - Returns cursor of all documents in a collection
func GetAllCursor(collection string) *mongo.Cursor {
	col := GetCollection(collection)

	cur, err := col.Find(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	return cur
}

// GetByIDCursor - Returns cursor with a single document based on _id
func GetByIDCursor(collection string, id string) *mongo.SingleResult {
	col := GetCollection(collection)

	objID, _ := primitive.ObjectIDFromHex(id)
	return col.FindOne(context.TODO(), bson.M{"_id": objID})
}

func connectToMongoDB() (*mongo.Client, error) {
	configuration := config.GetConfiguration()

	/// Set client options
	clientOptions := options.Client().ApplyURI(configuration.MongoDBConnectionString)

	// Connect to MongoDB
	return mongo.Connect(context.TODO(), clientOptions)
}
