package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Coach stuff
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

func (c Coach) getAll(collection string) []*Coach {
	var results []*Coach
	cur := getAllCursor(collection)

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

func (c Coach) getByID(collection string, id string) *Coach {
	var result *Coach

	getByIDCursor(collection, id).Decode(&result)

	return result
}

func (c Coach) create(collection string, coach Coach) *mongo.InsertOneResult {
	col := getCollection(collection)

	coach.ID = primitive.NewObjectID()
	insertResult, err := col.InsertOne(context.TODO(), coach)
	if err != nil {
		log.Fatal(err)
	}

	return insertResult
}

func (c Coach) update(collection string, coach Coach) *mongo.UpdateResult {
	col := getCollection(collection)

	updateResult, err := col.ReplaceOne(context.TODO(), bson.M{"_id": coach.ID}, coach)
	if err != nil {
		log.Fatal(err)
	}

	return updateResult
}

func (c Coach) delete(collection string, id string) *mongo.DeleteResult {
	col := getCollection(collection)

	objID, _ := primitive.ObjectIDFromHex(id)
	deleteResult, err := col.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
	}

	return deleteResult
}
