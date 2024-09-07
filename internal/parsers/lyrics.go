package parsers

import (
	"errors"

	btubeErrs "github.com/ghoshRitesh12/brooktube/internal/errors"
	"github.com/ghoshRitesh12/brooktube/internal/models/lyrics"
	"github.com/ghoshRitesh12/brooktube/internal/requests"
)

// {mediaId} refers to video or song id
func (p *Scraper) GetLyrics(mediaId string) (*lyrics.ScrapedData, error) {
	if mediaId == "" {
		return nil, btubeErrs.ErrInvalidMediaId
	}

	result := &lyrics.ScrapedData{}

	data, err := requests.FetchLyricsData(mediaId)
	if err != nil {
		if errors.Is(err, btubeErrs.ErrLyricsNotFound) {
			result.Lyrics = "Lyrics unavailable for this song"
			return result, nil
		}
		return nil, err
	}

	result.ScrapeAndSet(*data)
	return result, nil
}
