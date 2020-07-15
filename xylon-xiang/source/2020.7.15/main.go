package main

import (
	"2020.7.15/src"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	. "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var routinePool sync.WaitGroup

func writeJsonIntoDB(roomId string, interval time.Duration) error {
	// create a connection with the MongoDB
	client, err := NewClient(options.Client().ApplyURI("mongodb://@localhost:27017"))
	if err != nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	// choose the "bilibili live room info" collection
	collection := client.Database("test").Collection("bilibiliLiveRoomInfo")

	// get the []byte room information
	roomInfoByte, err := src.GetLiveRoomInfo(roomId)
	if err != nil {
		return err
	}

	// decode the []byte data into struct
	var roomInfoStruct src.RoomInfo
	err = json.Unmarshal(roomInfoByte, &roomInfoStruct)
	if err != nil {
		return err
	}

	// insert the data into db
	_, err = collection.InsertOne(context.Background(), bson.M{
		"title":       roomInfoStruct.Title,
		"description": roomInfoStruct.Description,
		"live_status": roomInfoStruct.LiveStatus,
		"uid":         roomInfoStruct.Uid,
		"user_cover":  roomInfoStruct.UserCover,
	})
	if err != nil {
		return err
	}

	// the interval time after a api request
	time.Sleep(interval)


	routinePool.Done()

	// if no error, then return the nil
	return nil
}

func roomIdEnterList(roomIdList []string, list chan string)  {

	for i := 0; i < len(roomIdList); i++ {
		list <- roomIdList[i]
	}

	routinePool.Done()
}

func getIdListAndTimeInterval() ([]string, time.Duration, error){

	// get the config from the src/config.go
	config := src.Config{}

	apiConfig, err := src.ApiConfig()
	if err != nil {
		return nil, 0, err
	}
	err = json.Unmarshal(apiConfig, &config)
	if err != nil {
		return nil, 0, err
	}

	roomIdList := config.IdList
	timeInterval := config.TimeInterval

	return roomIdList, timeInterval, nil

}

func main() {

	roomIdList, timeInterval, err := getIdListAndTimeInterval()
	if err != nil{
		fmt.Println(err)
	}

	roomId := make(chan string)

	routinePool.Add(1)
	go roomIdEnterList(roomIdList, roomId)


	for j := range roomId{
		routinePool.Add(1)
		go writeJsonIntoDB(j, timeInterval)
	}


	routinePool.Wait()

	//
	//pool.Wait()
}
