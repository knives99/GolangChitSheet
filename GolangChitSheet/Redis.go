package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type RediusStore struct {
	Client *redis.Client
}

var RediusStroe RediusStore

func main() {

	InitRedisStore()
	redisStore := RediusStroe.Client
	ctx := context.Background()

	value, _ := redisStore.Get(ctx, "").Result()

	err := redisStore.Set(ctx, "key", "value", time.Minute*10).Err()
	if err != nil {
		return
	}

	fmt.Println(value)
}

func InitRedisStore() *RediusStore {
	//config := InitCnfig()
	redis := redis.NewClient(&redis.Options{
		Addr:     "locahost:8000",
		Password: "",
		DB:       0,
	})
	RediusStroe = RediusStore{Client: redis}
	//base64Captcha.NewCaptcha("", RediusStroe)
	return &RediusStroe
}

//讀取自己本地端的JSON
func InitCnfig() *string {

	var conv *string
	file, err := os.Open("/kjdkdj/jdjdk")
	if err != nil {

	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conv)
	if err != nil {

	}

	return conv
}
