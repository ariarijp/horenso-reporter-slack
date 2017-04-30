package helper

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Songmu/horenso"
	"github.com/antonholmquist/jason"
	"github.com/bluele/slack"
)

// Getenvs get environment varibles
func Getenvs() (string, string, string, string, []string) {
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

	return token, channelName, groupName, mention, items
}

// GetReport get horenso report via STDIN
func GetReport(f *os.File) horenso.Report {
	v, err := jason.NewObjectFromReader(f)
	if err != nil {
		panic(err)
	}

	var r horenso.Report
	r.Command, _ = v.GetString("command")
	r.CommandArgs, _ = v.GetStringArray("commandArgs")
	r.Tag, _ = v.GetString("tag")
	r.Output, _ = v.GetString("output")
	r.Stdout, _ = v.GetString("stdout")
	r.Stderr, _ = v.GetString("stderr")
	r.Command, _ = v.GetString("command")
	r.Result, _ = v.GetString("result")
	r.Hostname, _ = v.GetString("hostname")
	exitCode, _ := getInt(v, "exitCode")
	r.ExitCode = &exitCode
	pid, _ := getInt(v, "pid")
	r.Pid = &pid
	r.StartAt, _ = getTime(v, "startAt")
	r.EndAt, _ = getTime(v, "endAt")
	systemTime, _ := v.GetFloat64("systemTime")
	r.SystemTime = &systemTime
	userTime, _ := v.GetFloat64("userTime")
	r.UserTime = &userTime

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

func getInt(v *jason.Object, key string) (int, error) {
	n, err := v.GetInt64(key)
	if err != nil {
		return 0, err
	}

	return int(n), nil
}

func getTime(v *jason.Object, key string) (*time.Time, error) {
	s, err := v.GetString(key)
	if err != nil {
		return nil, err
	}
	t, err := time.Parse("2006-01-02T15:04:05.999999-07:00", s)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
