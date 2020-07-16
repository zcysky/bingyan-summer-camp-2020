package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

//chorme UA

type LIVEinfo struct {
	LIVEinfo struct {
		Title       string `json:"title" bson:"Title"`
		Uid         int64  `json:"Uid" bson:"Uid"`
		Status      int64  `json:"live_status" bson:"Status"`
		Avatar      string `json:"user_cover" bson:"Avatar"`
		Description string `json:"description" bson:"Description"`
	} `json:"data"`
}

const (
	bilibiliRoomInfoUrl = "https://api.live.bilibili.com/room/v1/Room/get_info"
	myChromeUa          = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"
)

func HttpRequest(url string) (*http.Response,error){

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil,err
	}
	req.Header.Set("User-Agent", myChromeUa)
	var resp *http.Response
	resp, err = client.Do(req)
	//resp,err:=http.Get(bilibiliRoomInfoUrl+"?"+params)
	//fmt.Println(bilibiliRoomInfoUrl+"?"+params)
	if err != nil {
		return nil,err
	}
	return resp,nil
}

func getLIVEInfo(id int64) (LIVEinfo,error) {
	//contentType:="application/json"
	params := "id=" + strconv.Itoa(int(id))
	resp,err := HttpRequest(bilibiliRoomInfoUrl + "?" + params)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return LIVEinfo{},err
	}
	info := LIVEinfo{}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return LIVEinfo{},err
	}
	return info,nil
}
