package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func SetupDatabase() {
	//设置客户端连接配置
	mongodbURL := "mongodb://47.98.224.238:27017"
	clientOptions := options.Client().ApplyURI(mongodbURL)

	var err error
	//连接到MongoDB
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	//检查是否成功连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("conneted to mongodb successfully")
}

func TestConnection() error {
	err := client.Ping(context.TODO(), nil)
	return err
}