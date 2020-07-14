package main

import (
	"fmt"
	"time"

)

func main() {
	ConfigInfo := LoadConfig("./Config.json")
	for _,value :=range ConfigInfo.LIVEId {
		time.Sleep(time.Duration(ConfigInfo.FetchInterval) * time.Millisecond)
		newLiveInfo:=getLIVEInfo(value)
		fmt.Println(newLiveInfo)
	}
	fmt.Println(ConfigInfo)
}
