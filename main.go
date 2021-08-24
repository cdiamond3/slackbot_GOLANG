package main

import (
	"fmt"
	"os"
	"github.com/nlopes/slack"
)

var (
	slackClient *slack.Client
)

func main() {

	slackClient := slack.New(os.Getenv("SLACK_ACCESS_TOKEN"))
	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			go handleMessage(ev)
		}
	}

}

func handleMessage(ev *slack.MessageEvent) {
	fmt.Printf("%v\n", ev)
}
