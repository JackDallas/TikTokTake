# TikTokTake

This repo is 3 things:

- A set of libraries for interacting and scraping tiktok `pkg/` `internals/`
- A cli for downloading tiktok data `cmd/cli`
- An irc bot to facilitate adding jobs to scrape tiktok `cmd/ircserver`

## TODO

Todo's in order of priority

- Dockerfile
- Add a queue system and db
- Add more things you can download (hashtags, sounds..)
- Replace the `libs/tiktok-signature` submodule with something either native or a v8 instance inside the program that can generate the signatures

- Generalising the requests made in `user.go` and `request.go`
