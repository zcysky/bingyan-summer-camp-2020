package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QueryForm struct {
	UserID	string	`json:"user_id"`
	Limit 	int 	`json:"limit"`
	Page  	int 	`json:"page"`
}

func QueryAll(limit int, page int) (users []User, total int64, err error) {
	err = TestConnection()
	if err != nil {
		return nil, -1, err
	}

	collection := client.Database("users").Collection("users")
	total, err = collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, -1, err
	}

	queryOptions := options.Find().SetLimit(int64(limit))
	queryOptions.SetSkip(int64((page - 1) * limit))

	queryResults, err := collection.Find(context.TODO(), bson.M{}, queryOptions)
	if err != nil {
		return nil, total, err
	}

	err = queryResults.All(context.TODO(), &users)
	if err != nil {
		return users, total, err
	}
	return users, total, nil
}

func QueryOne(id string) (user User, err error) {
	err = TestConnection()
	if err != nil {
		return user, err
	}

	collection := client.Database("users").Collection("users")
	userObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	filter := bson.M{"_id":	userObjectID}

	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}