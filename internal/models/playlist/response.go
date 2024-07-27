package playlist

import (
	"strings"
	"sync"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/models"
)

type ScrapedData struct {
	Title         string               `json:"title"`
	Subtitle      string               `json:"subtitle"`
	Description   string               `json:"description"`
	YearOfRelease string               `json:"yearOfRelease"`
	CoverArts     models.AppThumbnails `json:"coverArts"`

	TrackCount    string `json:"trackCount"`
	Interactions  string `json:"interactions"`
	TotalDuration string `json:"totalDuration"`

	ChannelName string `json:"channelName"`
	ChannelId   string `json:"channelId"`

	Tracks Tracks `json:"tracks"`

	ContinuationTokens []string `json:"continuationTokens"`
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

	playlist.ChannelName = headerRenderer.StraplineTextOne.Runs.GetText()
	_, browseId := headerRenderer.StraplineTextOne.Runs.GetNavData(0)
	playlist.ChannelId = browseId

	playlist.Interactions = headerRenderer.SecondSubtitle.Runs.GetText(0)
	playlist.TotalDuration = headerRenderer.SecondSubtitle.Runs.GetText(4)
	playlist.TrackCount = strings.Split(
		headerRenderer.SecondSubtitle.Runs.GetText(2),
		" ",
	)[0]

	playlist.CoverArts = background.GetAllThumbnails()
}

type (
	Track struct {
		VideoId    string               `json:"videoId"`
		Name       string               `json:"name"`
		Duration   string               `json:"duration"` // from fixedColumns
		IsExplicit bool                 `json:"isExplicit"`
		IsDisabled bool                 `json:"isDisabled"`
		Thumbnails models.AppThumbnails `json:"thumbnails"`

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
			VideoId:    contentData.PlaylistItemData.VideoID,
			IsExplicit: contentData.Badges.IsExplicit(),
			IsDisabled: contentData.MusicItemRendererDisplayPolicy.IsDisabled(),
			Thumbnails: contentData.Thumbnail.GetAllThumbnails(),
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
				if (pageType == constants.MUSIC_PAGE_TYPE_ARTIST) ||
					(pageType == constants.MUSIC_PAGE_TYPE_USER_CHANNEL) {
					playlistTrack.ChannelId = browseId
				}

				// case 2:
				// 	playlistTrack.Interactions = textRuns.GetText(0)
			}
		}

		*tracks = append(*tracks, playlistTrack)
	}
}
