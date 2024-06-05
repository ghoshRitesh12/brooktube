package helpers

import (
	"github.com/ghoshRitesh12/brooktube/models/search"
	"github.com/ghoshRitesh12/brooktube/utils"
)

func ParseArtistSongContents(shelfContents *[]search.APIRespSectionContent) []search.SongOrVideo {
	songOrVideos := make([]search.SongOrVideo, 0, len(*shelfContents))

	for _, content := range *shelfContents {
		songOrVideo := search.SongOrVideo{
			SongOrVideoId: content.
				MusicResponsiveListItemRenderer.
				PlaylistItemData.VideoId,
		}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.
				Text.Runs

			switch i {
			case 0:
				songOrVideo.Name = textRuns.GetText()

			case 1:
				songOrVideo.ArtistName = textRuns.GetText(0)

				pageType, browseId, _ := flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs.GetNavData(0)

				if pageType == utils.MUSIC_PAGE_TYPE_ARTIST {
					songOrVideo.ArtistChannelId = browseId
				}

			case 2:
				songOrVideo.Interactions = textRuns.GetText(0)

			case 3:
				innerTextRuns := flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs

				songOrVideo.AlbumName = innerTextRuns.GetText(0)

				pageType, browseId, _ := innerTextRuns.GetNavData(0)
				if pageType == utils.MUSIC_PAGE_TYPE_ALBUM {
					songOrVideo.AlbumId = browseId
				}
			}
		}

		songOrVideos = append(songOrVideos, songOrVideo)
	}

	return songOrVideos
}
