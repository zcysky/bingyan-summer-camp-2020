package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func Mstart() *mongo.Client {
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

func Insert(client *mongo.Client, r *Remind) error {
	collection := client.Database(Info.DataBase).Collection(Info.Collection)
	_, err := collection.InsertOne(context.TODO(), r)
	if err != nil {
		fmt.Println("插入数据库失败")
		return err
	}
	return nil
}

func ReplyAll(client *mongo.Client, m *PostInfo) bool{
	collection := client.Database(Info.DataBase).Collection(Info.Collection)
	findOptions := options.Find()
	var results []*Remind
	// 把bson.D{{}}作为一个filter来匹配所有文档
	cur, err := collection.Find(context.TODO(), bson.D{{"qq", m.UserId}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem Remind
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
		if m.GroupId == 0 {
			SendPrivateMessage(m.UserId, "ID："+elem.ID)
			SendPrivateMessage(m.UserId, "标题："+elem.Title)
			SendPrivateMessage(m.UserId, "内容："+elem.Content)
		} else {
			SendGroupMessage(m.GroupId, "ID："+elem.ID)
			SendGroupMessage(m.GroupId, "标题："+elem.Title)
			SendGroupMessage(m.GroupId, "内容："+elem.Content)
		}

		time.Sleep(500000000)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// 完成后关闭游标
	cur.Close(context.TODO())
	if len(results) == 0 {
		return false
	}
	return true
}

func Delete(client *mongo.Client, m *PostInfo) bool {
	collection := client.Database(Info.DataBase).Collection(Info.Collection)
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"id", m.Message}})
	if err != nil {
		//log.Fatal(err)
		return false
	}
	//fmt.Printf("Deleted %v documents in the collection\n", deleteResult1.DeletedCount)
	if deleteResult.DeletedCount > 0 {
		return true
	}
	return false
}

func ChangeTitle(client *mongo.Client, m *PostInfo) int64 {
	collection := client.Database(Info.DataBase).Collection(Info.Collection)
	filter := bson.D{{"id", RemindID}}
	update := bson.D{
		{"$set", bson.D{
			{"title", m.Message},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0
	}
	return result.MatchedCount
}

func ChangeContent(client *mongo.Client, m *PostInfo) int64 {
	collection := client.Database(Info.DataBase).Collection(Info.Collection)
	filter := bson.D{{"id", RemindID}}
	update := bson.D{
		{"$set", bson.D{
			{"content", m.Message},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0
	}
	return result.MatchedCount
}

func ChangeTime(client *mongo.Client, t int64) int64 {
	collection := client.Database(Info.DataBase).Collection(Info.Collection)
	filter := bson.D{{"id", RemindID}}
	update := bson.D{
		{"$set", bson.D{
			{"ddl", t},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0
	}
	return result.MatchedCount
}

func ChangeAheadTime(client *mongo.Client, ahead int, interval int) int64{
	collection := client.Database(Info.DataBase).Collection(Info.Collection)
	filter := bson.D{{"id", RemindID}}
	update := bson.D{
		{"$set", bson.D{
			{"ahead", ahead},
			{"interval", interval},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0
	}
	return result.MatchedCount
}

func CheckID(client *mongo.Client, ID string) (bool, *Remind) {
	collection := client.Database(Info.DataBase).Collection(Info.Collection)
	filter := bson.D{{"id", ID}}
	result := new(Remind)
	err := collection.FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		return false, nil
	}
	return true, result
}
