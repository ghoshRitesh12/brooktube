package parsers_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
	"github.com/ghoshRitesh12/brooktube/models/search"
	"github.com/ghoshRitesh12/brooktube/parsers"
)

func TestSubGetSearchSongs(t *testing.T) {
	t.Parallel()
	brooktube := brooktube.New()

	d, err := brooktube.GetSearchResults(
		"black clover",
		parsers.SearchParserParams{
			Category: search.SONG_SEARCH_KEY,
		},
	)

	for _, song := range d.Content.SongOrVideos {
		spew.Dump(song)
	}

	spew.Dump(d.ContinuationToken)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSubGetSearchVideos(t *testing.T) {
	t.Parallel()
	brooktube := brooktube.New()

	d, err := brooktube.GetSearchResults(
		"black clover",
		parsers.SearchParserParams{
			Category: search.VIDEO_SEARCH_KEY,
		},
	)

	for _, song := range d.Content.SongOrVideos {
		spew.Dump(song)
	}

	spew.Dump(d.ContinuationToken)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSubGetSearchArtists(t *testing.T) {
	t.Parallel()
	brooktube := brooktube.New()

	d, err := brooktube.GetSearchResults(
		"black clover",
		parsers.SearchParserParams{
			Category: search.ARTIST_SEARCH_KEY,
		},
	)

	for _, artist := range d.Content.Artists {
		spew.Dump(artist)
	}

	spew.Dump(d.ContinuationToken)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSubGetSearchAlbums(t *testing.T) {
	t.Parallel()
	brooktube := brooktube.New()

	d, err := brooktube.GetSearchResults(
		"black clover",
		parsers.SearchParserParams{
			Category: search.ALBUM_SEARCH_KEY,
		},
	)

	for _, album := range d.Content.Albums {
		spew.Dump(album)
	}

	spew.Dump(d.ContinuationToken)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSubGetSearchCommunityPlaylists(t *testing.T) {
	t.Parallel()
	brooktube := brooktube.New()

	d, err := brooktube.GetSearchResults(
		"black clover",
		parsers.SearchParserParams{
			Category: search.COMMUNITY_PLAYLIST_SEARCH_KEY,
		},
	)

	for _, cmPlaylist := range d.Content.CommunityPlaylists {
		spew.Dump(cmPlaylist)
	}

	spew.Dump(d.ContinuationToken)

	if err != nil {
		t.Fatal(err)
	}
}
