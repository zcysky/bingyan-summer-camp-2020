package main

import (
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"sync"
)

var wg sync.WaitGroup

func loadLiveInfo(loadChannel <-chan int64,storeChannel chan<- LIVEinfo){
	defer wg.Done()
	for id:=range loadChannel {
		storeChannel <- getLIVEInfo(id)
		fmt.Println("finish", id)
	}
}

func main() {
	configInfo := LoadConfig("./Config.json")
	loadChannel:=make(chan int64,100)
	storeChannel:=make(chan LIVEinfo,100)
	for i:=0;i<3;i++{
		go loadLiveInfo(loadChannel,storeChannel)
		wg.Add(1)
	}
	for i, value := range configInfo.LiveId {
		fmt.Println(value,"is in the sequence")
		loadChannel<-value
		if (i+1)%6==0 {
			time.Sleep(time.Duration(configInfo.FetchInterval) * time.Millisecond)
		}
	}
	close(loadChannel)
	wg.Wait()
	close(storeChannel)
	var dataSet []LIVEinfo
	for value:=range storeChannel {
		dataSet=append(dataSet,value)
	}
	res,err:=json.Marshal(dataSet)
	if(err!=nil){
		fmt.Println("err when Marshal storeChannel",err)
	}
	ioutil.WriteFile("./data.json",res,0644)
}
