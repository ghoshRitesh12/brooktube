package playlist

import (
	"github.com/ghoshRitesh12/brooktube/models"
	"github.com/ghoshRitesh12/brooktube/models/search"
)

type APIResp struct {
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			SecondaryContents struct {
				SectionListRenderer struct {
					Contents [1]struct {
						MusicPlaylistShelfRenderer apiRespSection `json:"musicPlaylistShelfRenderer,omitempty"`
					} `json:"contents,omitempty"`

					Continuations models.Continuations `json:"continuations,omitempty"`
				} `json:"sectionListRenderer,omitempty"`
			} `json:"secondaryContents,omitempty"`

			Tabs [1]struct {
				TabRenderer struct {
					Content struct {
						SectionListRenderer struct {
							Contents [1]apiRespHeader `json:"contents,omitempty"`
							//
						} `json:"sectionListRenderer,omitempty"`
					} `json:"content,omitempty"`
				} `json:"tabRenderer,omitempty"`
			} `json:"tabs,omitempty"`
		} `json:"twoColumnBrowseResultsRenderer,omitempty"`
	} `json:"contents,omitempty"`

	Background models.Thumbnail `json:"background,omitempty"`
}

type apiRespSection struct {
	Contents           []apiRespSectionContent `json:"contents,omitempty"`
	PlaylistId         string                  `json:"playlistId,omitempty"`
	CollapsedItemCount int                     `json:"collapsedItemCount,omitempty"`
	Continuations      models.Continuations    `json:"continuations,omitempty"`
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

		MusicItemRendererDisplayPolicy models.DisplayPolicy `json:"musicItemRendererDisplayPolicy,omitempty"`

		Thumbnail models.Thumbnail `json:"thumbnail,omitempty"`
	} `json:"musicResponsiveListItemRenderer,omitempty"`
}

type apiRespHeader struct {
	MusicResponsiveHeaderRenderer struct {
		Title struct {
			Runs models.BasicRuns `json:"runs,omitempty"`
		} `json:"title,omitempty"`

		Subtitle struct {
			Runs models.BasicRuns `json:"runs,omitempty"`
		} `json:"subtitle,omitempty"`

		StraplineTextOne struct {
			Runs models.NavigationEndpointRuns `json:"runs,omitempty"`
		} `json:"straplineTextOne,omitempty"`

		Description struct {
			MusicDescriptionShelfRenderer struct {
				Description struct {
					Runs models.BasicRuns `json:"runs,omitempty"`
				} `json:"description,omitempty"`
			} `json:"musicDescriptionShelfRenderer,omitempty"`
		} `json:"description,omitempty"`

		SecondSubtitle struct {
			Runs models.BasicRuns `json:"runs,omitempty"`
		} `json:"secondSubtitle,omitempty"`
	} `json:"musicResponsiveHeaderRenderer,omitempty"`
}

type APIRespContinuation struct {
	// for continuation data
	ContinuationContents continuationContents `json:"continuationContents,omitempty"`
}

type continuationContents struct {
	MusicPlaylistShelfContinuation struct {
		Contents []apiRespSectionContent `json:"contents,omitempty"`
	} `json:"musicPlaylistShelfContinuation,omitempty"`
}
