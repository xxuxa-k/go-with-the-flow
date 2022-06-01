package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	// "net/http"
	"google.golang.org/api/calendar/v3"
	// "google.golang.org/api/option"
	// "google.golang.org/api/transport"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env can't be loaded.")
		return
	}

	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx)
	if err != nil {
		fmt.Println("Calendar service initialize error.")
	}
	eventList := calendarService.Events.List(os.Getenv("CALENDAR_ID"))
	fmt.Println(eventList)
}
