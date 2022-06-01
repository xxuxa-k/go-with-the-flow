package main

import (
	// "context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	// "net/http"
	// "google.golang.org/api/calendar/v3"
	// "google.golang.org/api/transport"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env can't be loaded.")
	}
	API_KEY := os.Getenv("API_KEY")
	if API_KEY != "" {
		fmt.Println(API_KEY)
	}
}
