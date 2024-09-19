package brooktube

import (
	"github.com/ghoshRitesh12/brooktube/internal/errors"
	"github.com/ghoshRitesh12/brooktube/internal/parsers"
)

var (
	ErrInvalidAlbumId           = errors.ErrInvalidAlbumId
	ErrInvalidAlbumBrowseId     = errors.ErrInvalidAlbumBrowseId
	ErrAlbumContentsNotFound    = errors.ErrAlbumContentsNotFound
	ErrCouldntGetAlbumBrowseId  = errors.ErrCouldntGetAlbumBrowseId
	ErrInvalidMediaId           = errors.ErrInvalidMediaId
	ErrInvalidPlaylistId        = errors.ErrInvalidPlaylistId
	ErrPlaylistContentsNotFound = errors.ErrPlaylistContentsNotFound
	ErrArtistContentNotFound    = errors.ErrArtistContentNotFound
	ErrLyricsContentNotFound    = errors.ErrLyricsContentNotFound
	ErrLyricsNotFound           = errors.ErrLyricsNotFound
	ErrInvalidContinuationToken = errors.ErrInvalidContinuationToken
	ErrSearchResultsNotFound    = errors.ErrSearchResultsNotFound
	ErrInvalidSearchCategory    = errors.ErrInvalidSearchCategory
)

type scraper struct {
	parsers.Scraper
}

// initialize a new [scraper] instance
func New() *scraper {
	return &scraper{}
}
