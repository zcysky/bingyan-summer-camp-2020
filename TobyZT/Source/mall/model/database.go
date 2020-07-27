package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var userColl *mongo.Collection
var commodityColl *mongo.Collection

func SetupDatabase() (err error) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	userColl = client.Database("mall").Collection("users")
	commodityColl = client.Database("mall").Collection("commodities")
	log.Println("Database connected successfully!")
	return err
}

func VerifyLogin(form LoginForm) (valid bool) {
	filter := bson.M{"username": form.Username}
	var res User
	userColl.FindOne(context.TODO(), filter).Decode(&res)
	if res.Username != "" && Compare(res.Password, form.Password) {
		return true
	}
	return false
}

func Signup(form User) (err error) {
	form.Password = Encrypt(form.Password)
	_, err = userColl.InsertOne(context.TODO(), form)
	return err
}

func QueryOne(username string) (form User, err error) {
	filter := bson.M{"username": username}
	var res User
	err = userColl.FindOne(context.TODO(), filter).Decode(&res)
	return res, err
}

func Update(username string, updateForm UpdateForm) (err error) {
	filter := bson.M{"username": username}
	update := bson.M{"$set": updateForm}
	_, err = userColl.UpdateOne(context.TODO(), filter, update)
	return err
}
