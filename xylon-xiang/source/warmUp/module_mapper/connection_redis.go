package module_mapper

import (
	"github.com/go-redis/redis/v8"
	"warmUp/config"
)

var RegisterRedis *redis.Client

func RegisterCodeDB() {
	RegisterRedis = redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.RedisAddress,
		Password: "",
		DB: 0,
	})
}

func init() {
	RegisterCodeDB()
}