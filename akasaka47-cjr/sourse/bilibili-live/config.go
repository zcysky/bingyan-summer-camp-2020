package main

import (
	"encoding/json"
)

type Rooms struct {
	Id     []int  `json:"id"`
	/*Head   string `json:"head"`
	UID    string `json:"uid"`
	Status string `json:"status"`
	Infom  string `json:"info"`*/
}

func Room(r *Rooms) {
	var roominfo = []byte(`{
  	"id": [
    9196015,
	1658468,
    3749977,
	7215179
  	]}`)

	err := json.Unmarshal(roominfo, r)
	if err != nil {
		return
	}
	//fmt.Println(r.Id)
}


