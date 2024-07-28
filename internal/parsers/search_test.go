package parsers_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
	"github.com/ghoshRitesh12/brooktube/internal/models/search"
	"github.com/ghoshRitesh12/brooktube/internal/parsers"
	"github.com/stretchr/testify/assert"
)

// go test -run TestGetSearchResults ./parsers -v -count=1
func TestGetSearchResults(t *testing.T) {
	innertube := brooktube.New()

	srch := parsers.Search{}
	srch.GetAlbumResults("bruh that is the coold")

	testsTable := []struct {
		query    string
		category search.SearchCategory
	}{
		{"black clover", search.SONG_SEARCH_KEY},
		// {"black clover", search.VIDEO_SEARCH_KEY},
	}

	for _, test := range testsTable {
		t.Run(test.query, func(t *testing.T) {
			// results, err := innertube.GetSearchResults(test.query, test.category)
			result, err := innertube.Search.GetVideoResults(test.query)

			assert.NoError(t, err)
			assert.NotEmpty(t, result.Contents)
			spew.Dump(result)

			t.Run("continuation", func(t *testing.T) {
				t.Parallel()
				res, err := innertube.Search.GetNextVideoResults(test.query, result.Info.ContinuationToken)

				assert.NoError(t, err)
				assert.NotEmpty(t, res.Contents)
				spew.Dump(res)
			})
		})
	}
}

// func TestSubGetSearchVideos(t *testing.T) {
// 	t.Parallel()
// 	brooktube := brooktube.New()

// 	d, err := brooktube.GetSearchResults(
// 		"black clover",
// 		parsers.SearchParserParams{
// 			Category: search.VIDEO_SEARCH_KEY,
// 		},
// 	)

// 	for _, video := range d.Content.Videos {
// 		spew.Dump(video)
// 	}

// 	spew.Dump(d.ContinuationToken)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestSubGetSearchArtists(t *testing.T) {
// 	t.Parallel()
// 	brooktube := brooktube.New()

// 	d, err := brooktube.GetSearchResults(
// 		"black clover",
// 		parsers.SearchParserParams{
// 			Category: search.ARTIST_SEARCH_KEY,
// 		},
// 	)

// 	for _, artist := range d.Content.Artists {
// 		spew.Dump(artist)
// 	}

// 	spew.Dump(d.ContinuationToken)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestSubGetSearchAlbums(t *testing.T) {
// 	t.Parallel()
// 	brooktube := brooktube.New()

// 	d, err := brooktube.GetSearchResults(
// 		"black clover",
// 		parsers.SearchParserParams{
// 			Category: search.ALBUM_SEARCH_KEY,
// 		},
// 	)

// 	for _, album := range d.Content.Albums {
// 		spew.Dump(album)
// 	}

// 	spew.Dump(d.ContinuationToken)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestSubGetSearchCommunityPlaylists(t *testing.T) {
// 	t.Parallel()
// 	brooktube := brooktube.New()

// 	d, err := brooktube.GetSearchResults(
// 		"black clover",
// 		parsers.SearchParserParams{
// 			Category: search.COMMUNITY_PLAYLIST_SEARCH_KEY,
// 		},
// 	)

// 	for _, cmPlaylist := range d.Content.CommunityPlaylists {
// 		spew.Dump(cmPlaylist)
// 	}

// 	spew.Dump(d.ContinuationToken)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
