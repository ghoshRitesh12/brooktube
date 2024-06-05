package album

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

	SongCount     string `json:"songCount"`
	TotalDuration string `json:"totalDuration"`

	ArtistName      string `json:"artistName"`
	ArtistChannelId string `json:"artistChannelId"`

	Songs AlbumSongs `json:"songs"`
}

// scrapes and sets basic info of the album
func (album *ScrapedData) ScrapeAndSetBasicInfo(
	wg *sync.WaitGroup,
	header *apiRespHeader,
) {
	defer wg.Done()

	headerRenderer := header.MusicDetailHeaderRenderer

	album.Title = headerRenderer.Title.Runs.GetText()
	album.Subtitle = headerRenderer.Subtitle.Runs.GetText()
	album.Description = strings.Split(
		headerRenderer.Description.Runs.GetText(),
		"\n\n",
	)[0]

	album.YearOfRelease = headerRenderer.Subtitle.Runs.GetText(
		uint8(len(headerRenderer.Subtitle.Runs) - 1),
	)
	album.ArtistName = headerRenderer.Subtitle.Runs.GetText(2)

	_, browseId := headerRenderer.Subtitle.Runs.GetNavData(2)
	album.ArtistChannelId = browseId

	album.TotalDuration = headerRenderer.SecondSubtitle.Runs.GetText(2)
	album.SongCount = strings.Split(
		headerRenderer.SecondSubtitle.Runs.GetText(0),
		" ",
	)[0]
}

type (
	AlbumSong struct {
		SongOrVideoId   string `json:"songOrVideoId"`
		Name            string `json:"name"`
		ArtistName      string `json:"artistName"`
		ArtistChannelId string `json:"artistChannelId"`
		Duration        string `json:"duration"` // from fixedColumns
		Interactions    string `json:"interactions"`
		IsExplicit      bool   `json:"isExplicit"`
	}
	AlbumSongs []AlbumSong
)

func (albumSongs *AlbumSongs) ScrapeAndSet(
	wg *sync.WaitGroup,
	sections *[]apiRespSectionContent,
) {
	defer wg.Done()

	preAlbumSongs := make(AlbumSongs, 0, len(*sections))
	*albumSongs = preAlbumSongs

	for _, content := range *sections {
		contentData := content.MusicResponsiveListItemRenderer
		albumSong := AlbumSong{
			SongOrVideoId: contentData.PlaylistItemData.VideoID,
			IsExplicit:    contentData.Badges.IsExplicit(),
		}

		for i, fixedColumn := range contentData.FixedColumns {
			if i == 0 {
				albumSong.Duration = fixedColumn.
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
				albumSong.Name = textRuns.GetText()

			case 1:
				albumSong.ArtistName = textRuns.GetText()

				pageType, browseId, _ := textRuns.GetNavData(0)
				if (pageType == utils.MUSIC_PAGE_TYPE_ARTIST) ||
					(pageType == utils.MUSIC_PAGE_TYPE_USER_CHANNEL) {
					albumSong.ArtistChannelId = browseId
				}

			case 2:
				albumSong.Interactions = textRuns.GetText(0)
			}
		}

		*albumSongs = append(*albumSongs, albumSong)
	}
}
