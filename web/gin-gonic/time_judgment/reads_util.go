package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func init() {
    //连接服务器
    redisdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    //心跳
    _, err := redisdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}