package parsers

import (
	"slices"

	"github.com/ghoshRitesh12/yt_music/requests"
	"github.com/ghoshRitesh12/yt_music/types/search"
	"github.com/ghoshRitesh12/yt_music/utils"
)

type SearchParserParams struct {
	Query    string
	Category search.SearchCategory
}

// Performs a POST to fetch search results
func (p *YtParser) GetSearchResults(params SearchParserParams) (search.Result, error) {
	result := search.Result{}

	if params.Category == "" || !slices.Contains(search.SEARCH_PARAMS_KEYS, params.Category) {
		params.Category = search.SONG_SEARCH_KEY
	}

	data, err := requests.FetchSearchResults(params.Query, params.Category)
	if err != nil {
		return result, err
	}

	sections := data.Contents.TabbedSearchResultsRenderer.
		Tabs[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(sections) == 0 || len(sections) > 1 {
		return result, nil
	}

	if len(sections[0].MusicShelfRenderer.Continuations) > 0 {
		result.ContinuationToken = sections[0].MusicShelfRenderer.Continuations[0].NextContinuationData.Continuation
	}
	result.Title = utils.ParseYtTextField(utils.ParseYtTextParams{
		NormalRuns: sections[0].MusicShelfRenderer.Title.Runs,
	})
	result.Content = utils.ParseSearchContent(params.Category, sections[0].MusicShelfRenderer.Contents)

	return result, nil
}
