package errors

import "errors"

var (
	ErrInvalidAlbumId          = errors.New("album_parser: invalid album id")
	ErrInvalidAlbumBrowseId    = errors.New("album_parser: invalid album browse id")
	ErrAlbumContentsNotFound   = errors.New("album_parser: album contents not found")
	ErrCouldntGetAlbumBrowseId = errors.New("album_parser: could't get album browse id")
)

var ErrInvalidMediaId = errors.New("brooktube: invalid song or video id")

var (
	ErrInvalidPlaylistId        = errors.New("playlist_parser: invalid playlist id")
	ErrPlaylistContentsNotFound = errors.New("playlist_parser: playlist contents not found")
)

var ErrArtistContentNotFound = errors.New("artist_parser: artist content not found")

var (
	ErrLyricsContentNotFound = errors.New("lyrics_parser: lyrics content not found")
	ErrLyricsNotFound        = errors.New("lyrics_parser: lyrics unavailable for this song")
)

var ErrInvalidContinuationToken = errors.New("continued_content: invalid continuation token")

var (
	ErrSearchResultsNotFound = errors.New("search_parser: search results not found")
	ErrInvalidSearchCategory = errors.New("search_parser: invalid or unsupported search category")
)
