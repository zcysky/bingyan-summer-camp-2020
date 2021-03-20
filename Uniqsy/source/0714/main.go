package main

import (
	"fmt"
	"time"
)

func main() {
	Init()

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
			for _, val := range Config.List {
				fmt.Println(GetRoomInfo(val.Roomid).Title)
			}
		}

	}()
}
