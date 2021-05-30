// Michael Peters
// main.go
// runner for the Kochava Postback Delivery Agent
// Last Modified: 05/29/21 17:45 PDT

package main

import (
	"context"
)

var ctx = context.Background()

func main() {

	// set up initial values (most importantly queueKey)
	name := "Delivery Agent v1.1.0"
	author := "Michael Peters"
	queueKey := "postback_queue"

	// greet the command line
	greet(name, author)

	// get logging running
	initializeLogger()

	// connect to Redis
	rdb := redisConnect()

	// start infinite loop to pull new objects from Redis until it is empty
	// then while empty, wait timeout time and check again
	dequeueLoop(rdb, queueKey)

}
