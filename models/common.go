package models

import "github.com/ghoshRitesh12/brooktube/utils"

type Continuations []struct {
	NextContinuationData struct {
		Continuation string `json:"continuation"`
	} `json:"nextContinuationData"`
}

func (continuations *Continuations) GetContinuationToken() string {
	if len(*continuations) > 0 {
		return (*continuations)[0].NextContinuationData.Continuation
	}
	return ""
}

type Badges []struct {
	MusicInlineBadgeRenderer struct {
		Icon struct {
			IconType string `json:"iconType,omitempty"`
		} `json:"icon,omitempty"`
	} `json:"musicInlineBadgeRenderer,omitempty"`
}

func (badges *Badges) IsExplicit() bool {
	if len(*badges) > 0 {
		return (*badges)[0].
			MusicInlineBadgeRenderer.
			Icon.IconType == utils.MUSIC_EXPLICIT_BADGE
	}
	return false
}
