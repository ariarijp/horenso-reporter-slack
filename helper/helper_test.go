package helper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func resetEnvs() {
	os.Setenv("HRS_SLACK_TOKEN", "")
	os.Setenv("HRS_SLACK_CHANNEL", "")
	os.Setenv("HRS_SLACK_GROUP", "")
	os.Setenv("HRS_SLACK_MENTION", "")
}

func TestGetenvs(t *testing.T) {
	func() {
		defer func() {
			err := recover()
			if err != nil {
				assert.Equal(t, "HRS_SLACK_TOKEN environment variable is required.", err)
			} else {
				t.Fail()
			}
		}()

		resetEnvs()
		token, _, _, _, _ := Getenvs()
		if token == "" {
			t.Fail()
		}
	}()

	func() {
		defer func() {
			err := recover()
			if err != nil {
				assert.Equal(t, "HRS_SLACK_CHANNEL or HRS_SLACK_GROUP environment variable is required.", err)
			} else {
				t.Fail()
			}
		}()

		resetEnvs()
		os.Setenv("HRS_SLACK_TOKEN", "token")
		token, _, _, _, _ := Getenvs()
		if token == "" {
			t.Fail()
		}
	}()

	func() {
		resetEnvs()
		os.Setenv("HRS_SLACK_TOKEN", "token")
		os.Setenv("HRS_SLACK_CHANNEL", "channel")
		os.Setenv("HRS_SLACK_GROUP", "group")
		os.Setenv("HRS_SLACK_MENTION", "here")

		token, channelName, groupName, mention, items := Getenvs()

		assert.Equal(t, "token", token)
		assert.Equal(t, "channel", channelName)
		assert.Equal(t, "group", groupName)
		assert.Equal(t, "here", mention)
		assert.Equal(t, []string{"all"}, items)
	}()
}

func TestGetReport(t *testing.T) {
	func() {
		f, _ := os.Open("../fixtures/report_exit_0.json")
		r := GetReport(f)
		assert.Equal(t, 0, *r.ExitCode)
	}()

	func() {
		f, _ := os.Open("../fixtures/report_exit_1.json")
		r := GetReport(f)
		assert.Equal(t, 1, *r.ExitCode)
	}()
}
