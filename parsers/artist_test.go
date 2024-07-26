package parsers_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
	"github.com/stretchr/testify/assert"
)

// go test -run TestGetArtist ./parsers -v -count=1
func TestGetArtist(t *testing.T) {
	testsTable := []struct {
		channelName string
		channelId   string
	}{
		{"eminem", "UCedvOgsKFzcK3hA5taf3KoQ"},
		{"naruto playlist guy", "UClX0RWY5bWVzGrWMTTD8iIQ"},
		{"black clover playlist guy", "UCUup0eHG2-iKm7kVKqq0HzQ"},
	}

	btube := brooktube.New()

	for _, test := range testsTable {
		t.Run(test.channelName, func(t *testing.T) {
			artist, err := btube.GetArtist(test.channelId)
			assert.NoError(t, err)
			assert.NotNil(t, artist.Info)

			spew.Dump(artist.Info)
		})
	}
}
