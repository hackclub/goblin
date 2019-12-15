package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

const GoblinUserID = "@URQ51A2P7"

func main() {
	fmt.Println("Waking up Goblin...")

	goblinToken := os.Getenv("GOBLIN_TOKEN")
	fmt.Println("Connecting with token:", goblinToken)
	if goblinToken == "" {
		log.Fatalf("GOBLIN_TOKEN not found!")
	}

	// init Slack connection
	api := slack.New(goblinToken)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	// init Goblin instance
	goblin := NewClient()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			fmt.Println("Slack - Hello")

		case *slack.ConnectedEvent:
			fmt.Println("Connected:", ev.Info)

		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)
			msg := ev

			if strings.Contains(msg.Text, GoblinUserID) {
				outMsg := rtm.NewOutgoingMessage(
					goblin.Respond(msg.Text),
					msg.Channel,
				)
				rtm.SendMessage(outMsg)
			}

		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:
			fmt.Printf("Unexpected event: %v\n", msg.Data)
		}
	}
}
