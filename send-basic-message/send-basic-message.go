package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelID, timestamp, err := api.PostMessage(
		"C030K5QDZ1S",
		slack.MsgOptionText("Hello World!", false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message sent successfully to channel %s at %s", channelID, timestamp)
}
