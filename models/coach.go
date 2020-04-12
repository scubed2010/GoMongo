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

// Coach - IYWT Employee
type Coach struct {
	ID                 primitive.ObjectID `bson:"_id"`
	CreatedOn          time.Time          `json:"CreatedOn"`
	CreatedBy          string             `json:"CreatedBy"`
	ModifiedOn         time.Time          `json:"ModifiedOn"`
	ModifiedBy         string             `json:"ModifiedBy"`
	DisplayName        string             `json:"DisplayName"`
	Email              string             `json:"Email"`
	Initials           string             `json:"Initials"`
	OutOfOfficeMessage string             `json:"OutOfOfficeMessage"`
	DisplayOOOMessage  bool               `json:"DisplayOOOMessage"`
	CountryAssignments []string           `json:"CountryAssignments"`
}

// GetAll - Gets all documents for the collection
func (c Coach) GetAll(collection string) []*Coach {
	var results []*Coach
	cur := data.GetAllCursor(collection)

	// Loop Cursor
	for cur.Next(context.TODO()) {
		var elem Coach
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
func (c Coach) GetByID(collection string, id string) *Coach {
	var result *Coach

	data.GetByIDCursor(collection, id).Decode(&result)

	return result
}

// Create - Inserts a single document
func (c Coach) Create(collection string, coach Coach) *mongo.InsertOneResult {
	col := data.GetCollection(collection)

	coach.ID = primitive.NewObjectID()
	insertResult, err := col.InsertOne(context.TODO(), coach)
	if err != nil {
		log.Fatal(err)
	}

	return insertResult
}

// Update - Replaces a document based on _id
func (c Coach) Update(collection string, coach Coach) *mongo.UpdateResult {
	col := data.GetCollection(collection)

	updateResult, err := col.ReplaceOne(context.TODO(), bson.M{"_id": coach.ID}, coach)
	if err != nil {
		log.Fatal(err)
	}

	return updateResult
}

// Delete - Removes a document
func (c Coach) Delete(collection string, id string) *mongo.DeleteResult {
	col := data.GetCollection(collection)

	objID, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := col.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
	}

	return deleteResult
}
