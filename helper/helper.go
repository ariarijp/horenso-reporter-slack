package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Songmu/horenso"
	"github.com/bluele/slack"
)

// Getenvs get environment varibles
func Getenvs() (string, string, string, string, []string, bool) {
	token := os.Getenv("HRS_SLACK_TOKEN")
	channelName := os.Getenv("HRS_SLACK_CHANNEL")
	groupName := os.Getenv("HRS_SLACK_GROUP")
	mention := os.Getenv("HRS_SLACK_MENTION")
	if len(mention) == 0 {
		mention = "channel"
	}

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

	notifyEverythingEnv := os.Getenv("HRS_SLACK_NOTIFY_EVERYTHING")
	notifyEverything := true
	if len(notifyEverythingEnv) != 0 && notifyEverythingEnv == "0" {
		notifyEverything = false
	}

	return token, channelName, groupName, mention, items, notifyEverything
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
func GetMessage(r horenso.Report, mention string) string {
	if *r.ExitCode == 0 {
		return ""
	}

	return fmt.Sprintf("<!%s>", mention)
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
