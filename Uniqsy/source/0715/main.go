package main

import (
	"encoding/json"
	"os"
	"time"
)

type ResultInfo struct {
	Time 		time.Time	`json:"time"`
	//这里将房间数硬编码了，不知道有啥好的解决办法
	RoomInfos	[3]RoomInfo	`json:"room_infos"`
}

func main() {
	Init()
	//设置计时器，固定时间拉取信息
	ticker := time.NewTicker(time.Second * time.Duration(Config.Interval))
	//清除文件内容，允许文件读写
	file, _ := os.OpenFile("result", os.O_RDWR|os.O_TRUNC, 0666)

	//判断文件是否为第一次打开，如果不是第一次打开，就先加一个逗号，使文件内容符合json格式
	first := true

	go func() {
		for t := range ticker.C {
			if !first {
				file.Write([]byte(","))
			}
			first = false

			var resultInfo *ResultInfo = new(ResultInfo)
			resultInfo.Time = t

			done := make(chan bool, 1)
			go func() {
				for key, val := range Config.List {
					resultInfo.RoomInfos[key] = GetRoomInfo(val.Roomid)
				}
				done <- true
			}()

			// 等待所有信息都获取完毕后，再将更新的信息追加写入文件
			<-done
			byteResult, err := json.Marshal(resultInfo)
			if err != nil {
				panic(err)
			}

			file.Write(byteResult)
		}

	}()

	//总的运行时长也是硬编码进入，不知道怎么处理
	time.Sleep(time.Second * 12)
	ticker.Stop()
	file.Close()
}
