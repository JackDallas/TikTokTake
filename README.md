# TikTokTake

This repo is 3 things; a set of libraries for interacting and scraping tiktok `pkg/` `internals/`, a cli for downloading tiktok data `cmd/cli`, and an irc bot to facilitate adding jobs to scrape tiktok `cmd/ircserver`

## TODO

Todo's in order of priority

- Downloading Videos
- Replace the `libs/tiktok-signature` submodule with something either native or a v8 instance inside the program that can generate the signatures

- Generalising the requests made in `user.go` and `httptiktok.go`
