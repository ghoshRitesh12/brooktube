package parsers

import (
	"github.com/ghoshRitesh12/brooktube/internal/errors"
	"github.com/ghoshRitesh12/brooktube/internal/models/search"
	"github.com/ghoshRitesh12/brooktube/internal/requests"
)

// {query} is the song search query
func (_search) GetSongResults(query string) (*search.ScrapedSongResult, error) {
	result := &search.ScrapedSongResult{}
	return getSearchResults(query, result, &result.Contents)
}

// {query} is the song search query
//
// {continuationToken} is the token used for fetching paginated data
func (_search) GetNextSongResults(query, continuationToken string) (*search.ScrapedSongResult, error) {
	result := &search.ScrapedSongResult{}
	return getNextSearchResults(query, continuationToken, result, &result.Contents)
}

// {query} is the video search query
func (_search) GetVideoResults(query string) (*search.ScrapedVideoResult, error) {
	result := &search.ScrapedVideoResult{}
	return getSearchResults(query, result, &result.Contents)
}

// {query} is the video search query
//
// {continuationToken} is the token used for fetching paginated data
func (_search) GetNextVideoResults(query, continuationToken string) (*search.ScrapedVideoResult, error) {
	result := &search.ScrapedVideoResult{}
	return getNextSearchResults(query, continuationToken, result, &result.Contents)
}

// {query} is the artist search query
func (_search) GetArtistResults(query string) (*search.ScrapedArtistResult, error) {
	result := &search.ScrapedArtistResult{}
	return getSearchResults(query, result, &result.Contents)
}

// {query} is the artist search query
//
// {continuationToken} is the token used for fetching paginated data
func (_search) GetNextArtistResults(query, continuationToken string) (*search.ScrapedArtistResult, error) {
	result := &search.ScrapedArtistResult{}
	return getNextSearchResults(query, continuationToken, result, &result.Contents)
}

// {query} is the album search query
func (_search) GetAlbumResults(query string) (*search.ScrapedAlbumResult, error) {
	result := &search.ScrapedAlbumResult{}
	return getSearchResults(query, result, &result.Contents)
}

// {query} is the album search query
//
// {continuationToken} is the token used for fetching paginated data
func (_search) GetNextAlbumResults(query, continuationToken string) (*search.ScrapedAlbumResult, error) {
	result := &search.ScrapedAlbumResult{}
	return getNextSearchResults(query, continuationToken, result, &result.Contents)
}

// {query} is the community playlist search query
func (_search) GetCommunityPlaylistResults(query string) (*search.ScrapedCommunityPlaylistResult, error) {
	result := &search.ScrapedCommunityPlaylistResult{}
	return getSearchResults(query, result, &result.Contents)
}

// {query} is the community playlist search query
//
// {continuationToken} is the token used for fetching paginated data
func (_search) GetNextCommunityPlaylistResults(query, continuationToken string) (*search.ScrapedCommunityPlaylistResult, error) {
	result := &search.ScrapedCommunityPlaylistResult{}
	return getNextSearchResults(query, continuationToken, result, &result.Contents)
}

// {query} is the featured playlist search query
func (_search) GetFeaturedPlaylistResults(query string) (*search.ScrapedFeaturedPlaylistResult, error) {
	result := &search.ScrapedFeaturedPlaylistResult{}
	return getSearchResults(query, result, &result.Contents)
}

// {query} is the featured playlist search query
//
// {continuationToken} is the token used for fetching paginated data
func (_search) GetNextFeaturedPlaylistResults(query, continuationToken string) (*search.ScrapedFeaturedPlaylistResult, error) {
	result := &search.ScrapedFeaturedPlaylistResult{}
	return getNextSearchResults(query, continuationToken, result, &result.Contents)
}

func getSearchResults[R search.SearchResult](query string, result R, contents search.SearchContent) (R, error) {
	var res R = result
	category := search.SearchCategory("")

	switch contents.(type) {
	case *search.Songs:
		category = search.SONG_SEARCH_KEY
	case *search.Videos:
		category = search.VIDEO_SEARCH_KEY
	case *search.Albums:
		category = search.ALBUM_SEARCH_KEY
	case *search.Artists:
		category = search.ARTIST_SEARCH_KEY
	case *search.CommunityPlaylists:
		category = search.COMMUNITY_PLAYLIST_SEARCH_KEY
	case *search.FeaturedPlaylists:
		category = search.FEATURED_PLAYLIST_SEARCH_KEY
	default:
		return res, errors.ErrInvalidSearchCategory
	}

	if _, found := search.SEARCH_PARAMS_KEYS[category]; !found || category == "" {
		category = search.SONG_SEARCH_KEY
	}

	data, err := requests.FetchSearchResults(query, category)
	if err != nil {
		return res, err
	}

	tabs := data.Contents.TabbedSearchResultsRenderer.Tabs
	if len(tabs) == 0 {
		return res, errors.ErrSearchResultsNotFound
	}

	sections := tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(sections) == 0 {
		return res, errors.ErrSearchResultsNotFound
	}

	section := &(sections[0])
	res.SetBasicInfo(section)
	contents.ScrapeAndSet(&section.MusicShelfRenderer.Contents)

	return res, nil
}

func getNextSearchResults[R search.SearchResult](query, cToken string, result R, contents search.SearchContent) (R, error) {
	var res R = result

	if cToken == "" {
		return res, errors.ErrInvalidContinuationToken
	}

	data, err := requests.FetchNextSearchResults(query, cToken)
	if err != nil {
		return res, err
	}

	section := &data.ContinuationContents.MusicShelfContinuation
	res.SetContinuationToken(&data.ContinuationContents)
	contents.ScrapeAndSet(&section.Contents)

	return res, nil
}
