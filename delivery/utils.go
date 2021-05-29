// Michael Peters
// utils.go
// utilities for the Kochava Postback Delivery Agent
// Last Modified: 05/28/21 19:00 PDT

package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// print information when program starts (to be seen in docker logs)
func greet(name string, author string) {
	fmt.Printf("%v, by %v\n", name, author)
	fmt.Println("Starting process...")
}

// check to ensure that JSON coming out of Redis is valid (it should be)
func isValidJSON(input string) bool {
	return json.Valid([]byte(input))
}

// simple solution for replacing the substrings in the endpoint (could be improved)
func constructURL(endpointUrl string, mascot string, location string) string {
	location = url.QueryEscape(location)

	respUrl := strings.Replace(endpointUrl, "{mascot}", mascot, -1)
	respUrl = strings.Replace(respUrl, "{location}", location, -1)

	bracketRegex := `{+[a-z0-9]+}`
	defaultValue := "default"

	matched, _ := regexp.MatchString(bracketRegex, respUrl)
	if matched {
		re := regexp.MustCompile(bracketRegex)
		respUrl = re.ReplaceAllString(respUrl, defaultValue)
	}

	return respUrl
}
