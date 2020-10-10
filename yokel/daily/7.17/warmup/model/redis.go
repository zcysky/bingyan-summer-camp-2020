package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"warmup/config"
)

var RedisClient redis.Conn

func ConnectRedisDataBase() error {
	var err error
	RedisClient, err = redis.Dial("tcp", "127.0.0.1:6379")
	fmt.Println("->>>>>>>>>>>>>>>>>",RedisClient)
	if err != nil {
		return  err
	}
	return nil
}

func FindCode(uid string) (string, error) {
	code, err := redis.String(RedisClient.Do("GET", uid))
	if err != nil {
		return "", err
	}
	return code, nil
}

func InsertCode(uid string, code string) error {
	_, err := RedisClient.Do("SET", uid, code)
	if err != nil {
		return err
	}
	return nil
}

func ExistUser(uid string) (bool, error) {
	uidExist, err := redis.Bool(RedisClient.Do("EXISTS", uid))
	if err != nil {
		return false, nil
	}
	return uidExist, nil
}

func DeleteCode(uid string) error {
	_, err := RedisClient.Do("DEL", uid)
	if err != nil {
		return err
	}
	return nil
}

func AddCode(uid string, code int) error {
	_, err := RedisClient.Do("SET", uid, code, "EX", config.RedisExpirationTime)
	if err != nil {
		return nil
	}
	return nil
}
