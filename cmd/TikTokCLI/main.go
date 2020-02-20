package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tiktok "github.com/JackDallas/TikTokTake/pkg/tiktok"
	warc "github.com/JackDallas/TikTokTake/pkg/warc"
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
		warcAllURLs(urls, user)
	} else {
		fmt.Println("Need 1 argument, username of account")
		os.Exit(1)
	}
}

func warcAllURLs(reqs []*http.Request, user tiktok.User) {
	warc := warc.WARC{user.Username, fmt.Sprintf("%s-%v", user.Username, time.Now().Unix())}
	errCount := 0

	log.Debugf("Found %v videos for %s sending to downloader, Job Name: %s\n", len(reqs), user.Username, warc.WARCName)

	err := warc.MakeRequest(user.ProfileRequest)
	if err != nil {
		errCount++
		log.Error(err.Error())
	}

	for _, req := range reqs {
		err = warc.MakeRequest(req)
		if err != nil {
			errCount++
			log.Error(err.Error())
		}
	}

	if errCount > 0 {
		log.Debugf("Job %s completed with %v error!\n", warc.WARCName, errCount)
	} else {
		log.Debugf("Job %s completed successfully!\n", warc.WARCName)
	}
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
