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

	if IsSelectedItem("Result", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "Result",
			Value: fmt.Sprintf("%v", r.Result),
			Short: true,
		})
	}

	if IsSelectedItem("Output", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "Output",
			Value: fmt.Sprintf("%v", r.Output),
			Short: true,
		})
	}

	if IsSelectedItem("Stdout", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "Stdout",
			Value: fmt.Sprintf("%v", r.Stdout),
			Short: true,
		})
	}

	if IsSelectedItem("Stderr", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "Stderr",
			Value: fmt.Sprintf("%v", r.Stderr),
			Short: true,
		})
	}

	if IsSelectedItem("Command", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "Command",
			Value: fmt.Sprintf("%v", r.Command),
			Short: true,
		})
	}

	if IsSelectedItem("CommandArgs", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "CommandArgs",
			Value: fmt.Sprintf("%v", r.CommandArgs),
			Short: true,
		})
	}

	if IsSelectedItem("Pid", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "Pid",
			Value: fmt.Sprintf("%d", r.Pid),
			Short: true,
		})
	}

	if IsSelectedItem("ExitCode", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "ExitCode",
			Value: fmt.Sprintf("%d", *r.ExitCode),
			Short: true,
		})
	}

	if IsSelectedItem("StartAt", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "StartAt",
			Value: fmt.Sprintf("%v", r.StartAt),
			Short: true,
		})
	}
	if IsSelectedItem("EndAt", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "EndAt",
			Value: fmt.Sprintf("%v", r.EndAt),
			Short: true,
		})
	}
	if IsSelectedItem("Hostname", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "Hostname",
			Value: fmt.Sprintf("%v", r.Hostname),
			Short: true,
		})
	}
	if IsSelectedItem("SystemTime", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "SystemTime",
			Value: fmt.Sprintf("%f", *r.SystemTime),
			Short: true,
		})
	}
	if IsSelectedItem("UserTime", items) {
		fields = append(fields, &slack.AttachmentField{
			Title: "UserTime",
			Value: fmt.Sprintf("%f", *r.UserTime),
			Short: true,
		})
	}

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

// IsSelectedItem returns key exists in slice
func IsSelectedItem(a string, list []string) bool {
	if len(list) == 0 {
		return false
	}

	if list[0] == "all" {
		return true
	}

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
