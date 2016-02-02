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

	a := GetAttachments(r)

	assert.Equal(t, "horenso Reporter", a[0].Fallback)
}

func TestGetSlackChatPostMessageOpt(t *testing.T) {
	f, _ := os.Open("../fixtures/report_exit_0.json")
	jsonBytes, _ := ioutil.ReadAll(f)

	var r horenso.Report
	json.Unmarshal(jsonBytes, &r)

	opts := GetSlackChatPostMessageOpt(r)

	assert.Equal(t, "slack.ChatPostMessageOpt", reflect.TypeOf(opts).String())
	assert.Equal(t, "[]*slack.Attachment", reflect.TypeOf(opts.Attachments).String())
	assert.Equal(t, "#00ff00", opts.Attachments[0].Color)
}
