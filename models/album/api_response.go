package album

import (
	"github.com/ghoshRitesh12/brooktube/models"
	"github.com/ghoshRitesh12/brooktube/models/search"
)

type APIResp struct {
	Contents struct {
		SingleColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Content struct {
						SectionListRenderer struct {
							Contents [1]struct {
								MusicShelfRenderer struct {
									Contents []apiRespSectionContent `json:"contents,omitempty"`
								} `json:"musicShelfRenderer,omitempty"`
								//
							} `json:"contents,omitempty"`
						} `json:"sectionListRenderer,omitempty"`
					} `json:"content,omitempty"`
				} `json:"tabRenderer,omitempty"`
			} `json:"tabs,omitempty"`
		} `json:"singleColumnBrowseResultsRenderer,omitempty"`
	} `json:"contents,omitempty"`

	// for meta data of the album
	Header apiRespHeader `json:"header,omitempty"`
}

// for songs
type apiRespSectionContent struct {
	MusicResponsiveListItemRenderer struct {
		FlexColumns []search.ApiRespFlexColumns `json:"flexColumns"`

		FixedColumns []struct {
			MusicResponsiveListItemFixedColumnRenderer struct {
				Text struct {
					Runs models.BasicRuns `json:"runs,omitempty"`
				} `json:"text,omitempty"`
			} `json:"musicResponsiveListItemFixedColumnRenderer,omitempty"`
		} `json:"fixedColumns,omitempty"`

		Badges models.Badges `json:"badges,omitempty"`

		PlaylistItemData struct {
			VideoID string `json:"videoId,omitempty"`
		} `json:"playlistItemData,omitempty"`
		//
	} `json:"musicResponsiveListItemRenderer,omitempty"`
}

type apiRespHeader struct {
	MusicDetailHeaderRenderer struct {
		Title struct {
			Runs models.BasicRuns `json:"runs,omitempty"`
		} `json:"title,omitempty"`

		Subtitle struct {
			Runs models.NavigationEndpointRuns `json:"runs,omitempty"`
		} `json:"subtitle,omitempty"`

		Description struct {
			Runs models.BasicRuns `json:"runs,omitempty"`
		} `json:"description,omitempty"`

		SecondSubtitle struct {
			Runs models.BasicRuns `json:"runs,omitempty"`
		} `json:"secondSubtitle,omitempty"`
	} `json:"musicDetailHeaderRenderer,omitempty"`
}
