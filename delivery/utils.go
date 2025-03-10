// Michael Peters
// utils.go
// utilities for the Kochava Postback Delivery Agent
// Last Modified: 05/29/21 18:15 PDT

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
	// important to URL encode the endpoint (to ensure final URL is valid)
	location = url.QueryEscape(location)

	// replacing the predetermined items
	respUrl := strings.Replace(endpointUrl, "{mascot}", mascot, -1)
	respUrl = strings.Replace(respUrl, "{location}", location, -1)

	// regular expression that matches anything (alphanum) surrounded by curly brackets
	bracketRegex := `{+[a-z0-9]+}`
	// default value to replace the brackets with if there is an unmatched key
	defaultValue := "default"

	// check string for remaining bracket keys
	matched, _ := regexp.MatchString(bracketRegex, respUrl)
	if matched {
		re := regexp.MustCompile(bracketRegex)
		// replace all unmatched bracket keys with the default value
		respUrl = re.ReplaceAllString(respUrl, defaultValue)
	}

	return respUrl
}
