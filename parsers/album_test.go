package parsers_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
)

func TestGetAlbum(t *testing.T) {
	const KAMIMAZE_ALBUM_ID = "OLAK5uy_kRVaDLvDemKrwYjkdUTryKHIyQa_RiiPo"
	parser := brooktube.NewParser()

	album, err := parser.GetAlbum(KAMIMAZE_ALBUM_ID)
	if err != nil {
		t.Error(err)
	}

	spew.Dump(album)
}
