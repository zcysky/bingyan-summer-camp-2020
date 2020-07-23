package module

import (
	"QQRobot/src/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	UserCol *mongo.Collection
	MemoCol *mongo.Collection
)

func connectMongoDB() {
	//creat a client
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.Database.DatabaseAddress))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	//create a connection
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	UserCol = client.Database(config.Config.Database.DatabaseName).
		Collection(config.Config.Database.CollectionFriendName)

	MemoCol = client.Database(config.Config.Database.DatabaseName).
		Collection(config.Config.Database.CollectionMemoName)
}

func init() {
	connectMongoDB()
}
