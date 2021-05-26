package main

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	initializeLogger()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	key := "postback_queue"

	getNextValue(rdb, key)

}

func initializeLogger() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func getNextValue(redisConn *redis.Client, key string) {
	// start infinite loop
	for {
		// five-second timeout on empty
		val, err := redisConn.BRPop(ctx, 5000000000, key).Result()
		if err == redis.Nil {
			log.Printf("'%v' is currently empty", key)
		} else if err != nil {
			panic(err)
		} else {
			// need to get index[1] when using BRPop vs normal RPop because
			// the val returned looks like "[postback_queue test]" and we want "test"
			log.Printf("Key/Val RPopped: %v:%v", key, val[1])
		}
	}
}
