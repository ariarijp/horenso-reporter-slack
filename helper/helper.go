package helper

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Songmu/horenso"
	"github.com/bluele/slack"
)

// Getenvs get environment varibles
func Getenvs() (string, string, string, []string) {
	token, channelName, groupName := os.Getenv("HRS_SLACK_TOKEN"), os.Getenv("HRS_SLACK_CHANNEL"), os.Getenv("HRS_SLACK_GROUP")

	if len(token) == 0 {
		panic("HRS_SLACK_TOKEN environment variable is required.")
	} else if len(channelName) == 0 && len(groupName) == 0 {
		panic("HRS_SLACK_CHANNEL or HRS_SLACK_GROUP environment variable is required.")
	}

	itemsStr := os.Getenv("HRS_SLACK_ITEMS")
	var items []string
	if len(itemsStr) > 0 {
		items = strings.Split(itemsStr, ",")
	} else {
		items = []string{"all"}
	}

	return token, channelName, groupName, items
}

// GetReport get horenso report via STDIN
func GetReport(f *os.File) horenso.Report {
	jsonBytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var r horenso.Report
	json.Unmarshal(jsonBytes, &r)

	return r
}

// GetID get Slack channel ID or group ID
func GetID(api *slack.Slack, r horenso.Report, channelName string, groupName string) string {
	if len(channelName) > 0 {
		return GetChannelID(api, r, channelName)
	} else if len(groupName) > 0 {
		return GetGroupID(api, r, groupName)
	}

	panic("Could not resolve ID.")
}

// GetMessage get message
func GetMessage(r horenso.Report) string {
	if *r.ExitCode == 0 {
		return ""
	}

	return "<!channel>"
}

// GetGroupID get Slack group ID by group name
func GetGroupID(api *slack.Slack, r horenso.Report, groupName string) string {
	group, err := api.FindGroupByName(groupName)
	if err != nil {
		panic(err)
	}

	return group.Id
}

// GetChannelID get Slack channel ID by channel name
func GetChannelID(api *slack.Slack, r horenso.Report, channelName string) string {
	channel, err := api.FindChannelByName(channelName)
	if err != nil {
		panic(err)
	}

	return channel.Id
}
