package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const colNameUser = "user"
var colUser *mongo.Collection

func initModelUser() {
	colUser = MongoDatabase.Collection(colNameUser)
}

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Phone    string             `bson:"phone" json:"phone"`
	Email    string             `bson:"email" json:"email"`
	IsAdmin  bool               `bson:"is_admin" json:"is_admin"`
	Verified bool               `bson:"verified" json:"verified"`
}

func GetUserWithID(idHex string) (User, bool, error) {
	var user User

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return user, false, err
	}

	err = colUser.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return user, false, nil
	}
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

func GetUserWithUsername(username string) (User, bool, error) {
	var user User
	err := colUser.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return user, false, nil
	}
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	result, err := colUser.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = result.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func IsUserAdmin(idHex string) (bool, error) {
	user, found, err := GetUserWithID(idHex)
	if !found {
		return false, errors.New("user with _id " + idHex + " not found")
	}
	if err != nil {
		return false, err
	}
	return user.IsAdmin, nil
}

func AddUser(user User) (string, error) {
	user.ID = primitive.NewObjectID()
	_, err := colUser.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}
	return user.ID.Hex(), nil
}

func UpdateUser(idHex string, info bson.M) error {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": info,
	}

	_, err = colUser.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	return err
}

func DeleteUser(idHex string) error {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return err
	}

	_, err = colUser.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
