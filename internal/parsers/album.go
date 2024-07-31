package parsers

import (
	"strings"

	"github.com/ghoshRitesh12/brooktube/internal/errors"
	"github.com/ghoshRitesh12/brooktube/internal/models/album"
	"github.com/ghoshRitesh12/brooktube/internal/requests"
	"github.com/ghoshRitesh12/brooktube/internal/utils"
)

const ALBUM_BROWSE_ID_PREFIX string = "MPREb_"

func (p *Scraper) GetAlbum(albumId string) (*album.ScrapedData, error) {
	if albumId == "" {
		return nil, errors.ErrInvalidAlbumBrowseId
	}

	albumBrowseId := ""
	result := &album.ScrapedData{}

	if strings.HasPrefix(albumId, ALBUM_BROWSE_ID_PREFIX) {
		albumBrowseId = albumId
	} else {
		browseId, err := utils.GetAlbumBrowseId(albumId)
		if err != nil {
			return nil, err
		}
		albumBrowseId = browseId
	}

	if albumBrowseId == "" {
		return nil, errors.ErrInvalidAlbumBrowseId
	}

	data, err := requests.FetchAlbum(albumBrowseId)
	if err != nil {
		return nil, err
	}

	// spew.Dump(data)

	tabs := data.Contents.TwoColumnBrowseResultsRenderer.Tabs
	if len(tabs) == 0 {
		return nil, errors.ErrAlbumContentsNotFound
	}

	headerContents := tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(headerContents) < 1 {
		return nil, errors.ErrAlbumContentsNotFound
	}

	outerContents := data.Contents.TwoColumnBrowseResultsRenderer.
		SecondaryContents.SectionListRenderer.Contents
	if len(outerContents) < 1 {
		return nil, errors.ErrAlbumContentsNotFound
	}

	sections := &(outerContents[0].MusicShelfRenderer.Contents)
	if len(*sections) == 0 {
		return result, nil
	}

	result.ScrapeAndSetBasicInfo(&headerContents[0], &data.Background)
	result.Tracks.ScrapeAndSet(sections)

	return result, nil
}
