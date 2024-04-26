package types

type SearchResults struct {
	Contents struct {
		TabbedSearchResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Content struct {
						SectionListRenderer struct {
							Contents []struct {
								// for top results
								MusicCardShelfRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`

									Buttons []struct {
										ButtonRenderer struct {
											Text struct {
												Runs []struct {
													Text string `json:"text"`
												} `json:"runs"`
											} `json:"text"`

											Accessibility struct {
												Label string `json:"label"`
											} `json:"accessibility"`

											Command struct {
												WatchPlaylistEndpoint struct {
													PlaylistID string `json:"playlistId"`
													Params     string `json:"params"`
												} `json:"watchPlaylistEndpoint"`
											} `json:"command"`
										} `json:"buttonRenderer"`
									} `json:"buttons"`
								} `json:"musicCardShelfRenderer,omitempty"`

								// for things like community playlists, songs, albums, etc.
								MusicShelfRenderer struct {
									Title struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"title"`

									Contents []struct {
										MusicResponsiveListItemRenderer struct {
											FlexColumns []struct {
												MusicResponsiveListItemFlexColumnRenderer struct {
													Text struct {
														Runs []struct {
															Text string `json:"text"`
														} `json:"runs"`
													} `json:"text"`
													DisplayPriority string `json:"displayPriority"`
												} `json:"musicResponsiveListItemFlexColumnRenderer"`
											} `json:"flexColumns"`

											Menu struct {
												MenuRenderer struct {
													Items []struct {
														MenuNavigationItemRenderer struct {
															Text struct {
																Runs []struct {
																	Text string `json:"text"`
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
										} `json:"musicResponsiveListItemRenderer"`
									} `json:"contents"`

									BottomText struct {
										Runs []struct {
											Text string `json:"text"`
										} `json:"runs"`
									} `json:"bottomText"`

									BottomEndpoint struct {
										SearchEndpoint struct {
											Query  string `json:"query"`
											Params string `json:"params"`
										} `json:"searchEndpoint"`
									} `json:"bottomEndpoint"`
								} `json:"musicShelfRenderer,omitempty"`
							} `json:"contents"`

							Header struct {
								ChipCloudRenderer struct {
									Chips []struct {
										ChipCloudChipRenderer struct {
											NavigationEndpoint struct {
												SearchEndpoint struct {
													Query  string `json:"query"`
													Params string `json:"params"`
												} `json:"searchEndpoint"`
											} `json:"navigationEndpoint"`

											UniqueID string `json:"uniqueId"`
										} `json:"chipCloudChipRenderer"`
									} `json:"chips"`
								} `json:"chipCloudRenderer"`
							} `json:"header"`
						} `json:"sectionListRenderer"`
					} `json:"content"`
				} `json:"tabRenderer"`
			} `json:"tabs"`
		} `json:"tabbedSearchResultsRenderer"`
	} `json:"contents"`
}
