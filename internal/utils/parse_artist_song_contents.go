package utils

import (
	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/models/search"
)

func ParseArtistSongContents(shelfContents *[]search.APIRespSectionContent) []search.Song {
	songs := make([]search.Song, 0, len(*shelfContents))

	for _, content := range *shelfContents {
		song := search.Song{
			VideoId: content.MusicResponsiveListItemRenderer.
				PlaylistItemData.VideoId,
			Thumbnails: content.MusicResponsiveListItemRenderer.Thumbnail.GetAllThumbnails(),
		}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.
				Text.Runs

			switch i {
			case 0:
				song.Name = textRuns.GetText()

			case 1:
				song.ChannelName = textRuns.GetText(0)

				pageType, browseId, _ := flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs.GetNavData(0)

				if pageType == constants.MUSIC_PAGE_TYPE_ARTIST {
					song.ChannelId = browseId
				}

			case 2:
				song.Interactions = textRuns.GetText(0)

			case 3:
				innerTextRuns := flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs

				song.AlbumName = innerTextRuns.GetText(0)

				pageType, browseId, _ := innerTextRuns.GetNavData(0)
				if pageType == constants.MUSIC_PAGE_TYPE_ALBUM {
					song.AlbumId = browseId
				}
			}
		}

		songs = append(songs, song)
	}

	return songs
}
