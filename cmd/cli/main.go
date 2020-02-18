package main

import (
	"fmt"
	"os"

	tiktok "github.com/JackDallas/TikTokTake/pkg/tiktok"
	log "github.com/sirupsen/logrus"
)

func main() {
	args := os.Args[1:]
	log.SetLevel(log.DebugLevel)

	if len(args) == 1 {
		user, err := tiktok.NewUser(args[0])
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Debug("Found User: " + user.Username)
		urls, err := user.GetVideos()
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, url := range urls {
			log.Println(url)
		}
		log.Printf("Found %v videos for user @%s", len(urls), user.Username)
	} else {
		fmt.Println("Need 1 argument, username of account")
		os.Exit(1)
	}
}
