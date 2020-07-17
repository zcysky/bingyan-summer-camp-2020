package src

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	BASEURL = "https://api.live.bilibili.com/room/v1/Room/get_info?id="

	USERAGENT = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 " +
		"(KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"
)

type ResponseFormat struct {
	Data RoomInfo `json:"data"`
}

type RoomInfo struct {
	Title       string `json:"title"`
	Uid         int64  `json:"uid"`
	LiveStatus  int    `json:"live_status"`
	UserCover   string `json:"user_cover"`
	Description string `json:"description"`
}

func GetLiveRoomInfo(roomID string) ([]byte, error) {

	responseBody, err := getResponseBody(roomID)
	if err != nil {
		return nil, err
	}


	// decode the response body Json data
	responseFormat := ResponseFormat{}
	err = json.Unmarshal(responseBody, &responseFormat)
	if err != nil {
		return nil, err
	}

	roomInfo := responseFormat.Data

	// encode the room information Json
	roomInfoJson, err := json.Marshal(roomInfo)
	if err != nil {
		return nil, err
	}

	return roomInfoJson, nil

}

func getResponseBody(uid string) ([]byte, error) {
	Url := BASEURL + uid

	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, Url, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", USERAGENT)

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}

	return nil, nil
}
