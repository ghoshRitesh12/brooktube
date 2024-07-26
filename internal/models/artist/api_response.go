package artist

import (
	"github.com/ghoshRitesh12/brooktube/internal/models"
	"github.com/ghoshRitesh12/brooktube/internal/models/search"
)

type APIResp struct {
	Contents struct {
		SingleColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Content struct {
						SectionListRenderer struct {
							Contents []apiRespSectionContent `json:"contents,omitempty"`
							//
						} `json:"sectionListRenderer,omitempty"`
					} `json:"content,omitempty"`
				} `json:"tabRenderer,omitempty"`
			} `json:"tabs,omitempty"`
		} `json:"singleColumnBrowseResultsRenderer,omitempty"`
	} `json:"contents,omitempty"`

	// for meta data of the album
	Header apiRespHeader `json:"header,omitempty"`
}

type apiRespSectionContent struct {
	// for songs
	MusicShelfRenderer APIRespMusicShelfRenderer `json:"musicShelfRenderer,omitempty"`
	// for albums, singles, videos, featured on & alike artists
	MusicCarouselShelfRenderer APIRespMusicCarouselShelfRenderer `json:"musicCarouselShelfRenderer,omitempty"`
	// for album views
	MusicDescriptionShelfRenderer apiRespMusicDescriptionShelfRenderer `json:"musicDescriptionShelfRenderer,omitempty"`
}

// for songs
type APIRespMusicShelfRenderer struct {
	Title struct {
		Runs models.NavigationEndpointRuns `json:"runs,omitempty"`
	} `json:"title,omitempty"`

	Contents []search.APIRespSectionContent `json:"contents,omitempty"`
}

// for albums, singles, videos, featured on & alike artists
type APIRespMusicCarouselShelfRenderer struct {
	Header struct {
		MusicCarouselShelfBasicHeaderRenderer struct {
			Title struct {
				Runs models.NavigationEndpointParamsRuns `json:"runs,omitempty"`
				//
			} `json:"title,omitempty"`
		} `json:"musicCarouselShelfBasicHeaderRenderer,omitempty"`
	} `json:"header,omitempty"`

	Contents []struct {
		MusicTwoRowItemRenderer struct {
			Title struct {
				// item name
				Runs models.BasicRuns `json:"runs,omitempty"`
			} `json:"title,omitempty"`

			Subtitle struct {
				Runs models.BasicRuns `json:"runs,omitempty"`
			} `json:"subtitle,omitempty"`

			NavigationEndpoint struct {
				BrowseEndpoint struct {
					BrowseID string `json:"browseId,omitempty"`

					BrowseEndpointContextSupportedConfigs struct {
						BrowseEndpointContextMusicConfig struct {
							PageType string `json:"pageType,omitempty"`
						} `json:"browseEndpointContextMusicConfig,omitempty"`
					} `json:"browseEndpointContextSupportedConfigs,omitempty"`
				} `json:"browseEndpoint,omitempty"`

				// for videos
				WatchEndpoint struct {
					VideoID string `json:"videoId,omitempty"`

					WatchEndpointMusicSupportedConfigs struct {
						WatchEndpointMusicConfig struct {
							MusicVideoType string `json:"musicVideoType,omitempty"`
						} `json:"watchEndpointMusicConfig,omitempty"`
					} `json:"watchEndpointMusicSupportedConfigs,omitempty"`
				} `json:"watchEndpoint,omitempty"`
			} `json:"navigationEndpoint,omitempty"`

			ThumbnailRenderer models.Thumbnail `json:"thumbnailRenderer,omitempty"`
		} `json:"musicTwoRowItemRenderer,omitempty"`
	} `json:"contents,omitempty"`
}

// for album views
type apiRespMusicDescriptionShelfRenderer struct {
	Subheader struct {
		Runs models.BasicRuns `json:"runs,omitempty"` // views
	} `json:"subheader,omitempty"`
}

// for meta data of the album
type apiRespHeader struct {
	// for artists
	MusicImmersiveHeaderRenderer struct {
		Title struct {
			Runs models.BasicRuns
		} `json:"title,omitempty"`

		SubscriptionButton struct {
			SubscribeButtonRenderer struct {
				SubscriberCountText struct {
					Runs models.BasicRuns
				} `json:"subscriberCountText,omitempty"`
			} `json:"subscribeButtonRenderer,omitempty"`
		} `json:"subscriptionButton,omitempty"`

		Description struct {
			Runs models.BasicRuns
		} `json:"description,omitempty"`

		Thumbnail models.Thumbnail `json:"thumbnail,omitempty"`
	} `json:"musicImmersiveHeaderRenderer,omitempty"`

	// for non artists
	MusicVisualHeaderRenderer struct {
		Title struct {
			Runs models.BasicRuns
		} `json:"title,omitempty"`

		SubscriptionButton struct {
			SubscribeButtonRenderer struct {
				SubscriberCountText struct {
					Runs models.BasicRuns
				} `json:"subscriberCountText,omitempty"`
			} `json:"subscribeButtonRenderer,omitempty"`
		} `json:"subscriptionButton,omitempty"`

		Description struct {
			Runs models.BasicRuns
		} `json:"description,omitempty"`

		Thumbnail           models.Thumbnail `json:"thumbnail,omitempty"`
		ForegroundThumbnail models.Thumbnail `json:"foregroundThumbnail,omitempty"`
	} `json:"musicVisualHeaderRenderer,omitempty"`
}
