/*Day3 Start learning mongodb*/
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Username string
	Password string
	RegTime  string
	Email    string
}

func main() {
	// Set up client option
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to database
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Print(err)
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected successfully!")

	collection := client.Database("test").Collection("TestCollection")

	//Insert a record

	info := Student{"TobyZT", "19260817", "20200715", ""}
	id := InsertUser(info, collection)
	fmt.Println(id)

	FindUsername("TobyZT", collection)

	objID, _ := primitive.ObjectIDFromHex("5f0f007701736b547d88e186")
	UpdatePassword(objID, "23336666", collection)
}

func InsertUser(info Student, collection *mongo.Collection) interface{} {
	res, err := collection.InsertOne(context.TODO(), info)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return res.InsertedID
}

func FindUsername(username string, collection *mongo.Collection) {
	filter := bson.M{"username": username}
	var res Student
	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Matched!")
	fmt.Println(res)
}

func UpdatePassword(userID interface{}, newPassword string, collection *mongo.Collection) {
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"password": newPassword}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated successfully")
	}
}
