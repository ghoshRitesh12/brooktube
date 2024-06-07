package playlist

import (
	"strings"
	"sync"

	"github.com/ghoshRitesh12/brooktube/utils"
)

type ScrapedData struct {
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	Description   string `json:"description"`
	YearOfRelease string `json:"yearOfRelease"`

	TrackCount   string `json:"trackCount"`
	Interactions string `json:"interactions"`

	ArtistName      string `json:"artistName"`
	ArtistChannelId string `json:"artistChannelId"`

	Tracks PlaylistTracks `json:"tracks"`
}

// scrapes and sets basic info of the playlist
func (playlist *ScrapedData) ScrapeAndSetBasicInfo(
	wg *sync.WaitGroup,
	header *apiRespHeader,
) {
	defer wg.Done()

	headerRenderer := header.MusicDetailHeaderRenderer

	playlist.Title = headerRenderer.Title.Runs.GetText()
	playlist.Subtitle = headerRenderer.Subtitle.Runs.GetText()
	playlist.Description = headerRenderer.Description.Runs.GetText()
	playlist.YearOfRelease = headerRenderer.Subtitle.Runs.GetText(
		uint8(len(headerRenderer.Subtitle.Runs) - 1),
	)

	playlist.ArtistName = headerRenderer.Subtitle.Runs.GetText(2)
	_, browseId := headerRenderer.Subtitle.Runs.GetNavData(2)
	playlist.ArtistChannelId = browseId

	playlist.Interactions = headerRenderer.SecondSubtitle.Runs.GetText(0)
	playlist.TrackCount = strings.Split(
		headerRenderer.SecondSubtitle.Runs.GetText(2),
		" ",
	)[0]
}

type (
	PlaylistTrack struct {
		SongOrVideoId   string `json:"songOrVideoId"`
		Name            string `json:"name"`
		ArtistName      string `json:"artistName"`
		ArtistChannelId string `json:"artistChannelId"`
		Duration        string `json:"duration"` // from fixedColumns
		// Interactions    string `json:"interactions"`
		IsExplicit bool `json:"isExplicit"`
		IsDisabled bool `json:"isDisabled"`
	}
	PlaylistTracks []PlaylistTrack
)

func (playlistTracks *PlaylistTracks) ScrapeAndSet(
	wg *sync.WaitGroup,
	contents *[]apiRespSectionContent,
) {
	defer wg.Done()

	prePlaylistTracks := make(PlaylistTracks, 0, len(*contents))
	*playlistTracks = prePlaylistTracks

	for _, content := range *contents {
		contentData := content.MusicResponsiveListItemRenderer
		playlistTrack := PlaylistTrack{
			SongOrVideoId: contentData.PlaylistItemData.VideoID,
			IsExplicit:    contentData.Badges.IsExplicit(),
			IsDisabled:    contentData.MusicItemRendererDisplayPolicy.IsDisabled(),
		}

		for i, fixedColumn := range contentData.FixedColumns {
			if i == 0 {
				playlistTrack.Duration = fixedColumn.
					MusicResponsiveListItemFixedColumnRenderer.
					Text.Runs.GetText()
				continue
			}
		}

		for i, flexColumn := range contentData.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.
				Text.Runs

			switch i {
			case 0:
				playlistTrack.Name = textRuns.GetText()

			case 1:
				playlistTrack.ArtistName = textRuns.GetText()

				pageType, browseId, _ := textRuns.GetNavData(0)
				if (pageType == utils.MUSIC_PAGE_TYPE_ARTIST) ||
					(pageType == utils.MUSIC_PAGE_TYPE_USER_CHANNEL) {
					playlistTrack.ArtistChannelId = browseId
				}

				// case 2:
				// 	playlistTrack.Interactions = textRuns.GetText(0)
			}
		}

		*playlistTracks = append(*playlistTracks, playlistTrack)
	}
}
