package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RemoveForm struct {
	UserID	string	`json:"user_id"`
}

func Remove(id string) (err error) {
	err = TestConnection()
	if err != nil {
		return err
	}

	collection := client.Database("users").Collection("users")
	userObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id":	userObjectID}

	_, err = collection.DeleteOne(context.TODO(), filter)
	return err
}