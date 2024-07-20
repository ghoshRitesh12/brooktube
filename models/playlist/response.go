package playlist

import (
	"strings"
	"sync"

	"github.com/ghoshRitesh12/brooktube/models"
	"github.com/ghoshRitesh12/brooktube/utils"
)

type ScrapedData struct {
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	CoverArt      string `json:"coverArt"`
	Description   string `json:"description"`
	YearOfRelease string `json:"yearOfRelease"`

	TrackCount     string `json:"trackCount"`
	Interactions   string `json:"interactions"`
	ListeningHours string `json:"listeningHours"`

	ArtistName      string `json:"artistName"`
	ArtistChannelId string `json:"artistChannelId"`

	Tracks Tracks `json:"tracks"`

	ContinuationTokens []string `json:"continuationToken"`
}

// scrapes and sets basic info of the playlist
func (playlist *ScrapedData) ScrapeAndSetBasicInfo(wg *sync.WaitGroup, header *apiRespHeader, background *models.Thumbnail) {
	defer wg.Done()

	headerRenderer := header.MusicResponsiveHeaderRenderer

	playlist.Title = headerRenderer.Title.Runs.GetText()
	playlist.Subtitle = headerRenderer.Subtitle.Runs.GetText()
	playlist.Description = headerRenderer.Description.
		MusicDescriptionShelfRenderer.Description.Runs.GetText()

	playlist.YearOfRelease = headerRenderer.Subtitle.Runs.GetText(
		uint8(len(headerRenderer.Subtitle.Runs) - 1),
	)

	playlist.ArtistName = headerRenderer.StraplineTextOne.Runs.GetText()
	_, browseId := headerRenderer.StraplineTextOne.Runs.GetNavData(0)
	playlist.ArtistChannelId = browseId

	playlist.Interactions = headerRenderer.SecondSubtitle.Runs.GetText(0)
	playlist.ListeningHours = headerRenderer.SecondSubtitle.Runs.GetText(4)
	playlist.TrackCount = strings.Split(
		headerRenderer.SecondSubtitle.Runs.GetText(2),
		" ",
	)[0]

	playlist.CoverArt = background.GetThumbnail(0)
}

type (
	Track struct {
		SongOrVideoId string `json:"songOrVideoId"`
		Name          string `json:"name"`
		Duration      string `json:"duration"` // from fixedColumns
		Thumbnail     string `json:"thumbnail"`
		IsExplicit    bool   `json:"isExplicit"`
		IsDisabled    bool   `json:"isDisabled"`

		ChannelName string `json:"channelName"`
		ChannelId   string `json:"channelId"`
	}
	Tracks []Track
)

func (tracks *Tracks) ScrapeAndSet(wg *sync.WaitGroup, contents *[]apiRespSectionContent) {
	if wg != nil {
		defer wg.Done()
	}

	*tracks = make(Tracks, 0, len(*contents))

	for _, content := range *contents {
		contentData := content.MusicResponsiveListItemRenderer
		playlistTrack := Track{
			SongOrVideoId: contentData.PlaylistItemData.VideoID,
			IsExplicit:    contentData.Badges.IsExplicit(),
			IsDisabled:    contentData.MusicItemRendererDisplayPolicy.IsDisabled(),
			Thumbnail:     contentData.Thumbnail.GetThumbnail(0),
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
				playlistTrack.ChannelName = textRuns.GetText()

				pageType, browseId, _ := textRuns.GetNavData(0)
				if (pageType == utils.MUSIC_PAGE_TYPE_ARTIST) ||
					(pageType == utils.MUSIC_PAGE_TYPE_USER_CHANNEL) {
					playlistTrack.ChannelId = browseId
				}

				// case 2:
				// 	playlistTrack.Interactions = textRuns.GetText(0)
			}
		}

		*tracks = append(*tracks, playlistTrack)
	}
}
