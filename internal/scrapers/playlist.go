package scrapers

import (
	"strings"

	"github.com/ghoshRitesh12/brooktube/internal/errors"
	"github.com/ghoshRitesh12/brooktube/internal/models/playlist"
	"github.com/ghoshRitesh12/brooktube/internal/requests"
)

const PLAYLIST_ID_PREFIX string = "VL"

func (p *Scraper) GetPlaylist(playlistId string) (*playlist.ScrapedData, error) {
	if playlistId == "" {
		return nil, errors.ErrInvalidPlaylistId
	}

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
		return nil, errors.ErrPlaylistContentsNotFound
	}

	headerContents := tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(headerContents) < 1 {
		return nil, errors.ErrPlaylistContentsNotFound
	}

	outerContents := data.Contents.TwoColumnBrowseResultsRenderer.
		SecondaryContents.SectionListRenderer.Contents
	if len(outerContents) < 1 {
		return nil, errors.ErrPlaylistContentsNotFound
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

	result.ScrapeAndSetBasicInfo(&headerContents[0], &data.Background)
	result.Tracks.ScrapeAndSet(sections)

	return result, nil
}

// get more playlist tracks, if no more content is found, an empty array is returned
func (p *Scraper) GetMorePlaylistTracks(playlistId, continuationToken string) (*playlist.Tracks, error) {
	tracks := &playlist.Tracks{}

	if playlistId == "" {
		return nil, errors.ErrInvalidPlaylistId
	}
	if continuationToken == "" {
		return nil, errors.ErrInvalidContinuationToken
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

	tracks.ScrapeAndSet(contents)

	return tracks, nil
}
