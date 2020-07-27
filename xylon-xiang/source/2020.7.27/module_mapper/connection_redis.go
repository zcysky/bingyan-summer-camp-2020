package module_mapper

import (
	"2020.7.27/config"
	"github.com/go-redis/redis/v8"
)

var RegisterRedis *redis.Client

func RegisterCodeDB() {
	RegisterRedis = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.RedisAddress,
		Password: "",
		DB:       0,
	})
}

func init() {
	RegisterCodeDB()
}
