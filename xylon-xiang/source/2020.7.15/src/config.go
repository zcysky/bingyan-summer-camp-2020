package src

import (
	"encoding/json"
	"time"
)

//const BASELISTURL = "https://api.live.bilibili.com/room/v3/area/getRoomList"

type Config struct {
	IdList       []string      `json:"id_list"`
	TimeInterval time.Duration `json:"time_interval"`
}

func ApiConfig() ([]byte, error) {

	//the room id list
	list := []string{"81711", "5050", "271744", "90713", "923833", "5441", "1010",
		"1440094", "5096", "1569975", "8385390", "313585", "713115", "21504767", "12326244", "4190415", "3066386"}

	config := Config{
		IdList:       list,
		TimeInterval: 5 * time.Millisecond,
	}

	//encode the config Json to transport
	configJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	return configJson, nil
}
