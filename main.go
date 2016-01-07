package main

import (
	"os"

	slackreporter "github.com/ariarijp/horenso-reporter-slack/reporter"
	"github.com/bluele/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	groupName := os.Getenv("SLACK_GROUP")

	api := slack.New(token)
	r := slackreporter.GetReport([]byte(os.Args[1]))

	slackreporter.NotifyToGroup(*api, r, groupName)
}
