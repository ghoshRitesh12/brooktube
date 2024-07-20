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
		return (*badges)[0].MusicInlineBadgeRenderer.
			Icon.IconType == utils.MUSIC_EXPLICIT_BADGE
	}
	return false
}

type DisplayPolicy string

func (displayPolicy DisplayPolicy) IsDisabled() bool {
	policy := string(displayPolicy)
	return policy == utils.MUSIC_ITEM_RENDERER_DISPLAY_POLICY_GREY_OUT || policy != ""
}

// for getting playlist and album thumbnail
type Thumbnail struct {
	MusicThumbnailRenderer struct {
		Thumbnail struct {
			Thumbnails []struct {
				URL    string `json:"url,omitempty"`
				Width  int    `json:"width,omitempty"`
				Height int    `json:"height,omitempty"`
			} `json:"thumbnails,omitempty"`
		} `json:"thumbnail,omitempty"`
	} `json:"musicThumbnailRenderer,omitempty"`
}

func (thmbnl *Thumbnail) GetThumbnail(index uint8) string {
	url := ""
	thumbnails := thmbnl.MusicThumbnailRenderer.Thumbnail.Thumbnails

	if len(thumbnails) == 0 {
		return url
	}
	if len(thumbnails) > 0 && int(index) < len(thumbnails) {
		url = thumbnails[index].URL
	}

	return url
}
