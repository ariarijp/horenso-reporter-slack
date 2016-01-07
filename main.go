package main

import (
	"io/ioutil"
	"os"

	slackreporter "github.com/ariarijp/horenso-reporter-slack/reporter"
	"github.com/bluele/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	groupName := os.Getenv("SLACK_GROUP")

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	api := slack.New(token)
	r := slackreporter.GetReport(stdin)

	if *r.ExitCode != 0 {
		slackreporter.NotifyToGroup(*api, r, groupName)
	}
}
