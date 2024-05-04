package utils

import (
	"strings"

	"github.com/ghoshRitesh12/brooktube/models/search"
)

func ParseSearchContent(category search.SearchCategory, shelfContents []search.RespSectionContent) search.ResultContent {
	resultContent := search.ResultContent{}

	switch category {
	case search.SONG_SEARCH_KEY, search.VIDEO_SEARCH_KEY:
		for _, content := range shelfContents {
			separator := " • "
			var otherInfoBuilder strings.Builder
			songOrVideo := search.SongOrVideo{
				VideoId: content.MusicResponsiveListItemRenderer.PlaylistItemData.VideoId,
			}

			for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
				if i == 0 {
					songOrVideo.Name = ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					})
					continue
				}

				otherInfoBuilder.WriteString(
					ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					}) + separator,
				)

				if i == 1 {
					browseEndpoint := flexColumn.MusicResponsiveListItemFlexColumnRenderer.
						Text.Runs[0].NavigationEndpoint.BrowseEndpoint

					if strings.Contains(
						browseEndpoint.BrowseEndpointContextSupportedConfigs.
							BrowseEndpointContextMusicConfig.PageType,
						"MUSIC_PAGE_TYPE_ARTIST",
					) {
						songOrVideo.ArtistChannelId = browseEndpoint.BrowseID
					}
				}
			}

			otherInfo, _ := strings.CutSuffix(otherInfoBuilder.String(), separator)
			songOrVideo.OtherInfo = otherInfo

			resultContent.SongOrVideos = append(resultContent.SongOrVideos, songOrVideo)
		}

	case search.ARTIST_SEARCH_KEY:
		for _, content := range shelfContents {
			artist := search.Artist{}
			var otherInfoBuilder strings.Builder

			for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
				if i == 0 {
					artist.Name = ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					})
				} else {
					otherInfoBuilder.WriteString(ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					}))
				}
			}

			artist.OtherInfo = otherInfoBuilder.String()

			browseEndpoint := content.MusicResponsiveListItemRenderer.
				NavigationEndpoint.BrowseEndpoint

			if strings.Contains(
				browseEndpoint.BrowseEndpointContextSupportedConfigs.
					BrowseEndpointContextMusicConfig.PageType,
				"MUSIC_PAGE_TYPE_ARTIST",
			) {
				artist.ArtistChannelId = browseEndpoint.BrowseID
			}

			resultContent.Artists = append(resultContent.Artists, artist)
		}

	case search.ALBUM_SEARCH_KEY:
		for _, content := range shelfContents {
			album := search.Album{}
			var otherInfoBuilder strings.Builder

			for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
				if i == 0 {
					album.Name = ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					})
				} else {
					otherInfoBuilder.WriteString(ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					}))
				}
			}

			album.OtherInfo = otherInfoBuilder.String()
			resultContent.Albums = append(resultContent.Albums, album)
		}

	case search.COMMUNITY_PLAYLIST_SEARCH_KEY:
		for _, content := range shelfContents {
			communityPlaylist := search.CommunityPlaylist{}
			var otherInfoBuilder strings.Builder

			for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
				if i == 0 {
					communityPlaylist.Name = ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					})
				} else {
					otherInfoBuilder.WriteString(ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					}))
				}
			}

			communityPlaylist.OtherInfo = otherInfoBuilder.String()
			resultContent.CommunityPlaylists = append(resultContent.CommunityPlaylists, communityPlaylist)
		}

	case search.FEATURED_PLAYLIST_SEARCH_KEY:
		for _, content := range shelfContents {
			featuredPlaylist := search.FeaturedPlaylist{}
			var otherInfoBuilder strings.Builder

			for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
				if i == 0 {
					featuredPlaylist.Name = ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					})
				} else {
					otherInfoBuilder.WriteString(ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					}))
				}
			}

			featuredPlaylist.OtherInfo = otherInfoBuilder.String()
			resultContent.FeaturedPlaylists = append(resultContent.FeaturedPlaylists, featuredPlaylist)
		}
	}

	return resultContent
}
