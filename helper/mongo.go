package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	fmt.Println("init called")
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin123@mongo.default.svc.cluster.local:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	fmt.Println("connection established")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("admin").Collection("scalecollection")
	fmt.Println("collection initiated")
}

type TestData struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

// Insert documents into MongoDB collection
func InsertMongoData() error {

	fmt.Println("insertMongoData Called")
	testData := &TestData{
		ID:   primitive.NewObjectID(),
		Name: "test",
	}
	_, err := collection.InsertOne(ctx, testData)
	fmt.Println("insertMongoData Inserted Data")
	return err

}

// Delete all documents from MongoDB collection
func DeleteMongoData() error {
	fmt.Println("deleteMongoData called")
	// delete where name = "test"
	_, err := collection.DeleteMany(ctx, bson.M{"name": "test"})

	return err
}
