package parsers

import (
	"strings"
	"sync"

	"github.com/ghoshRitesh12/brooktube/helpers"
	"github.com/ghoshRitesh12/brooktube/models/album"
	"github.com/ghoshRitesh12/brooktube/requests"
	"github.com/ghoshRitesh12/brooktube/utils"
)

const ALBUM_SCRAPE_OPERATIONS int = 2
const ALBUM_BROWSE_ID_PREFIX string = "MPREb_"

func (p *YTMusicAPI) GetAlbum(albumId string) (*album.ScrapedData, error) {
	wg := &sync.WaitGroup{}
	result := &album.ScrapedData{}
	albumBrowseId := ""

	if strings.HasPrefix(albumId, ALBUM_BROWSE_ID_PREFIX) {
		albumBrowseId = albumId
	} else {
		browseId, err := helpers.GetAlbumBrowseId(albumId)
		if err != nil {
			return nil, err
		}
		albumBrowseId = browseId
	}

	if albumBrowseId == "" {
		return nil, utils.ErrInvalidAlbumBrowseId
	}

	data, err := requests.FetchAlbum(albumBrowseId)
	if err != nil {
		return nil, err
	}

	outerContents := data.Contents.
		SingleColumnBrowseResultsRenderer.Tabs[0].
		TabRenderer.Content.SectionListRenderer.Contents
	if len(outerContents) < 1 {
		return nil, utils.ErrAlbumContentsNotFound
	}

	sections := &(outerContents[0].MusicShelfRenderer.Contents)

	if len(*sections) == 0 {
		return result, nil
	}

	wg.Add(ALBUM_SCRAPE_OPERATIONS)

	go result.ScrapeAndSetBasicInfo(wg, &data.Header)
	go result.Songs.ScrapeAndSet(wg, sections)

	wg.Wait()

	return result, nil
}
