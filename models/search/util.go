package search

type SearchCategory string

const (
	SONG_SEARCH_KEY               SearchCategory = "song"
	VIDEO_SEARCH_KEY              SearchCategory = "video"
	ALBUM_SEARCH_KEY              SearchCategory = "album"
	ARTIST_SEARCH_KEY             SearchCategory = "artist"
	COMMUNITY_PLAYLIST_SEARCH_KEY SearchCategory = "community_playlist"
	FEATURED_PLAYLIST_SEARCH_KEY  SearchCategory = "featured_playlist"
)

var SEARCH_PARAMS_MAP = map[SearchCategory]string{
	SONG_SEARCH_KEY:               "EgWKAQIIAWoKEAkQBRAKEAMQBA%3D%3D",
	VIDEO_SEARCH_KEY:              "EgWKAQIQAWoKEAkQChAFEAMQBA%3D%3D",
	ALBUM_SEARCH_KEY:              "EgWKAQIYAWoKEAkQChAFEAMQBA%3D%3D",
	ARTIST_SEARCH_KEY:             "EgWKAQIgAWoKEAkQChAFEAMQBA%3D%3D",
	COMMUNITY_PLAYLIST_SEARCH_KEY: "EgeKAQQoAEABagoQAxAEEAoQCRAF",
	FEATURED_PLAYLIST_SEARCH_KEY:  "EgeKAQQoADgBagwQDhAKEAMQBRAJEAQ%3D",
}

var SEARCH_PARAMS_KEYS = map[SearchCategory]bool{
	SONG_SEARCH_KEY:               true,
	VIDEO_SEARCH_KEY:              true,
	ALBUM_SEARCH_KEY:              true,
	ARTIST_SEARCH_KEY:             true,
	COMMUNITY_PLAYLIST_SEARCH_KEY: true,
	FEATURED_PLAYLIST_SEARCH_KEY:  true,
}
