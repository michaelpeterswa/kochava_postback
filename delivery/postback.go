// Michael Peters
// postback.go
// Postback struct and handling for the Kochava Postback Delivery Agent
// Last Modified: 05/28/21 18:50 PDT

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

// https://stackoverflow.com/questions/36322141/unmarshal-inconsistent-json
// https://github.com/Jeffail/gabs
// https://bencane.com/2020/12/08/maps-vs-structs-for-json/

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

func dequeueLoop(redisConn *redis.Client, key string) {
	// start infinite loop
	for {
		// one-second timeout on empty
		val, err := redisConn.BRPop(ctx, 1000000000, key).Result()
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
				start := time.Now()
				prepareResponse(obj, start)
			}
		}
	}
}

func prepareResponse(data string, start time.Time) {
	var pb Postback

	err := json.Unmarshal([]byte(data), &pb)
	if err != nil {
		log.Println("error:", err)
	}

	for _, data := range pb.Data {

		respUrl := constructURL(pb.Endpoint.Url, data.Mascot, data.Location)
		sendResponse(respUrl, start)
	}
}

func sendResponse(respUrl string, start time.Time) {
	resp, err := http.Get(respUrl)
	if err != nil {
		log.Println(err)
	} else {
		t := time.Now()
		elapsed := t.Sub(start)
		log.Printf("(%v) - time: %v - %v", resp.StatusCode, elapsed, respUrl)
		fmt.Printf("(%v) - time: %v - %v\n", resp.StatusCode, elapsed, respUrl)
	}
}
