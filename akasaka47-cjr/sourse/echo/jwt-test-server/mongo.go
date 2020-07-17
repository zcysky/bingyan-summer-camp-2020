package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Mstart() *mongo.Client{
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
	//开始操作
}

func Mend(client *mongo.Client){
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func Insert(client *mongo.Client, u *User) error{
	collection := client.Database("test").Collection("user-data-test")
	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		fmt.Println("插入数据库失败")
		return err
	}
	return nil
}

func Check(client *mongo.Client, u *User) bool{
	collection := client.Database("test").Collection("user-data-test")
	filter := bson.D{{"email", u.Email}}
	var result User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}
	return true
}