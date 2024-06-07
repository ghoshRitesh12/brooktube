package parsers

import (
	"strings"
	"sync"

	"github.com/ghoshRitesh12/brooktube/models/playlist"
	"github.com/ghoshRitesh12/brooktube/requests"
	"github.com/ghoshRitesh12/brooktube/utils"
)

const PLAYLIST_ID_PREFIX string = "VL"
const (
	PLAYLIST_SCRAPE_OPERATIONS        int = 2
	PLAYLIST_SCRAPE_OPERATIONS_CTOKEN int = 1
)

// `contiuationToken` is optional, if provided only 0th element will be used
func (p *YTMusicAPI) GetPlaylist(
	playlistId string,
	continuationToken ...string,
) (*playlist.ScrapedData, error) {
	if playlistId == "" {
		return nil, utils.ErrInvalidPlaylistId
	}

	wg := &sync.WaitGroup{}
	result := &playlist.ScrapedData{}
	cToken := ""

	if strings.HasPrefix(playlistId, PLAYLIST_ID_PREFIX) {
		playlistId = playlistId[len(PLAYLIST_ID_PREFIX):]
	} else {
		playlistId = PLAYLIST_ID_PREFIX + playlistId
	}

	if len(continuationToken) > 0 {
		cToken = continuationToken[0]
		if cToken == "" {
			return nil, utils.ErrInvalidContinuationToken
		}
	}

	data, err := requests.FetchPlaylist(playlistId, cToken)
	if err != nil {
		return nil, err
	}

	// if ctoken is not provided, scrape basic info and tracks
	if cToken == "" {
		outerContents := data.Contents.
			SingleColumnBrowseResultsRenderer.Tabs[0].
			TabRenderer.Content.SectionListRenderer.Contents

		if len(outerContents) < 1 {
			return nil, utils.ErrPlaylistContentsNotFound
		}

		sections := &(outerContents[0].MusicPlaylistShelfRenderer.Contents)

		if len(*sections) == 0 {
			return result, nil
		}

		wg.Add(PLAYLIST_SCRAPE_OPERATIONS)

		go result.ScrapeAndSetBasicInfo(wg, &data.Header)
		go result.Tracks.ScrapeAndSet(wg, sections)

		wg.Wait()

		return result, nil
	}

	contents := &data.ContinuationContents.
		MusicPlaylistShelfContinuation.Contents

	wg.Add(PLAYLIST_SCRAPE_OPERATIONS_CTOKEN)
	go result.Tracks.ScrapeAndSet(wg, contents)
	wg.Wait()

	return result, nil
}
