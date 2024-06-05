package utils

import "errors"

var (
	ErrCouldntGetAlbumBrowseId = errors.New("could't get album browse id")
	ErrInvalidAlbumBrowseId    = errors.New("invalid album browse id")
	ErrAlbumContentsNotFound   = errors.New("album contents not found")
)
