package ytmusic_test

import (
	"fmt"
	"testing"

	ytmusic "github.com/ghoshRitesh12/yt_music"
	"github.com/ghoshRitesh12/yt_music/parsers"
	"github.com/ghoshRitesh12/yt_music/types/search"
)

func TestMain(t *testing.T) {
	parser := ytmusic.NewParser()

	d, err := parser.GetSearchResults(parsers.SearchParserParams{
		Query:    "black clover",
		Category: search.ARTIST_SEARCH_KEY,
	})

	fmt.Printf("Data: %+v", d)
	if err != nil {
		t.Fatal(err)
		return
	}
}
