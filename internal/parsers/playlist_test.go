package parsers_test

import (
	"testing"

	"github.com/ghoshRitesh12/brooktube"
	"github.com/stretchr/testify/assert"
)

// go test -run TestGetPlaylist ./internal/parsers -v -count=1
func TestGetPlaylist(t *testing.T) {
	testsTable := []struct {
		name       string
		playlistId string
	}{
		{"black clover playlist", "PLtwDCqqblBclwxAvPP0lTN56iWBjshBv7"},
		{"naruto playlist", "VLPLF7A13C44809B5893"},
	}

	btube := brooktube.New()

	for i, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			playlist, err := btube.GetPlaylist(test.playlistId)

			assert.NoError(t, err)
			assert.NotNil(t, playlist)

			t.Logf("\n%#v\n\n", playlist.CoverArts)

			t.Run("continuation content", func(t *testing.T) {
				if i == 0 {
					t.Skip("skipping, as black clover playlist has no continuation tokens\n")
				}
				if len(playlist.ContinuationTokens) == 0 {
					return
				}

				t.Logf("\nTracks count: %s\nTracks length: %d\n", playlist.TrackCount, len(playlist.Tracks))
				t.Logf("\nPlaylist: %s\ncTokens: %#v\n\n", playlist.Title, playlist.ContinuationTokens)

				tracks, err := btube.GetMorePlaylistTracks(test.playlistId, playlist.ContinuationTokens[0])
				assert.NoError(t, err)
				assert.NotEmpty(t, tracks)

				// spew.Dump(tracks)
			})
		})
	}
}
