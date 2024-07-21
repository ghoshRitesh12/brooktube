package parsers_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
)

// go test -run TestGetArtist ./parsers -v -count=1
func TestGetArtist(t *testing.T) {
	testsTable := []struct {
		channelName string
		channelId   string
	}{
		{"eminem", "UCedvOgsKFzcK3hA5taf3KoQ"},
		{"naruto playlist guy", "UClX0RWY5bWVzGrWMTTD8iIQ"},
	}

	btube := brooktube.New()

	for _, test := range testsTable {
		t.Run(test.channelName, func(t *testing.T) {

			artist, err := btube.GetArtist(test.channelId)
			if err != nil {
				t.Error(err)
			}

			spew.Dump(artist)
		})
	}
}
