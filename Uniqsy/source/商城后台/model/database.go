package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mall/config"
)

//链接
var client *mongo.Client
var (
	//表
	usersCol, commoditiesCol, keyWordsCol *mongo.Collection
)

//初始化数据库配置信息
func InitDataBase() {
	//配置连接数据库的信息
	URL := config.Config.MongoDB.URL
	clientOptions := options.Client().ApplyURI(URL)

	//建立链接
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	//检测连接是否正常
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	log.Println("Connected to remote database")

	usersCol = client.Database("mall").Collection("users")
	commoditiesCol = client.Database("mall").Collection("commodities")
	keyWordsCol = client.Database("mall").Collection("keywords")
	log.Println("Initialize database successfully")
}