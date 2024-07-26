package models

import (
	"github.com/ghoshRitesh12/brooktube/internal/constants"
)

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
		return (*badges)[0].MusicInlineBadgeRenderer.
			Icon.IconType == constants.MUSIC_EXPLICIT_BADGE
	}
	return false
}

type DisplayPolicy string

func (displayPolicy DisplayPolicy) IsDisabled() bool {
	policy := string(displayPolicy)
	return policy == constants.MUSIC_ITEM_RENDERER_DISPLAY_POLICY_GREY_OUT || policy != ""
}

// for getting playlist and album thumbnail
type Thumbnail struct {
	MusicThumbnailRenderer struct {
		Thumbnail struct {
			Thumbnails AppThumbnails `json:"thumbnails,omitempty"`
		} `json:"thumbnail,omitempty"`
	} `json:"musicThumbnailRenderer,omitempty"`
}

func (thmbnl *Thumbnail) GetThumbnail(index uint8) string {
	url := ""
	defaultIndex := 0
	thumbnails := thmbnl.MusicThumbnailRenderer.Thumbnail.Thumbnails

	if len(thumbnails) == 0 {
		return url
	}

	if len(thumbnails) > 0 && int(index) < len(thumbnails) {
		url = thumbnails[index].URL
	} else if len(thumbnails) > 0 && int(index) > len(thumbnails) {
		url = thumbnails[defaultIndex].URL
	}

	return url
}

func (thmbnl *Thumbnail) GetAllThumbnails() AppThumbnails {
	return thmbnl.MusicThumbnailRenderer.Thumbnail.Thumbnails
}
