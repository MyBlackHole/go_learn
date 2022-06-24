package main

import (
	"encoding/json"
	"time"
	"log"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func init() {
	redisdb = redis.NewClient(&redis.Options{
		// Addr:     "localhost:6379",
		Addr:     "localhost:6380",
		Password: "",              
		DB:       6,               
	})

	//心跳
	pong, err := redisdb.Ping().Result()
	if err!=nil {
		log.Printf("pong: [%s], err: [%v]\n", pong, err)
	}
}


func save(postjson Json,t time.Duration) {
	var redisData RedisData
	redisData.Time = CurrentDateTime()
	redisData.Json = postjson
	key := CurrentDate()

	redisDataStr, err := json.Marshal(redisData)
	if err != nil {
		log.Print(err)
	}

    pipe := redisdb.TxPipeline()
	pipe.SAdd(key, redisDataStr)
	pipe.Expire(key, t)
    _, err = pipe.Exec()
	if err != nil {
		log.Print(err)
	}
}