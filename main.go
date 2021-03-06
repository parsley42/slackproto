package main

import (
	"bytes"
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

func msgBlockFixed(msg string) slack.MsgOption {
	text := fmt.Sprintf("```%s```", msg)
	messageText := slack.NewTextBlockObject("mrkdwn", text, false, false)
	messageSection := slack.NewSectionBlock(messageText, nil, nil)
	return slack.MsgOptionBlocks(messageSection)
}

func sendBlock(channel string, client slack.Client, block slack.MsgOption) {
	channelID, timestamp, err := client.PostMessage(channel, slack.MsgOptionText("the robot sent you a message", false), block)
	if err != nil {
		fmt.Printf("Could not send message: %v\n", err)
	} else {
		fmt.Printf("Message send successfully to channel %s at %s\n", channelID, timestamp)
	}
}

func sendUserRaw(user, channel, msg string, client slack.Client) {
	text := fmt.Sprintf("%s %s", user, msg)
	block := msgBlockRaw(text)
	sendBlock(channel, client, block)
}

func sendUserPlain(user, channel, msg string, client slack.Client) {
	userMsg := msgBlockRaw(user)
	sendBlock(channel, client, userMsg)
	block := msgBlockPlain(msg)
	sendBlock(channel, client, block)
}

func sendUserFixed(user, channel, msg string, client slack.Client) {
	sbytes := []byte(msg)
	sbytes = bytes.Replace(sbytes, []byte("&"), []byte("&amp;"), -1)
	sbytes = bytes.Replace(sbytes, []byte("<"), []byte("&lt;"), -1)
	sbytes = bytes.Replace(sbytes, []byte(">"), []byte("&gt;"), -1)
	rmsg := string(sbytes)
	text := fmt.Sprintf("%s %s", user, rmsg)
	block := msgBlockFixed(text)
	sendBlock(channel, client, block)
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
	text := "Have a look:\n*<https://github.com/parsley42|David Parsley - Coder>*"
	sendUserRaw(user, channel, text, *client)
	sendUserPlain(user, channel, text, *client)
	sendUserFixed(user, channel, text, *client)
}
