package helper

import (
	"os"
	"testing"

	"github.com/Songmu/horenso"
	"github.com/stretchr/testify/assert"
)

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

func TestGetMessage(t *testing.T) {
	var r horenso.Report

	exitCode := 0
	r.ExitCode = &exitCode
	assert.Equal(t, "<!here>", GetMessage(r))

	exitCode = 1
	assert.Equal(t, "<!channel>", GetMessage(r))
}
