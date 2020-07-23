package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"

	"github.com/Logiase/gomirai/message"
)

var MongoClient *(mongo.Client)

func ConnectMongoDataBase() error {
	var err error
	MongoClient, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = MongoClient.Connect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func InsertMemo(userId string, eventInfo message.Event) error {

	col := MongoClient.Database("robot").Collection(userId)
	_, err := col.InsertOne(context.Background(), eventInfo)
	if err != nil {
		return err
	}
	return nil
}

func Find(userId string, eventId string) (message.Event, error) {
	col := MongoClient.Database("robot").Collection(userId)
	filter := bson.D{{"eventId", eventId}}
	var UserMemo message.Event

	err := col.FindOne(context.TODO(), filter).Decode(&UserMemo)
	if err != nil {
		return message.Event{}, err
	}
	//fmt.Println(registerInfo)
	return UserMemo, nil
}

func ShowAllEvent(userId string) ([]message.Event, error) {
	var AllEvent []message.Event
	col := MongoClient.Database("robot").Collection(userId)
	cur, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var Event message.Event
		err := cur.Decode(&Event)
		if err != nil {
			return nil, err
		}
		AllEvent = append(AllEvent, Event)
	}
	return AllEvent, nil
}

//func Update(UserInfo config.RegisterInfo) error {
//	col := MongoClient.Database("project").Collection("users")
//	filter := bson.D{{"uid", UserInfo.Uid}}
//	update := bson.D{{"$set", UserInfo}}
//	//fmt.Println("->>>",UserInfo.Uid)
//	var updatedDocument bson.M
//	err := col.FindOneAndUpdate(context.TODO(), filter, update).Decode(&updatedDocument)
//	if err != nil {
//		// ErrNoDocuments means that the filter did not match any documents in the collection
//		if err == mongo.ErrNoDocuments {
//			return err
//		}
//	}
//	return nil
//}

func Delete(userId string, eventId string) error { //delete删除前检查文档是否存在
	col := MongoClient.Database("robot").Collection(userId)
	eventIdInt, err := strconv.Atoi(eventId)
	if err != nil {
		return err
	}
	//fmt.Println("->>>>>>", eventIdInt)
	var result bson.M
	err = col.FindOne(context.TODO(), bson.D{{"eventid", eventIdInt}}).Decode(&result)
	if err != nil {
		return err
	}

	filter := bson.D{{"eventid", eventIdInt}}
	_, err = col.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	err := ConnectMongoDataBase()
	if err != nil {
		panic(err)
	}
}

func InsertNoti(userId string, eventInfo message.Event) error {

	col := MongoClient.Database("robot").Collection(userId)
	_, err := col.InsertOne(context.Background(), eventInfo)
	if err != nil {
		return err
	}
	return nil
}