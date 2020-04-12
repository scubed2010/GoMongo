package models

import (
	"context"
	"log"
	"time"

	"../data"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ContentArea stuff
type ContentArea struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedOn   time.Time          `json:"CreatedOn"`
	CreatedBy   string             `json:"CreatedBy"`
	ModifiedOn  time.Time          `json:"ModifiedOn"`
	ModifiedBy  string             `json:"ModifiedBy"`
	Key         string             `json:"Key"`
	Description string             `json:"Description"`
	Content     string             `json:"Content"`
}

// GetAll - Gets all documents for the collection
func (c ContentArea) GetAll(collection string) []*ContentArea {
	var results []*ContentArea
	cur := data.GetAllCursor(collection)

	// Loop Cursor
	for cur.Next(context.TODO()) {
		var elem ContentArea
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}

// GetByID - Returns a single document based on _id
func (c ContentArea) GetByID(collection string, id string) *ContentArea {
	var result *ContentArea

	data.GetByIDCursor(collection, id).Decode(&result)

	return result
}

// Create - Inserts a single document
func (c ContentArea) Create(collection string, contentArea ContentArea) *mongo.InsertOneResult {
	col := data.GetCollection(collection)

	contentArea.ID = primitive.NewObjectID()
	insertResult, err := col.InsertOne(context.TODO(), contentArea)
	if err != nil {
		log.Fatal(err)
	}

	return insertResult
}

// Update - Replaces a document based on _id
func (c ContentArea) Update(collection string, contentArea ContentArea) *mongo.UpdateResult {
	col := data.GetCollection(collection)

	updateResult, err := col.ReplaceOne(context.TODO(), bson.M{"_id": contentArea.ID}, contentArea)
	if err != nil {
		log.Fatal(err)
	}

	return updateResult
}

// Delete - Removes a document
func (c ContentArea) Delete(collection string, id string) *mongo.DeleteResult {
	col := data.GetCollection(collection)

	objID, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := col.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
	}

	return deleteResult
}
