package main

import (
	"io/ioutil"
	"os"

	slackreporter "github.com/ariarijp/horenso-reporter-slack/reporter"
	"github.com/bluele/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	channelName := os.Getenv("SLACK_CHANNEL")
	groupName := os.Getenv("SLACK_GROUP")

	if len(token) == 0 {
		panic("SLACK_TOKEN environment variable is required.")
	} else if len(channelName) == 0 && len(groupName) == 0 {
		panic("SLACK_CHANNEL or SLACK_GROUP environment variable is required.")
	}

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	api := slack.New(token)
	r := slackreporter.GetReport(stdin)

	if *r.ExitCode != 0 {
		if len(channelName) > 0 {
			slackreporter.NotifyToChannel(*api, r, channelName)
		}
		if len(groupName) > 0 {
			slackreporter.NotifyToGroup(*api, r, groupName)
		}
	}
}
