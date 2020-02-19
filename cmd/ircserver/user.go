package main

import (
	"fmt"
	"time"

	tiktok "github.com/JackDallas/TikTokTake/pkg/tiktok"
	warc "github.com/JackDallas/TikTokTake/pkg/warc"
	log "github.com/sirupsen/logrus"
)

func archiveUser(username string) {
	User, err := tiktok.NewUser(username)
	if err != nil {
		irccon.Privmsgf(channel, "Error getting user %s, %e", username, err)
	}

	reqs, err := User.GetVideos()
	if err != nil {
		irccon.Privmsgf(channel, "Error getting videos for %s, %e", username, err)
	} else {
		warc := warc.WARC{username, fmt.Sprintf("%s-%v", username, time.Now().Unix())}

		irccon.Privmsgf(channel, "Found %v videos for %s sending to downloader, Job Name: %s", len(reqs), username, warc.WARCName)

		errCount := 0

		err = warc.MakeRequest(User.ProfileRequest)
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
			irccon.Privmsgf(channel, "Job %s completed with %v error!", warc.WARCName, errCount)
		} else {
			irccon.Privmsgf(channel, "Job %s completed successfully!", warc.WARCName)
		}
	}
}
