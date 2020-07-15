package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"time"
)
import (
	//"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var wg sync.WaitGroup

func loadLiveInfo(loadChannel <-chan int64, storeChannel chan<- LIVEinfo) {
	defer wg.Done()
	for id := range loadChannel {
		storeChannel <- getLIVEInfo(id)
		fmt.Println("finish", id)
	}
}

func StoreToFile(url string, storeChannel chan LIVEinfo) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	fmt.Println(client, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = client.Connect(ctx)
	collection := client.Database("test").Collection("project1")

	for value := range storeChannel {
		_, err = collection.InsertOne(context.Background(), bson.M{
			"Title":       value.LIVEinfo.Title,
			"Uid":         value.LIVEinfo.Uid,
			"Status":      value.LIVEinfo.Status,
			"Avatar":      value.LIVEinfo.Avatar,
			"Description": value.LIVEinfo.Description,
		})
		if err != nil {
			fmt.Println(err)
		}
	}

}

func main() {
	configInfo := LoadConfig("./Config.json")
	loadChannel := make(chan int64, 100)
	storeChannel := make(chan LIVEinfo, 100)
	for i := 0; i < 3; i++ {
		go loadLiveInfo(loadChannel, storeChannel)
		wg.Add(1)
	}
	for i, value := range configInfo.LiveId {
		fmt.Println(value, "is in the sequence")
		loadChannel <- value
		if (i+1)%6 == 0 {
			time.Sleep(time.Duration(configInfo.FetchInterval) * time.Millisecond)
		}
	}
	close(loadChannel)
	wg.Wait()
	close(storeChannel)
	StoreToFile("./data.json", storeChannel)
}
