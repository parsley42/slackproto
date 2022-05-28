package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func msgBlockRaw(msg string) slack.MsgOption {
	messageText := slack.NewTextBlockObject("mrkdwn", msg, false, false)
	messageSection := slack.NewSectionBlock(messageText, nil, nil)
	return slack.MsgOptionBlocks(messageSection)
}

func msgBlockPlain(msg string) slack.MsgOption {
	messageText := slack.NewTextBlockObject("plain_text", msg, false, false)
	messageSection := slack.NewSectionBlock(messageText, nil, nil)
	return slack.MsgOptionBlocks(messageSection)
}

func sendBlock(channel string, client slack.Client, block slack.MsgOption) {
	channelID, timestamp, err := client.PostMessage(channel, slack.MsgOptionText("", false), block)
	if err != nil {
		fmt.Printf("Could not send message: %v\n", err)
	} else {
		fmt.Printf("Message send successfully to channel %s at %s\n", channelID, timestamp)
	}
}

func main() {
	token, ok := os.LookupEnv("SLACK_TOKEN")
	if !ok {
		fmt.Println("Missing SLACK_TOKEN in environment")
		os.Exit(1)
	}
	channel, ok := os.LookupEnv("SLACK_CHANNEL")
	if !ok {
		fmt.Println("Missing SLACK_CHANNEL in environment")
		os.Exit(1)
	}
	user, ok := os.LookupEnv("SLACK_USER")
	if !ok {
		fmt.Println("Missing SLACK_CHANNEL in environment")
		os.Exit(1)
	}
	client := slack.New(
		token,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)
	text := fmt.Sprintf("%s - Have a look:\n*<https://github.com/parsley42|David Parsley - Coder>*", user)
	rawMessage := msgBlockRaw(text)
	sendBlock(channel, *client, rawMessage)
	plainMessage := msgBlockPlain(text)
	sendBlock(channel, *client, plainMessage)
}
