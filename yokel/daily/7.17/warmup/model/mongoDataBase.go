package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"warmup/config"
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

func ShowAllUser()([]config.RegisterInfo,error){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil,err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil,err
	}
	var AllUser []config.RegisterInfo
	col := client.Database("project").Collection("users")
	cur,err:=col.Find(context.TODO(),bson.D{})
	if err != nil {
		return nil,err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()){
		var User config.RegisterInfo
		err:=cur.Decode(&User)
		if err != nil {
			return nil,err
		}
		AllUser=append(AllUser, User)
	}
	return AllUser,nil
}

func FindUser(userId string) (config.RegisterInfo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return config.RegisterInfo{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return config.RegisterInfo{}, err
	}
	col := client.Database("project").Collection("users")
	filter := bson.D{{"uid", userId}}
	var registerInfo config.RegisterInfo

	findResult := col.FindOne(context.TODO(), filter)
	err = findResult.Decode(&registerInfo)
	if err != nil {
		return config.RegisterInfo{}, err
	}
	//fmt.Println(registerInfo)
	return registerInfo, nil
}

func UpdateUser(UserInfo config.RegisterInfo)error{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return  err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	col := client.Database("project").Collection("users")
	filter := bson.D{{"uid", UserInfo.Uid}}
	update:=bson.D{{"$set", UserInfo}}
	//fmt.Println("->>>",UserInfo.Uid)
	var updatedDocument bson.M
	err=col.FindOneAndUpdate(context.TODO(), filter, update).Decode(&updatedDocument)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return err
		}
	}
	return nil
}

func DeleteUser(userId string)error{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	col := client.Database("project").Collection("users")
	filter := bson.D{{"uid",userId}}
	_,err = col.DeleteOne(context.TODO(), filter)
	if(err!=nil){
		return err
	}
	return nil
}