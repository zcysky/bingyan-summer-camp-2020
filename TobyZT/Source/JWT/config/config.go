package config

type User struct {
	Username string
	Password string
}

type JsonInfo struct {
	Secret string `json:"secret"`
	Expire int `json:"expire"`
}
