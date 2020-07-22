package model

import (
	"context"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"strconv"
	"time"
	"warmup-ref/config"
)

var (
	MongoDatabase *mongo.Database
	RedisClient *redis.Client
)

func InitModel() {
	// init mongodb
	mongoUri := "mongodb://" + config.Config.Mongo.Address + ":" + strconv.Itoa(config.Config.Mongo.Port)
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Println("model: error when init mongo client with uri " + mongoUri)
		log.Panic(err)
	}

	// ping test for mongodb
	pingContext, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = mongoClient.Ping(pingContext, readpref.Primary())
	if err != nil {
		log.Println("model: unsuccessful ping to mongodb with uri " + mongoUri)
		log.Panic(err)
	}

	MongoDatabase = mongoClient.Database(config.Config.Mongo.Database)

	// init redis
	redisUri := config.Config.Redis.Address + ":" + strconv.Itoa(config.Config.Redis.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisUri,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// ping test for redis
	err = redisClient.Ping().Err()
	if err != nil {
		log.Println("model: unsuccessful ping to redis with uri " + redisUri)
		log.Panic(err)
	}

	RedisClient = redisClient

	initModelUser()
}
