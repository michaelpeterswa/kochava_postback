package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

// https://stackoverflow.com/questions/36322141/unmarshal-inconsistent-json
// https://github.com/Jeffail/gabs
type Postback struct {
	Endpoint Endpoint    `json:"endpoint,omitempty"`
	Data     []DataValue `json:"data,omitempty"`
}
type Endpoint struct {
	Method string `json:"method,omitempty"`
	Url    string `json:"url,omitempty"`
}
type DataValue struct {
	Mascot   string `json:"mascot,omitempty"`
	Location string `json:"location,omitempty"`
}

var ctx = context.Background()

func main() {

	name := "Delivery Agent v1.0.0"
	author := "Michael Peters"

	greet(name, author)
	initializeLogger()
	rdb := redisConnect()

	key := "postback_queue"

	dequeueLoop(rdb, key)

}

func redisConnect() *redis.Client {
	fmt.Println("Connecting to Redis...")
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func greet(name string, author string) {
	fmt.Printf("%v, by %v\n", name, author)
	fmt.Println("Starting process...")

}

func isValidJSON(input string) bool {
	return json.Valid([]byte(input))
}

func initializeLogger() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	fmt.Println("Opening logger... (logs.txt)")

}

func dequeueLoop(redisConn *redis.Client, key string) {
	// start infinite loop
	for {
		// five-second timeout on empty
		val, err := redisConn.BRPop(ctx, 5000000000, key).Result()
		if err == redis.Nil {
			// log.Printf("'%v' is currently empty", key)
		} else if err != nil {
			panic(err)
		} else {
			// need to get index[1] when using BRPop vs normal RPop because
			// the val returned looks like "[postback_queue test]" and we want "test"
			obj := val[1]
			log.Printf("Key/Val RPopped: %v:%v", key, val[1])
			if isValidJSON(obj) && obj != "null" {
				// log.Printf("-- Valid JSON --")
				prepareResponse(obj)
			}
		}
	}
}

func constructURL(endpointUrl string, mascot string, location string) string {
	location = url.QueryEscape(location)

	respUrl := strings.Replace(endpointUrl, "{mascot}", mascot, -1)
	respUrl = strings.Replace(respUrl, "{location}", location, -1)

	// fmt.Println(respUrl)
	return respUrl
}

func prepareResponse(data string) {
	// fmt.Println("Recieved Object")

	var pb Postback

	err := json.Unmarshal([]byte(data), &pb)
	if err != nil {
		log.Println("error:", err)
	}
	// fmt.Println(pb)

	for _, data := range pb.Data {
		// fmt.Println(data)
		respUrl := constructURL(pb.Endpoint.Url, data.Mascot, data.Location)
		sendResponse(respUrl)
	}

}

func sendResponse(respUrl string) {
	resp, err := http.Get(respUrl)
	if err != nil {
		log.Println(err)
	}

	log.Printf("(%v) - %v", resp.StatusCode, respUrl)
}
