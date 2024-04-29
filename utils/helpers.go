package utils

import (
	"strings"

	"github.com/ghoshRitesh12/yt_music/types/search"
)

type ParseYtTextParams struct {
	FlexColumnRuns []search.RespFlexColumnRun
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
	} else if len(params.FlexColumnRuns) > 0 {
		for _, val := range params.FlexColumnRuns {
			str.WriteString(val.Text)
		}
	}

	return str.String()
}

func GetWatchUrl(videoId string) string {
	return HOST + "/watch?v=" + videoId
}
