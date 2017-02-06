package main

import (
	"os"

	"github.com/ariarijp/horenso-reporter-slack/helper"
	"github.com/ariarijp/horenso-reporter-slack/reporter"
	"github.com/bluele/slack"
)

func main() {
	token, channelName, groupName, mention, items, notifyEverything := helper.Getenvs()
	r := helper.GetReport(os.Stdin)

	if *r.ExitCode == 0 && !notifyEverything {
		os.Exit(0)
	}

	api := slack.New(token)

	id := helper.GetID(api, r, channelName, groupName)
	msg := helper.GetMessage(r, mention)

	reporter.SendReportToSlack(api, r, id, msg, items)
}
