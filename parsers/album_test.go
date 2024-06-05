package parsers_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
)

func TestGetAlbum(t *testing.T) {
	const KAMIMAZE_ALBUM_ID = "OLAK5uy_kRVaDLvDemKrwYjkdUTryKHIyQa_RiiPo"
	// const THE_EMINEM_SHOW_ALBUM_ID = "OLAK5uy_lqWe7SUa0zi9eDcuCSCi1eeiakfPi2skg"
	brooktube := brooktube.New()

	album, err := brooktube.GetAlbum(KAMIMAZE_ALBUM_ID)
	if err != nil {
		t.Error(err)
	}

	spew.Dump(album, len(album.Songs))
}
