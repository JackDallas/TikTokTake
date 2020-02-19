package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tiktok "github.com/JackDallas/TikTokTake/pkg/tiktok"
	wget "github.com/JackDallas/TikTokTake/pkg/wget"
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
		warcAllURLs(urls, user.Username)
	} else {
		fmt.Println("Need 1 argument, username of account")
		os.Exit(1)
	}
}

func warcAllURLs(urls []*http.Request, username string) {
	log.Debugf("Warcing %v videos...\n", len(urls))
	t1 := time.Now()

	for _, url := range urls {
		wget.MakeWARCRequest(url, username, fmt.Sprintf("%s-%v", username, t1.Unix()))
	}

	log.Debugf("Downloaded %v videos in %vs", len(urls), time.Since(t1).Seconds())
}

// func downloadAllURLs(urls []string, username string) {
// 	log.Debugf("Downloading %v videos...\n", len(urls))
// 	t1 := time.Now()
// 	var wg sync.WaitGroup

// 	for i, url := range urls {
// 		wg.Add(1)

// 		go func(url, username string, i int) {
// 			defer wg.Done()

// 			download.Video(url, strconv.FormatInt(int64(i), 10), username)
// 			log.Debugf("Downloaded %v/%v\n", i, len(urls))
// 		}(url, username, i)
// 	}
// 	wg.Wait()
// 	log.Debugf("Downloaded %v videos in %vs", len(urls), time.Since(t1).Seconds())
// }
