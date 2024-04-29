package search

type RespResult struct {
	Contents struct {
		TabbedSearchResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Content struct {
						SectionListRenderer struct {
							Contents []RespResultSection `json:"contents"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"tabbedSearchResultsRenderer"`
	} `json:"contents"`
}

type RespResultSection struct {
	// for community playlists, songs, albums, videos, etc.
	MusicShelfRenderer struct {
		Title struct {
			Runs []struct {
				Text string
			} `json:"runs"`
		} `json:"title"`

		Contents      []RespSectionContent `json:"contents"`
		Continuations []RespContinuation   `json:"continuations,omitempty"`
	} `json:"musicShelfRenderer,omitempty"`
}

type RespSectionContent struct {
	MusicResponsiveListItemRenderer struct {
		FlexColumns []RespFlexColumns `json:"flexColumns"`

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

		NavigationEndpoint RespNavigationEndpoint `json:"navigationEndpoint,omitempty"`
	} `json:"musicResponsiveListItemRenderer"`
}

type RespFlexColumns struct {
	MusicResponsiveListItemFlexColumnRenderer struct {
		Text struct {
			Runs []RespFlexColumnRun
		} `json:"text"`
	} `json:"musicResponsiveListItemFlexColumnRenderer"`
}

// `json:"navigationEndpoint,omitempty"`
type RespContinuation struct {
	NextContinuationData struct {
		Continuation        string `json:"continuation"`
		ClickTrackingParams string `json:"clickTrackingParams"`
	} `json:"nextContinuationData"`
}

type RespFlexColumnRun struct {
	Text               string
	NavigationEndpoint RespNavigationEndpoint `json:"navigationEndpoint,omitempty"`
}

type RespNavigationEndpoint struct {
	WatchEndpoint struct {
		VideoID string `json:"videoId,omitempty"`

		WatchEndpointMusicSupportedConfigs struct {
			WatchEndpointMusicConfig struct {
				MusicVideoType string `json:"musicVideoType,omitempty"`
			} `json:"watchEndpointMusicConfig,omitempty"`
		} `json:"watchEndpointMusicSupportedConfigs,omitempty"`
		//
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
