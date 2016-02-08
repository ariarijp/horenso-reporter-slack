package reporter

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/Songmu/horenso"
	"github.com/stretchr/testify/assert"
)

func TestGetAttachments(t *testing.T) {
	f, _ := os.Open("../fixtures/report_exit_0.json")
	jsonBytes, _ := ioutil.ReadAll(f)

	var r horenso.Report
	json.Unmarshal(jsonBytes, &r)

	items := []string{"all"}
	a := GetAttachments(r, items)
	assert.Equal(t, 13, len(a[0].Fields))
	assert.Equal(t, "command exited with code: 0", a[0].Fields[0].Value)
	assert.Equal(t, "1\n95030\n", a[0].Fields[1].Value)

	items = []string{"Output"}
	a = GetAttachments(r, items)
	assert.Equal(t, 1, len(a[0].Fields))
	assert.Equal(t, "1\n95030\n", a[0].Fields[0].Value)

	items = []string{"Output", "ExitCode"}
	a = GetAttachments(r, items)
	assert.Equal(t, 2, len(a[0].Fields))
	assert.Equal(t, "1\n95030\n", a[0].Fields[0].Value)
	assert.Equal(t, "0", a[0].Fields[1].Value)
}

func TestGetSlackChatPostMessageOpt(t *testing.T) {
	f, _ := os.Open("../fixtures/report_exit_0.json")
	jsonBytes, _ := ioutil.ReadAll(f)

	var r horenso.Report
	json.Unmarshal(jsonBytes, &r)

	items := []string{"all"}
	opts := GetSlackChatPostMessageOpt(r, items)

	assert.Equal(t, "slack.ChatPostMessageOpt", reflect.TypeOf(opts).String())
	assert.Equal(t, "[]*slack.Attachment", reflect.TypeOf(opts.Attachments).String())
	assert.Equal(t, "#00ff00", opts.Attachments[0].Color)
}

func TestIsSelectedItem(t *testing.T) {
	assert.True(t, IsSelectedItem("ExitCode", []string{"all"}))
	assert.True(t, IsSelectedItem("ExitCode", []string{"ExitCode", "Output"}))
	assert.False(t, IsSelectedItem("Stdout", []string{"ExitCode", "Output"}))
	assert.False(t, IsSelectedItem("ExitCode", []string{}))
}
