// Michael Peters
// redis.go
// Redis Connection for the Kochava Postback Delivery Agent
// Last Modified: 05/28/21 18:50 PDT

package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

// activate connection to Redis and return Client object
func redisConnect() *redis.Client {
	fmt.Println("Connecting to Redis...")
	return redis.NewClient(&redis.Options{
		Addr:     "10.0.0.15:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
