// Michael Peters
// redis.go
// Redis Connection for the Kochava Postback Delivery Agent
// Last Modified: 05/28/21 18:50 PDT

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

// activate connection to Redis and return Client object
func redisConnect() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ip := os.Getenv("REDIS_IP")
	port := os.Getenv("REDIS_PORT")

	address := fmt.Sprintf("%v:%v", ip, port)

	fmt.Println("Connecting to Redis...")
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
