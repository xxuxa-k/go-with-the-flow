package main

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	config := NewGoogleAuthConf()
	fmt.Println(config)
}

func NewGoogleAuthConf() *oauth2.Config {
	json, err := ioutil.ReadFile("./client_secret.json")
	if err != nil {
		fmt.Printf("Err: %v", err)
	}
	config, err := google.ConfigFromJSON(json)
	if err != nil {
		fmt.Printf("Err: %v", err)
	}

	return config
}
