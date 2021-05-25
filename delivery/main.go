package main

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	key := "test_val"
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	log.Println(key, val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// // Output: key value
	// // key2 does not exist
}
