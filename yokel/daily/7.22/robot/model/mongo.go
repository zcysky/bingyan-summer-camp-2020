package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"regexp"
	"strconv"
	"time"
	"github.com/Logiase/gomirai/message"
)

type Notification struct {
	UserId   uint `json:"user-id"`
	EventId  int64  `json:"event-id"`
	Time     int64 `json:"time"`
	Interval int64 `json:"interval"`
	Advance  int64 `json:"advance"`
	NotiId int64 `json:"noti-id"`
}
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
	eventIdInt, err := strconv.Atoi(eventId)
	if err != nil {
		return message.Event{},err
	}
	fmt.Println(eventId)
	filter := bson.D{{"eventid", eventIdInt}}
	var UserMemo message.Event

	err = col.FindOne(context.TODO(), filter).Decode(&UserMemo)
	if err != nil {
		return message.Event{}, err
	}
	//fmt.Println(registerInfo)
	return UserMemo, nil
}

func ShowAllCollection()([]string,error){//应该能该用正则表达式匹配
	allCol,err:=MongoClient.Database("robot").ListCollectionNames(context.TODO(),bson.M{})
	if err!=nil {
		return []string{},err
	}

	var result []string
	rgex:="([0-9]+)Noti$"
	for _,col:=range allCol{
		//fmt.Println(col)
		match,err:=regexp.MatchString(rgex,col)
		if err!=nil{
			return []string{},err
		}
		if match ==true{
			result=append(result, col)
		}
	}
	return result,nil
}

func ShowAllNoti(userId string)([]Notification,error){
	var AllNoti []Notification
	col := MongoClient.Database("robot").Collection(userId)
	cur, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var Noti Notification
		err := cur.Decode(&Noti)
		if err != nil {
			return nil, err
		}
		AllNoti = append(AllNoti, Noti)
	}
	return AllNoti, nil

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


func AddNoti(newNoti Notification,userId string) error {
	//通知的事件必须在数据库内存在，且通知时也要进行检查，是否已经移出数据库
	_,err:=Find(userId,fmt.Sprint(newNoti.EventId))
	//fmt.Println(userId,fmt.Sprint(newNoti.EventId),err)
	if err!=nil{
		return err
	}
	col := MongoClient.Database("robot").Collection(userId+"Noti")
	_, err = col.InsertOne(context.Background(), newNoti)
	if err != nil {
		return err
	}
	return nil
}

func DeleteNoti(notiId int64,userId string)error{
	col := MongoClient.Database("robot").Collection(userId+"Noti")
	fmt.Println(userId+"Noti")
	var result bson.M
	err := col.FindOne(context.TODO(), bson.D{{"notiid", notiId}}).Decode(&result)
	if err != nil {
		return err
	}
	filter := bson.D{{"notiid", notiId}}
	_, err = col.DeleteOne(context.TODO(), filter)
	if err!=nil{
		return err
	}
	return nil
}