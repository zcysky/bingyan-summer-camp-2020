/*将项目一的结果存入数据库*/

package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

type ResInfo struct {
	ID       []int
	UserName []string
	Title    []string
}

type SingleInfo struct {
	ID       int
	UserName string
	Title    string
}

func SaveResInfo(resInfo ResInfo) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Connected successfully!")

	collection := client.Database("live").Collection("HostInfo")

	for i := 0; i < len(resInfo.ID); i++ {
		filter := bson.M{"id": resInfo.ID[i]}
		var m SingleInfo
		err = collection.FindOne(context.TODO(), filter).Decode(&m)
		if m.ID == 0 { //which means it doesn't exist
			newInfo := bson.M{
				"id":       resInfo.ID[i],
				"username": resInfo.UserName[i],
				"title":    resInfo.Title[i],
			}
			_, err = collection.InsertOne(context.TODO(), newInfo)
			if err != nil {
				log.Fatal(err)
				return
			}
		} else {
			update := bson.M{
				"$set": bson.M{
					"id":       resInfo.ID[i],
					"username": resInfo.UserName[i],
					"title":    resInfo.Title[i],
				},
			}
			_, err = collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
		fmt.Println("Host " + strconv.Itoa(resInfo.ID[i]) + " saved.")
	}

}
