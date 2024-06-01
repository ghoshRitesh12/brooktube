package artist

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
							Contents []apiRespSectionContent `json:"contents,omitempty"`
							//
						} `json:"sectionListRenderer,omitempty"`
					} `json:"content,omitempty"`
				} `json:"tabRenderer,omitempty"`
			} `json:"tabs,omitempty"`
		} `json:"singleColumnBrowseResultsRenderer,omitempty"`
	} `json:"contents,omitempty"`

	Header apiRespHeader `json:"header,omitempty"`
}

type apiRespSectionContent struct {
	MusicShelfRenderer            apiRespMusicShelfRenderer         `json:"musicShelfRenderer,omitempty"`
	MusicCarouselShelfRenderer    apiRespMusicCarouselShelfRenderer `json:"musicCarouselShelfRenderer,omitempty"`
	MusicDescriptionShelfRenderer apiRespMusicDescriptionSection    `json:"musicDescriptionShelfRenderer,omitempty"`
}

// for Songs
type apiRespMusicShelfRenderer struct {
	Title struct {
		Runs models.NavigationEndpointRuns `json:"runs,omitempty"`
	} `json:"title,omitempty"`

	Contents []search.APIRespSectionContent `json:"contents,omitempty"`
}

/*
// for Albums, Singles and Videos
NavigationEndpoint struct {
	BrowseEndpoint struct {
		BrowseID string `json:"browseId,omitempty"`

		// determines discography being either single or album
		Params string `json:"params,omitempty"`

		BrowseEndpointContextSupportedConfigs struct {
			BrowseEndpointContextMusicConfig struct {
				PageType string `json:"pageType,omitempty"`
			} `json:"browseEndpointContextMusicConfig,omitempty"`
		} `json:"browseEndpointContextSupportedConfigs,omitempty"`
	} `json:"browseEndpoint,omitempty"`
} `json:"navigationEndpoint,omitempty"`
*/

type apiRespMusicCarouselShelfRenderer struct {
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
				//
			} `json:"navigationEndpoint,omitempty"`
			//
		} `json:"musicTwoRowItemRenderer,omitempty"`
	} `json:"contents,omitempty"`
}

// about sectionListRendererContent
type apiRespMusicDescriptionSection struct {
	Subheader struct {
		Runs models.BasicRuns `json:"runs,omitempty"`
	} `json:"subheader,omitempty"`
}

type apiRespHeader struct {
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
	} `json:"musicImmersiveHeaderRenderer,omitempty"`
}
