package reporter

import (
	"fmt"

	"github.com/Songmu/horenso"
	"github.com/bluele/slack"
)

func GetAttachments(r horenso.Report) []*slack.Attachment {
	var attachments []*slack.Attachment

	var a slack.Attachment
	a.Fallback = "horenso Reporter"
	a.AuthorName = "horenso Reporter"
	a.Color = "#ff0000"
	a.Fields = []*slack.AttachmentField{
		&slack.AttachmentField{
			Title: "Result",
			Value: fmt.Sprintf("%v", r.Result),
		},
		&slack.AttachmentField{
			Title: "Output",
			Value: fmt.Sprintf("%v", r.Output),
		},
		&slack.AttachmentField{
			Title: "Stdout",
			Value: fmt.Sprintf("%v", r.Stdout),
		},
		&slack.AttachmentField{
			Title: "Stderr",
			Value: fmt.Sprintf("%v", r.Stderr),
		},
		&slack.AttachmentField{
			Title: "Command",
			Value: fmt.Sprintf("%v", r.Command),
		},
		&slack.AttachmentField{
			Title: "CommandArgs",
			Value: fmt.Sprintf("%v", r.CommandArgs),
		},
		&slack.AttachmentField{
			Title: "Pid",
			Value: fmt.Sprintf("%d", r.Pid),
			Short: true,
		},
		&slack.AttachmentField{
			Title: "ExitCode",
			Value: fmt.Sprintf("%d", *r.ExitCode),
			Short: true,
		},
		&slack.AttachmentField{
			Title: "StartAt",
			Value: fmt.Sprintf("%v", r.StartAt),
			Short: true,
		},
		&slack.AttachmentField{
			Title: "EndAt",
			Value: fmt.Sprintf("%v", r.EndAt),
			Short: true,
		},
		&slack.AttachmentField{
			Title: "Hostname",
			Value: fmt.Sprintf("%v", r.Hostname),
			Short: true,
		},
		&slack.AttachmentField{
			Title: "SystemTime",
			Value: fmt.Sprintf("%f", *r.SystemTime),
			Short: true,
		},
	}

	return append(attachments, &a)
}

func GetSlackChatPostMessageOpt(r horenso.Report) slack.ChatPostMessageOpt {
	return slack.ChatPostMessageOpt{
		Attachments: GetAttachments(r),
	}
}

func NotifyToGroup(api slack.Slack, r horenso.Report, groupName string) {
	group, err := api.FindGroupByName(groupName)
	if err != nil {
		panic(err)
	}

	Notify(api, r, group.Id)
}

func NotifyToChannel(api slack.Slack, r horenso.Report, channelName string) {
	channel, err := api.FindChannelByName(channelName)
	if err != nil {
		panic(err)
	}

	Notify(api, r, channel.Id)
}

func Notify(api slack.Slack, r horenso.Report, id string) {
	opt := GetSlackChatPostMessageOpt(r)

	err := api.ChatPostMessage(id, "", &opt)
	if err != nil {
		panic(err)
	}
}
