package reporter

import (
	"encoding/json"

	"github.com/Songmu/horenso"
)

func GetReport(jsonBytes []byte) horenso.Report {
	var r horenso.Report
	json.Unmarshal(jsonBytes, &r)

	return r
}
