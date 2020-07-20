package model

import (
	"github.com/garyburd/redigo/redis"
	"warmup/config"
)

func FindCode(uid string) (string, error) {
	redisDataBase, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return "", err
	}
	defer redisDataBase.Close()
	code, err := redis.String(redisDataBase.Do("GET", uid))
	if err != nil {
		return "", err
	}
	return code, nil
}

func InsertCode(uid string, code string) error {
	redisDataBase, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return err
	}
	defer redisDataBase.Close()
	_, err = redisDataBase.Do("SET", uid, code)
	if err != nil {
		return err
	}
	return nil
}

func ExistUser(uid string) (bool, error) {
	redisDataBase, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return false, err
	}
	defer redisDataBase.Close()
	uidExist, err := redis.Bool(redisDataBase.Do("EXISTS", uid))
	if err != nil {
		return false, nil
	}
	return uidExist, nil
}

func DeleteCode(uid string) error {
	redisDataBase, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return err
	}
	defer redisDataBase.Close()
	_, err = redisDataBase.Do("DEL", uid)
	if err != nil {
		return err
	}
	return nil
}

func AddCode(uid string,code int)error{
	redisDataBase, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return  err
	}
	defer redisDataBase.Close()
	_, err = redisDataBase.Do("SET", uid,code,"EX",config.RedisExpirationTime)
	if err != nil {
		return nil
	}
	return nil
}