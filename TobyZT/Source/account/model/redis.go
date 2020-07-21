package model

import (
	"github.com/garyburd/redigo/redis"
	"strconv"
)

var c redis.Conn

func SetKey(key string, value string, expMinute int) (err error) {
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return err
	}
	_, err = c.Do("SET", key, value, "EX", strconv.Itoa(expMinute*60))
	if err != nil {
		return err
	}
	return nil
}

func GetKey(key string) (value string, err error) {
	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return "", err
	}
	value, err = redis.String(c.Do("GET", key))
	if err != nil {
		return "", err
	}
	return value, nil
}
