package parsers

import (
	"github.com/ghoshRitesh12/brooktube/helpers"
	"github.com/ghoshRitesh12/brooktube/models/search"
	"github.com/ghoshRitesh12/brooktube/requests"
)

type SearchParserParams struct {
	Category          search.SearchCategory // eg: community_playlists, song, video, album, artist, etc.
	ContinuationToken string                // token used for fetching paginated data
}

func (p *YTMusicAPI) GetSearchResults(query string, params SearchParserParams) (search.ScrapedData, error) {
	result := search.ScrapedData{}

	if _, ok := search.SEARCH_PARAMS_KEYS[params.Category]; !ok || params.Category == "" {
		params.Category = search.SONG_SEARCH_KEY
	}

	data, err := requests.FetchSearchResults(
		query,
		params.Category,
		params.ContinuationToken,
	)
	if err != nil {
		return result, err
	}

	if params.ContinuationToken != "" {
		section := data.ContinuationContents.MusicShelfContinuation

		result.ContinuationToken = helpers.PickContinuationToken(section.Continuations)
		result.Content = helpers.ParseSearchContent(params.Category, section.Contents)

		return result, nil
	}

	sections := data.Contents.TabbedSearchResultsRenderer.
		Tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(sections) == 0 || len(sections) > 1 {
		return result, nil
	}
	section := sections[0]

	result.ContinuationToken = helpers.PickContinuationToken(section.MusicShelfRenderer.Continuations)
	result.Title = helpers.ParseYtTextField(helpers.ParseYtTextParams{
		NormalRuns: section.MusicShelfRenderer.Title.Runs,
	})
	result.Content = helpers.ParseSearchContent(params.Category, section.MusicShelfRenderer.Contents)

	return result, nil
}
