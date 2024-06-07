package utils

import "errors"

var (
	ErrCouldntGetAlbumBrowseId = errors.New("could't get album browse id")
	ErrInvalidAlbumBrowseId    = errors.New("invalid album browse id")
	ErrAlbumContentsNotFound   = errors.New("album contents not found")
	ErrInvalidAlbumId          = errors.New("invalid album id")
)

var (
	ErrPlaylistContentsNotFound = errors.New("playlist contents not found")
	ErrInvalidPlaylistId        = errors.New("invalid playlist id")
)

var ErrInvalidContinuationToken = errors.New("invalid continuation token")
