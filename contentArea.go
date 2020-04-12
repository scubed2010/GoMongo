package main

import (
	"context"
	"log"
	"time"

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

func (c ContentArea) getAll(collection string) []*ContentArea {
	var results []*ContentArea
	cur := getAllCursor(collection)

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

func (c ContentArea) getByID(collection string, id string) *ContentArea {
	var result *ContentArea

	getByIDCursor(collection, id).Decode(&result)

	return result
}

func (c ContentArea) create(collection string, contentArea ContentArea) *mongo.InsertOneResult {
	col := getCollection(collection)

	contentArea.ID = primitive.NewObjectID()
	insertResult, err := col.InsertOne(context.TODO(), contentArea)
	if err != nil {
		log.Fatal(err)
	}

	return insertResult
}

func (c ContentArea) update(collection string, contentArea ContentArea) *mongo.UpdateResult {
	col := getCollection(collection)

	updateResult, err := col.ReplaceOne(context.TODO(), bson.M{"_id": contentArea.ID}, contentArea)
	if err != nil {
		log.Fatal(err)
	}

	return updateResult
}

func (c ContentArea) delete(collection string, id string) *mongo.DeleteResult {
	col := getCollection(collection)

	objID, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := col.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
	}

	return deleteResult
}
