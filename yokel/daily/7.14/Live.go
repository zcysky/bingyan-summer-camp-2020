package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//chorme UA

type LIVEinfo struct {
	LIVEinfo struct {
		Title       string `json:"title"`
		Uid         int64  `json:"Uid"`
		Status      int64  `json:"live_status"`
		Avatar      string `json:"user_cover"`
		Description string `json:"description"`
	} `json:"data"`
}

const (
	bilibiliRoomInfoUrl = "https://api.live.bilibili.com/room/v1/Room/get_info"
	myChromeUa          = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"
)

func getLIVEInfo(id int64) LIVEinfo {
	//contentType:="application/json"
	params := "id=" + strconv.Itoa(int(id))

	client := &http.Client{}
	req, err := http.NewRequest("GET", bilibiliRoomInfoUrl+"?"+params, nil)
	if err != nil {
		fmt.Println("err when creating NewRequest", err)
	}
	req.Header.Set("User-Agent", myChromeUa)
	var resp *http.Response
	resp, err = client.Do(req)
	//resp,err:=http.Get(bilibiliRoomInfoUrl+"?"+params)
	//fmt.Println(bilibiliRoomInfoUrl+"?"+params)
	if err != nil {
		fmt.Println("err when GET bilibili LIVE room", id)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err when ReadAll", err)
	}
	info := LIVEinfo{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		fmt.Println("erro when Unmarshal", err)
	}
	return info
}
