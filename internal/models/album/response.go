package album

import (
	"strconv"
	"strings"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/models"
)

type ScrapedData struct {
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	CoverArt      string `json:"coverArt"`
	Description   string `json:"description"`
	YearOfRelease string `json:"yearOfRelease"`

	TrackCount    string `json:"trackCount"`
	TotalDuration string `json:"totalDuration"`
	IsExplicit    bool   `json:"isExplicit"`

	ChannelName string `json:"channelName"`
	ChannelId   string `json:"channelId"`

	Tracks Tracks `json:"tracks"`
}

// scrapes and sets basic info of the album
func (album *ScrapedData) ScrapeAndSetBasicInfo(header *apiRespHeader, background *models.Thumbnail) {
	headerRenderer := header.MusicResponsiveHeaderRenderer

	album.Title = headerRenderer.Title.Runs.GetText()
	album.Subtitle = headerRenderer.Subtitle.Runs.GetText()
	album.Description = headerRenderer.Description.
		MusicDescriptionShelfRenderer.Description.Runs.GetText()

	album.IsExplicit = headerRenderer.SubtitleBadge.IsExplicit()
	album.YearOfRelease = headerRenderer.Subtitle.Runs.GetText(
		uint8(len(headerRenderer.Subtitle.Runs) - 1),
	)

	album.ChannelName = headerRenderer.StraplineTextOne.Runs.GetText(0)
	_, browseId := headerRenderer.StraplineTextOne.Runs.GetNavData(0)
	album.ChannelId = browseId

	album.TotalDuration = headerRenderer.SecondSubtitle.Runs.GetText(2)
	album.TrackCount = strings.Split(
		headerRenderer.SecondSubtitle.Runs.GetText(0),
		" ",
	)[0]

	album.CoverArt = background.GetThumbnail(2)
}

type (
	Track struct {
		Index      int    `json:"index"`
		VideoId    string `json:"videoId"`
		Name       string `json:"name"`
		Duration   string `json:"duration"` // from fixedColumns
		IsExplicit bool   `json:"isExplicit"`
		// IsDisabled    bool   `json:"isDisabled"`
		Interactions string `json:"interactions"`

		ChannelName string `json:"channelName"`
		ChannelId   string `json:"channelId"`
	}
	Tracks []Track
)

func (tracks *Tracks) ScrapeAndSet(contents *[]apiRespSectionContent) {
	*tracks = make(Tracks, 0, len(*contents))

	for _, content := range *contents {
		contentData := content.MusicResponsiveListItemRenderer
		track := Track{
			VideoId:    contentData.PlaylistItemData.VideoID,
			IsExplicit: contentData.Badges.IsExplicit(),
		}

		idx, err := strconv.Atoi(contentData.Index.Runs.GetText(0))
		if err != nil {
			idx = -1
		}
		track.Index = idx

		for i, fixedColumn := range contentData.FixedColumns {
			if i == 0 {
				track.Duration = fixedColumn.
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
				track.Name = textRuns.GetText()

			case 1:
				track.ChannelName = textRuns.GetText()

				pageType, browseId, _ := textRuns.GetNavData(0)
				if (pageType == constants.MUSIC_PAGE_TYPE_ARTIST) ||
					(pageType == constants.MUSIC_PAGE_TYPE_USER_CHANNEL) {
					track.ChannelId = browseId
				}

			case 2:
				track.Interactions = textRuns.GetText(0)
			}
		}

		*tracks = append(*tracks, track)
	}
}
