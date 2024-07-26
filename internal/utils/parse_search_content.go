package utils

import (
	"github.com/ghoshRitesh12/brooktube/internal/models/search"
)

func ParseSearchContent(category search.SearchCategory, shelfContents []search.APIRespSectionContent) search.ResultContent {
	resultContent := search.ResultContent{}

	switch category {
	case search.SONG_SEARCH_KEY:
		resultContent.Songs.ScrapeAndSet(shelfContents)
	case search.VIDEO_SEARCH_KEY:
		resultContent.Videos.ScrapeAndSet(shelfContents)
	case search.ARTIST_SEARCH_KEY:
		resultContent.Artists.ScrapeAndSet(shelfContents)
	case search.ALBUM_SEARCH_KEY:
		resultContent.Albums.ScrapeAndSet(shelfContents)
	case search.COMMUNITY_PLAYLIST_SEARCH_KEY:
		resultContent.CommunityPlaylists.ScrapeAndSet(shelfContents)
	case search.FEATURED_PLAYLIST_SEARCH_KEY:
		resultContent.FeaturedPlaylists.ScrapeAndSet(shelfContents)
	default:
		return resultContent
	}

	return resultContent
}
