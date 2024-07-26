package search

import "github.com/ghoshRitesh12/brooktube/internal/models"

type APIResp struct {
	Contents struct {
		TabbedSearchResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Content struct {
						SectionListRenderer struct {
							// for normal search results
							Contents []apiRespSection `json:"contents"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"tabbedSearchResultsRenderer"`
	} `json:"contents,omitempty"`

	// for paginated search results
	ContinuationContents apiRespSectionContinuation `json:"continuationContents,omitempty"`
}

// for normal search results
type apiRespSection struct {
	// for community playlists, songs, albums, videos, etc.
	MusicShelfRenderer struct {
		Title struct {
			Runs models.BasicRuns `json:"runs"`
		} `json:"title"`

		Contents      []APIRespSectionContent `json:"contents"`
		Continuations models.Continuations    `json:"continuations,omitempty"`
	} `json:"musicShelfRenderer,omitempty"`
}

// for paginated search results
type apiRespSectionContinuation struct {
	// for community playlists, songs, albums, videos, etc.
	MusicShelfContinuation struct {
		Contents      []APIRespSectionContent `json:"contents"`
		Continuations models.Continuations    `json:"continuations,omitempty"`
	} `json:"musicShelfContinuation,omitempty"`
}

type APIRespSectionContent struct {
	MusicResponsiveListItemRenderer struct {
		FlexColumns []ApiRespFlexColumns `json:"flexColumns"`

		Menu struct {
			MenuRenderer struct {
				Items []struct {
					MenuNavigationItemRenderer struct {
						Text struct {
							Runs []struct {
								Text string
							} `json:"runs"`
						} `json:"text"`

						NavigationEndpoint struct {
							WatchPlaylistEndpoint struct {
								PlaylistID string `json:"playlistId"`
								Params     string `json:"params"`
							} `json:"watchPlaylistEndpoint"`
						} `json:"navigationEndpoint"`
					} `json:"menuNavigationItemRenderer,omitempty"`
				} `json:"items"`
			} `json:"menuRenderer"`
		} `json:"menu"`

		PlaylistItemData struct {
			VideoId string `json:"videoId"`
		} `json:"playlistItemData,omitempty"`

		NavigationEndpoint apiRespNavigationEndpoint `json:"navigationEndpoint,omitempty"`

		Thumbnail models.Thumbnail `json:"thumbnail,omitempty"`
	} `json:"musicResponsiveListItemRenderer"`
}

type ApiRespFlexColumns struct {
	MusicResponsiveListItemFlexColumnRenderer struct {
		Text struct {
			Runs models.NavigationAndWatchEndpointRuns
		} `json:"text"`
	} `json:"musicResponsiveListItemFlexColumnRenderer"`
}

type APIRespContinuation struct {
	NextContinuationData struct {
		Continuation string `json:"continuation"`
	} `json:"nextContinuationData"`
}

type apiRespNavigationEndpoint struct {
	WatchEndpoint struct {
		VideoID string `json:"videoId,omitempty"`
	} `json:"watchEndpoint,omitempty"`

	BrowseEndpoint struct {
		BrowseID string `json:"browseId,omitempty"`

		BrowseEndpointContextSupportedConfigs struct {
			BrowseEndpointContextMusicConfig struct {
				PageType string `json:"pageType,omitempty"`
			} `json:"browseEndpointContextMusicConfig,omitempty"`
		} `json:"browseEndpointContextSupportedConfigs,omitempty"`
	} `json:"browseEndpoint,omitempty"`
}
