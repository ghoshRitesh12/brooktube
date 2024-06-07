package parsers_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
)

func TestGetPlaylist(t *testing.T) {
	const NARUTO_OST_PLAYLIST = "VLPLF7A13C44809B5893"
	// const BLACK_CLOVER_OST_PLAYLIST = "PLtwDCqqblBclwxAvPP0lTN56iWBjshBv7"
	brooktube := brooktube.New()

	playlist, err := brooktube.GetPlaylist(NARUTO_OST_PLAYLIST)
	if err != nil {
		t.Error(err)
	}

	spew.Dump(playlist, len(playlist.Tracks))
}
