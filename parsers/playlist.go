package parsers

import (
	"strings"
	"sync"

	"github.com/ghoshRitesh12/brooktube/models/playlist"
	"github.com/ghoshRitesh12/brooktube/requests"
	"github.com/ghoshRitesh12/brooktube/utils"
)

const PLAYLIST_ID_PREFIX string = "VL"
const PLAYLIST_SCRAPE_OPERATIONS int = 2

func (p *Scraper) GetPlaylist(playlistId string) (*playlist.ScrapedData, error) {
	if playlistId == "" {
		return nil, utils.ErrInvalidPlaylistId
	}

	wg := &sync.WaitGroup{}
	result := &playlist.ScrapedData{}

	if !strings.HasPrefix(playlistId, PLAYLIST_ID_PREFIX) {
		playlistId = PLAYLIST_ID_PREFIX + playlistId
	}

	data, err := requests.FetchPlaylist(playlistId)
	if err != nil {
		return nil, err
	}

	tabs := data.Contents.TwoColumnBrowseResultsRenderer.Tabs
	if len(tabs) == 0 {
		return nil, utils.ErrPlaylistContentsNotFound
	}

	headerContents := tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(headerContents) < 1 {
		return nil, utils.ErrPlaylistContentsNotFound
	}

	outerContents := data.Contents.TwoColumnBrowseResultsRenderer.
		SecondaryContents.SectionListRenderer.Contents
	if len(outerContents) < 1 {
		return nil, utils.ErrPlaylistContentsNotFound
	}

	sections := &(outerContents[0].MusicPlaylistShelfRenderer.Contents)
	if len(*sections) == 0 {
		return result, nil
	}

	result.ContinuationTokens = append(
		result.ContinuationTokens,
		outerContents[0].
			MusicPlaylistShelfRenderer.
			Continuations.GetContinuationToken(),
		data.Contents.TwoColumnBrowseResultsRenderer.
			SecondaryContents.SectionListRenderer.
			Continuations.GetContinuationToken(),
	)

	wg.Add(PLAYLIST_SCRAPE_OPERATIONS)

	go result.ScrapeAndSetBasicInfo(wg, &headerContents[0], &data.Background)
	go result.Tracks.ScrapeAndSet(wg, sections)

	wg.Wait()

	return result, nil
}

// get more playlist tracks, if no more content is found, an empty array is returned
func (p *Scraper) GetMorePlaylistTracks(playlistId, continuationToken string) (*playlist.Tracks, error) {
	tracks := &playlist.Tracks{}

	if playlistId == "" {
		return nil, utils.ErrInvalidPlaylistId
	}
	if continuationToken == "" {
		return nil, utils.ErrInvalidContinuationToken
	}

	if !strings.HasPrefix(playlistId, PLAYLIST_ID_PREFIX) {
		playlistId = PLAYLIST_ID_PREFIX + playlistId
	}

	data, err := requests.FetchMorePlaylistTracks(playlistId, continuationToken)
	if err != nil {
		return nil, err
	}

	contents := &data.ContinuationContents.
		MusicPlaylistShelfContinuation.Contents

	if len(*contents) == 0 {
		return tracks, nil
	}

	tracks.ScrapeAndSet(nil, contents)

	return tracks, nil
}
