package reporter

import (
	"fmt"

	"github.com/Songmu/horenso"
	"github.com/bluele/slack"
)

// GetAttachments get attachments for message
func GetAttachments(r horenso.Report, items []string) []*slack.Attachment {
	var attachments []*slack.Attachment

	var a slack.Attachment
	a.Fallback = "horenso Reporter"
	a.AuthorName = "horenso Reporter"

	if *r.ExitCode == 0 {
		a.Color = "#00ff00"
	} else {
		a.Color = "#ff0000"
	}

	fields := []*slack.AttachmentField{}

	fields = append(fields, &slack.AttachmentField{
		Title: "Result",
		Value: fmt.Sprintf("%v", r.Result),
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "Output",
		Value: fmt.Sprintf("%v", r.Output),
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "Stdout",
		Value: fmt.Sprintf("%v", r.Stdout),
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "Stderr",
		Value: fmt.Sprintf("%v", r.Stderr),
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "Command",
		Value: fmt.Sprintf("%v", r.Command),
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "CommandArgs",
		Value: fmt.Sprintf("%v", r.CommandArgs),
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "Pid",
		Value: fmt.Sprintf("%d", r.Pid),
		Short: true,
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "ExitCode",
		Value: fmt.Sprintf("%d", *r.ExitCode),
		Short: true,
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "StartAt",
		Value: fmt.Sprintf("%v", r.StartAt),
		Short: true,
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "EndAt",
		Value: fmt.Sprintf("%v", r.EndAt),
		Short: true,
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "Hostname",
		Value: fmt.Sprintf("%v", r.Hostname),
		Short: true,
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "SystemTime",
		Value: fmt.Sprintf("%f", *r.SystemTime),
		Short: true,
	})
	fields = append(fields, &slack.AttachmentField{
		Title: "UserTime",
		Value: fmt.Sprintf("%f", *r.UserTime),
		Short: true,
	})

	a.Fields = fields

	return append(attachments, &a)
}

// GetSlackChatPostMessageOpt message options for message
func GetSlackChatPostMessageOpt(r horenso.Report, items []string) slack.ChatPostMessageOpt {
	return slack.ChatPostMessageOpt{
		Attachments: GetAttachments(r, items),
	}
}

// SendReportToSlack send Report to Slack
func SendReportToSlack(api *slack.Slack, r horenso.Report, id string, m string, items []string) {
	opt := GetSlackChatPostMessageOpt(r, items)

	err := api.ChatPostMessage(id, m, &opt)
	if err != nil {
		panic(err)
	}
}
