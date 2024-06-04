package helpers

import (
	"strings"

	"github.com/ghoshRitesh12/brooktube/models/search"
)

type ParseYtTextParams struct {
	FlexColumnRuns []search.APIRespFlexColumnRun
	NormalRuns     []struct {
		Text string
	}
}

func ParseYtTextField(params ParseYtTextParams) string {
	var str strings.Builder

	if len(params.NormalRuns) > 0 {
		for _, val := range params.NormalRuns {
			str.WriteString(val.Text)
		}
		return str.String()
	}

	if len(params.FlexColumnRuns) > 0 {
		for _, val := range params.FlexColumnRuns {
			str.WriteString(val.Text)
		}
		return str.String()
	}

	return str.String()
}
