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

	r := slackreporter.GetReport(stdin)

	var id, m string
	api := slack.New(token)

	if len(channelName) > 0 {
		id = slackreporter.GetChannelId(*api, r, channelName)
		m = "<!channel>"
	} else if len(groupName) > 0 {
		id = slackreporter.GetGroupId(*api, r, groupName)
		m = "<!group>"
	}

	if *r.ExitCode != 0 {
		slackreporter.Notify(*api, r, id, m)
	} else {
		slackreporter.Notify(*api, r, id, "<!here>")
	}
}
