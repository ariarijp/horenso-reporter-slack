package main

import (
	"os"

	"github.com/ariarijp/horenso-reporter-slack/helper"
	"github.com/ariarijp/horenso-reporter-slack/reporter"
	"github.com/bluele/slack"
)

func main() {
	token, channelName, groupName, mention, items := helper.Getenvs()
	r := helper.GetReport(os.Stdin)

	api := slack.New(token)

	id := helper.GetID(api, r, channelName, groupName)

	reporter.SendReportToSlack(api, r, id, mention, items)
}
