package utils

import (
	"strings"

	"github.com/ghoshRitesh12/brooktube/models/search"
)

func ParseSearchContent(category search.SearchCategory, shelfContents []search.RespSectionContent) search.ResultContent {
	resultContent := search.ResultContent{}

	switch category {
	case search.SONG_SEARCH_KEY, search.VIDEO_SEARCH_KEY:
		resultContent.SongOrVideos = parseSongOrVideoContents(shelfContents)

	case search.ARTIST_SEARCH_KEY:
		otherInfoBuilder := strings.Builder{}

		for _, content := range shelfContents {
			artist := search.Artist{}

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

			if browseEndpoint.
				BrowseEndpointContextSupportedConfigs.
				BrowseEndpointContextMusicConfig.
				PageType == "MUSIC_PAGE_TYPE_ARTIST" {
				artist.ArtistChannelId = browseEndpoint.BrowseID
			}

			resultContent.Artists = append(resultContent.Artists, artist)
			otherInfoBuilder.Reset()
		}

	case search.ALBUM_SEARCH_KEY:
		otherInfoBuilder := strings.Builder{}

		for _, content := range shelfContents {
			album := search.Album{}

			for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
				if i == 0 {
					album.Name = ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					})
					continue
				}

				otherInfoBuilder.WriteString(ParseYtTextField(ParseYtTextParams{
					FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
				}))

				channelBrowseEndpoint := flexColumn.MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs[2].NavigationEndpoint.BrowseEndpoint

				if channelBrowseEndpoint.
					BrowseEndpointContextSupportedConfigs.
					BrowseEndpointContextMusicConfig.
					PageType == "MUSIC_PAGE_TYPE_ARTIST" {
					album.ArtistChannelId = channelBrowseEndpoint.BrowseID
				}
			}

			album.OtherInfo = otherInfoBuilder.String()
			resultContent.Albums = append(resultContent.Albums, album)
			otherInfoBuilder.Reset()
		}

	case search.COMMUNITY_PLAYLIST_SEARCH_KEY:
		otherInfoBuilder := strings.Builder{}

		for _, content := range shelfContents {
			communityPlaylist := search.CommunityPlaylist{}

			for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
				if i == 0 {
					communityPlaylist.Name = ParseYtTextField(ParseYtTextParams{
						FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
					})
					continue
				}

				otherInfoBuilder.WriteString(ParseYtTextField(ParseYtTextParams{
					FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
				}))
			}

			playlistBrowseEndpoint := content.MusicResponsiveListItemRenderer.
				NavigationEndpoint.BrowseEndpoint

			if playlistBrowseEndpoint.
				BrowseEndpointContextSupportedConfigs.
				BrowseEndpointContextMusicConfig.
				PageType == "MUSIC_PAGE_TYPE_PLAYLIST" {
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

func parseSongOrVideoContents(shelfContents []search.RespSectionContent) []search.SongOrVideo {
	songOrVideos := make([]search.SongOrVideo, 0, len(shelfContents))

	for _, content := range shelfContents {
		songOrVideo := search.SongOrVideo{
			SongOrVideoId: content.MusicResponsiveListItemRenderer.PlaylistItemData.VideoId,
		}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			if i == 0 {
				songOrVideo.Name = ParseYtTextField(ParseYtTextParams{
					FlexColumnRuns: flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs,
				})
				continue
			}

			if i == 1 {
				songOrVideo.ArtistName = flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs[0].Text

				channelBrowseEndpoint := flexColumn.MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs[0].NavigationEndpoint.BrowseEndpoint

				if channelBrowseEndpoint.
					BrowseEndpointContextSupportedConfigs.
					BrowseEndpointContextMusicConfig.
					PageType == "MUSIC_PAGE_TYPE_ARTIST" {
					songOrVideo.ArtistChannelId = channelBrowseEndpoint.BrowseID
				}

				songOrVideo.AlbumName = flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs[2].Text

				albumBrowseEndpoint := flexColumn.MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs[2].NavigationEndpoint.BrowseEndpoint

				if albumBrowseEndpoint.
					BrowseEndpointContextSupportedConfigs.
					BrowseEndpointContextMusicConfig.
					PageType == "MUSIC_PAGE_TYPE_ALBUM" {
					songOrVideo.AlbumId = albumBrowseEndpoint.BrowseID
				}

				songOrVideo.Duration = flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs[4].Text

				continue
			}

			songOrVideo.PlaysCount = strings.Split(
				flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs[0].Text,
				" plays",
			)[0]

		}
		// OTHER_INFO_SEPARATOR

		songOrVideos = append(songOrVideos, songOrVideo)
	}

	return songOrVideos
}
