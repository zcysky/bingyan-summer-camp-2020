package model

import (
	"../config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func InsertNewUser(UserInfo config.RegisterInfo) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	col := client.Database("project").Collection("users")
	_, err = col.InsertOne(context.Background(), UserInfo)
	if err != nil {
		return err
	}
	return nil
}
