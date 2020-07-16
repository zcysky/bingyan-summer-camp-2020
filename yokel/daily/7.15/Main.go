package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
import (
	//"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var wg sync.WaitGroup

func loadLiveInfo(loadChannel <-chan int64, storeSlice *[]LIVEinfo)error {
	defer wg.Done()
	for id := range loadChannel {
		value,err:=getLIVEInfo(id)
		if err!=nil{
			return err
		}
		*storeSlice=append(*storeSlice, value)
		//fmt.Println(value)
		//fmt.Println("finish", id)
	}
	return nil
}

func StoreToFile(url string, storeSlice []LIVEinfo) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	//fmt.Println(client, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = client.Connect(ctx)
	collection := client.Database("test").Collection("project1")

	for _,value := range storeSlice  {
		//fmt.Print(value)
		_, err = collection.InsertOne(context.Background(), value)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func main() {
	configInfo,err:= LoadConfig("./Config.json")
	if(err!=nil){
		fmt.Println(err)
	}
	loadChannel := make(chan int64, 100)
	storeSlice := []LIVEinfo{}
	for i := 0; i < 3; i++ {
		go loadLiveInfo(loadChannel, &storeSlice)
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
	fmt.Println(storeSlice)
	StoreToFile("./data.json", storeSlice)
}
