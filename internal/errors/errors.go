package errors

import "errors"

var (
	ErrInvalidAlbumId          = errors.New("album_parser: invalid album id")
	ErrInvalidAlbumBrowseId    = errors.New("album_parser: invalid album browse id")
	ErrAlbumContentsNotFound   = errors.New("album_parser: album contents not found")
	ErrCouldntGetAlbumBrowseId = errors.New("album_parser: could't get album browse id")
)

var (
	ErrInvalidPlaylistId        = errors.New("playlist_parser: invalid playlist id")
	ErrPlaylistContentsNotFound = errors.New("playlist_parser: playlist contents not found")
)

var ErrArtistContentNotFound = errors.New("artist_parser: artist content not found")

var ErrInvalidContinuationToken = errors.New("continued_content: invalid continuation token")
