package parsers

import (
	models "github.com/ghoshRitesh12/brooktube/models"
	"github.com/ghoshRitesh12/brooktube/requests"
)

func (p *YtMusicParser) GetSearchSuggestions(query string) (models.SearchSuggestions, error) {
	return requests.FetchSearchSuggestions(query)
}
