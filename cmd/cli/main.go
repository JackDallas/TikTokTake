package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

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
		log.Printf("Found %v videos for user @%s", len(urls), user.Username)
		downloadAllURLs(urls, user.Username)
	} else {
		fmt.Println("Need 1 argument, username of account")
		os.Exit(1)
	}
}

func downloadAllURLs(urls []string, username string) {
	log.Debugf("Downloading %v videos...\n", len(urls))
	t1 := time.Now()
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)

		go func(url, username string, i int) {
			defer wg.Done()

			tiktok.DownloadVideo(url, strconv.FormatInt(int64(i), 10), username)
			log.Debugf("Downloaded %v/%v\n", i, len(urls))
		}(url, username, i)
	}
	wg.Wait()
	log.Debugf("Downloaded %v videos in %vs", len(urls), time.Since(t1).Seconds())
}
