package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxb-214491577398-jj2xf0vQeEEO624PPgKGEnYo")
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	api.SetDebug(true)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: \n")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)

		case *slack.MessageEvent:
			fmt.Printf("## Message: %+v\n", ev)
      rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", ev.Channel))

		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials\n")
			return

		default:
      fmt.Printf("default msg: %+v", msg.Data)

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
