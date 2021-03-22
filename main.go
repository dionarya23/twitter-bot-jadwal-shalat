package main

import (
	"fmt"
	"log"

	"github.com/dionarya23/twitter-bot-jadwal-shalat/twitter"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Go-Twitter Jadwal Shalat Reminder Bot v0.01")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := gocron.NewScheduler()
	s.Every(1).Minutes().Do(twitter.UpdateStatus)
	<-s.Start()
}
