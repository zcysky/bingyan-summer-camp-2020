package main

import (
	"2020.7.14/src"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var pool sync.WaitGroup

func WriteJsonIntoFile(roomId string, interval time.Duration) {

	//open the file "BilibiliLiveRoomInfo.json"
	file, err := os.OpenFile(
		"BilibiliLiveRoomInfo.json",
		os.O_APPEND|os.O_WRONLY,
		os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	// close the file after the function return
	defer file.Close()

	//get the room info
	roomInfo := src.GetLiveRoomInfo(roomId)

	//byteNum is the length of the written data
	byteNum, err := file.Write(roomInfo)
	if err != nil {
		log.Fatal(err)
	}

	//print the length to specify the written data
	fmt.Println(byteNum)

	time.Sleep(interval)

	//sub 1
	pool.Done()

}

func main() {

	// get the config from the src/config.go
	config := src.Config{}
	err := json.Unmarshal(src.ApiConfig(), &config)
	if err != nil {
		fmt.Println(err)
	}


	roomIdList := config.IdList
	timeInterval := config.TimeInterval


	for i := 0; i < len(roomIdList); i = i + 3 {
		pool.Add(1)
		go WriteJsonIntoFile(roomIdList[i], timeInterval)
		if i+1 >= len(roomIdList) {
			break
		}

		pool.Add(1)
		go WriteJsonIntoFile(roomIdList[i+1], timeInterval)

		if i+2 >= len(roomIdList) {
			break
		}

		pool.Add(1)
		go WriteJsonIntoFile(roomIdList[i+2], timeInterval)

		pool.Wait()
	}

	pool.Wait()
}
