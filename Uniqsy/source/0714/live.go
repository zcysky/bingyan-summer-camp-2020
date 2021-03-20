package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type RoomInfo struct {
	Title		string `json:"title"`
	Uid			int    `json:"uid"`
	LiveStatus	int    `json:"live_status"`
	UserCover	string	`json:"user_cover"`
	Description	string	`json:"description"`
}

type ResponseInfo struct {
	Message	string	`json:"message"`
	Data 	RoomInfo	`json:"data"`
}

func GetRoomInfo(roomid int) *RoomInfo {
	client := &http.Client{}

	url := "https://api.live.bilibili.com/room/v1/Room/get_info?id="
	url += strconv.Itoa(roomid)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent", Config.UserAgent)

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var responseInfo *ResponseInfo = new(ResponseInfo)
	err = json.Unmarshal(s, &responseInfo)
	return &responseInfo.Data
}