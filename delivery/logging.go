// Michael Peters
// logging.go
// logger initialization for the Kochava Postback Delivery Agent
// Last Modified: 05/28/21 18:50 PDT

package main

import (
	"fmt"
	"log"
	"os"
)

// initialize the log library to write to a file: "logs.txt"
func initializeLogger() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	fmt.Println("Opening logger... (logs.txt)")

}
