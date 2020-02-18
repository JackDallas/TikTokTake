package main

import (
	"crypto/tls"
	"fmt"
	"strings"

	irc "github.com/thoj/go-ircevent"

	tiktok "github.com/JackDallas/TikTokTake/pkg/tiktok"
)

const channel = "#TikTokBot"
const serverssl = "irc.hackint.org:6697"
const nick = "TikTokBot"

var irccon *irc.Connection

func main() {
	irccon = irc.IRC(nick, nick)
	irccon.VerboseCallbackHandler = true
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })

	irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
		msg := handleMSG(e.Arguments[1])
		if msg != "" {
			irccon.Privmsg(channel, msg)
		}
	})

	err := irccon.Connect(serverssl)

	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}

	irccon.Loop()
}

const helpMessage = "Commands: !user <username> (Archives a TikTok user), !help (Displays the help message)"

func handleMSG(args string) string {
	if len(args) > 0 && args[0] == '!' {

		command := strings.Split(args, " ")

		switch command[0] {
		case "!help":
			return helpMessage

		case "!user":
			if len(command) == 2 {
				go func(username string) {
					User, err := tiktok.NewUser(username)
					if err != nil {
						irccon.Privmsgf(channel, "Error getting user %s, %e", username, err)
					}
					urls, err := User.GetVideos()
					irccon.Privmsgf(channel, "Found %v videos for %s", len(urls), username)
				}(command[1])
				return fmt.Sprintf("Queuing job to archive user: %s", command[1])
			}
			return "Invalid Usage: !user <username>"

		}

		return fmt.Sprintf("Invalid Command %s : %s", command, helpMessage)
	}

	return ""
}
