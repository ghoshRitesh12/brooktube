package utils

import (
	"strings"

	"github.com/ghoshRitesh12/brooktube/models/search"
)

func ParseSearchContent(category search.SearchCategory, shelfContents []search.APIRespSectionContent) search.ResultContent {
	resultContent := search.ResultContent{}

	switch category {
	case search.SONG_SEARCH_KEY, search.VIDEO_SEARCH_KEY:
		resultContent.SongOrVideos = parseSongOrVideoContents(shelfContents, category)

	case search.ARTIST_SEARCH_KEY:
		for _, content := range shelfContents {
			artist := search.Artist{}

			for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
				if i == 0 {
					artist.Name = ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					})
					continue
				}

				artist.Subscribers = flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs[2].Text
			}

			browseEndpoint := content.MusicResponsiveListItemRenderer.
				NavigationEndpoint.BrowseEndpoint

			if browseEndpoint.
				BrowseEndpointContextSupportedConfigs.
				BrowseEndpointContextMusicConfig.
				PageType == MUSIC_PAGE_TYPE_ARTIST {
				artist.ChannelId = browseEndpoint.BrowseID
			}

			resultContent.Artists = append(resultContent.Artists, artist)
		}

	case search.ALBUM_SEARCH_KEY:
		otherInfoBuilder := strings.Builder{}

		for _, content := range shelfContents {
			album := search.Album{}

			for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
				textRuns := flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs

				if i == 0 {
					album.Name = ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: textRuns,
					})
					continue
				}

				otherInfoBuilder.WriteString(ParseYtTextField(ParseYtTextParams{
					FlexColumnRuns: textRuns,
				}))

				channelBrowseEndpoint := flexColumn.MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs[2].NavigationEndpoint.BrowseEndpoint

				if channelBrowseEndpoint.
					BrowseEndpointContextSupportedConfigs.
					BrowseEndpointContextMusicConfig.
					PageType == MUSIC_PAGE_TYPE_ARTIST {
					album.ArtistChannelId = channelBrowseEndpoint.BrowseID
				}

				album.YearOfRelease = textRuns[len(textRuns)-1].Text
			}

			album.OtherInfo = otherInfoBuilder.String()
			album.ArtistName = strings.Split(
				album.OtherInfo,
				OTHER_INFO_SEPARATOR,
			)[1]

			resultContent.Albums = append(resultContent.Albums, album)
			otherInfoBuilder.Reset()
		}

	case search.COMMUNITY_PLAYLIST_SEARCH_KEY:
		otherInfoBuilder := strings.Builder{}

		for _, content := range shelfContents {
			communityPlaylist := search.CommunityPlaylist{}

			for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
				textRuns := flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs

				if i == 0 {
					communityPlaylist.Name = ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: textRuns,
					})
					continue
				}

				otherInfoBuilder.WriteString(ParseYtTextField(ParseYtTextParams{
					FlexColumnRuns: textRuns,
				}))

				communityPlaylist.ArtistChannelId = textRuns[0].NavigationEndpoint.BrowseEndpoint.BrowseID
				communityPlaylist.Interactions = textRuns[len(textRuns)-1].Text
			}

			playlistBrowseEndpoint := content.MusicResponsiveListItemRenderer.
				NavigationEndpoint.BrowseEndpoint

			if playlistBrowseEndpoint.
				BrowseEndpointContextSupportedConfigs.
				BrowseEndpointContextMusicConfig.
				PageType == MUSIC_PAGE_TYPE_PLAYLIST {
				communityPlaylist.PlaylistId = playlistBrowseEndpoint.BrowseID
			}

			communityPlaylist.OtherInfo = otherInfoBuilder.String()
			resultContent.CommunityPlaylists = append(resultContent.CommunityPlaylists, communityPlaylist)
			otherInfoBuilder.Reset()
		}

	case search.FEATURED_PLAYLIST_SEARCH_KEY:
		otherInfoBuilder := strings.Builder{}

		for _, content := range shelfContents {
			featuredPlaylist := search.FeaturedPlaylist{}

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
			otherInfoBuilder.Reset()
		}
	default:
		return resultContent
	}

	return resultContent
}

func parseSongOrVideoContents(shelfContents []search.APIRespSectionContent, category search.SearchCategory) []search.SongOrVideo {
	songOrVideos := make([]search.SongOrVideo, 0, len(shelfContents))

	for _, content := range shelfContents {
		songOrVideo := search.SongOrVideo{
			SongOrVideoId: content.MusicResponsiveListItemRenderer.PlaylistItemData.VideoId,
		}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs

			if i == 0 {
				songOrVideo.Name = ParseYtTextField(ParseYtTextParams{
					FlexColumnRuns: textRuns,
				})
				continue
			}

			if i == 1 {
				songOrVideo.ArtistName = textRuns[0].Text

				channelBrowseEndpoint := flexColumn.MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs[0].NavigationEndpoint.BrowseEndpoint

				if (channelBrowseEndpoint.
					BrowseEndpointContextSupportedConfigs.
					BrowseEndpointContextMusicConfig.
					PageType == MUSIC_PAGE_TYPE_ARTIST) ||
					(channelBrowseEndpoint.
						BrowseEndpointContextSupportedConfigs.
						BrowseEndpointContextMusicConfig.
						PageType == MUSIC_PAGE_TYPE_USER_CHANNEL) {
					songOrVideo.ArtistChannelId = channelBrowseEndpoint.BrowseID
				}

				if category != search.VIDEO_SEARCH_KEY {
					songOrVideo.AlbumName = textRuns[2].Text
				}

				albumBrowseEndpoint := flexColumn.MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs[2].NavigationEndpoint.BrowseEndpoint

				if albumBrowseEndpoint.
					BrowseEndpointContextSupportedConfigs.
					BrowseEndpointContextMusicConfig.
					PageType == MUSIC_PAGE_TYPE_ALBUM {
					songOrVideo.AlbumId = albumBrowseEndpoint.BrowseID
				}

				songOrVideo.Duration = textRuns[len(textRuns)-1].Text
				songOrVideo.Interactions = textRuns[2].Text

				continue
			}

			songOrVideo.Interactions = textRuns[0].Text

		}

		songOrVideos = append(songOrVideos, songOrVideo)
	}

	return songOrVideos
}
