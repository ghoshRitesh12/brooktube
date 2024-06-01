package utils

import (
	"github.com/ghoshRitesh12/brooktube/models/search"
)

func ParseArtistSongContents(shelfContents *[]search.APIRespSectionContent) []search.SongOrVideo {
	songOrVideos := make([]search.SongOrVideo, 0, len(*shelfContents))

	for _, content := range *shelfContents {
		songOrVideo := search.SongOrVideo{
			SongOrVideoId: content.MusicResponsiveListItemRenderer.PlaylistItemData.VideoId,
		}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs

			switch i {
			case 0:
				songOrVideo.Name = ParseYtTextField(ParseYtTextParams{
					FlexColumnRuns: textRuns,
				})

			case 1:
				songOrVideo.ArtistName = textRuns[0].Text

			case 2:
				songOrVideo.Interactions = textRuns[0].Text

			case 3:
				albumBrowseEndpoint := flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs[0].NavigationEndpoint.BrowseEndpoint

				if albumBrowseEndpoint.
					BrowseEndpointContextSupportedConfigs.
					BrowseEndpointContextMusicConfig.
					PageType == MUSIC_PAGE_TYPE_ALBUM {
					songOrVideo.AlbumName = flexColumn.
						MusicResponsiveListItemFlexColumnRenderer.
						Text.Runs[0].Text

					songOrVideo.AlbumId = albumBrowseEndpoint.BrowseID
				}
			}
		}

		songOrVideos = append(songOrVideos, songOrVideo)
	}

	return songOrVideos
}
