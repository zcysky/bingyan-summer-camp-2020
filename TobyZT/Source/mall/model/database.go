package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var userColl *mongo.Collection
var commodityColl *mongo.Collection
var keywordColl *mongo.Collection
var commentColl *mongo.Collection

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
	keywordColl = client.Database("mall").Collection("keywords")
	commentColl = client.Database("mall").Collection("comments")
	log.Println("Database connected successfully!")
	return err
}
