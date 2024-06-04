package parsers_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
)

func TestGetArtist(t *testing.T) {
	const EMINEM_CHANNEL_ID = "UCedvOgsKFzcK3hA5taf3KoQ"
	brooktube := brooktube.New()

	artist, err := brooktube.GetArtist(EMINEM_CHANNEL_ID)
	if err != nil {
		t.Error(err)
	}

	spew.Dump(artist)
}
