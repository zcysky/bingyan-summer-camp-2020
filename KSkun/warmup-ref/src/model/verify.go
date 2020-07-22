package model

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
	"warmup-ref/config"
)

const formatVerifyCodeKey = "code:%s"

var codeExpire time.Duration

func InitModelVerify() {
	codeExpire = time.Duration(config.Config.App.VerifyCodeExpire) * time.Minute
}

func AddVerifyCode(code string, idHex string) error {
	key := fmt.Sprintf(formatVerifyCodeKey, code)
	return RedisClient.Set(key, idHex, codeExpire).Err()
}

func GetVerifyCode(code string) (string, bool, error) {
	key := fmt.Sprintf(formatVerifyCodeKey, code)
	result, err := RedisClient.Get(key).Result()
	if err == redis.Nil {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return result, true, nil
}

func DeleteVerifyCode(code string) error {
	key := fmt.Sprintf(formatVerifyCodeKey, code)
	return RedisClient.Del(key).Err()
}
