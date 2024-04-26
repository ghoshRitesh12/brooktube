package parsers

import (
	"fmt"

	"github.com/ghoshRitesh12/yt_music/types"
	"github.com/ghoshRitesh12/yt_music/utils"
)

// Performs a POST to fetch search suggestions
func (p *YtParser) GetSearchResults(query string) (types.SearchResults, error) {
	url := fmt.Sprintf("%s%s?prettyPrint=false", utils.HOST, utils.SEARCH_PATH)

	method := "POST"
	body := map[string]any{
		"query": query,
	}
	headers := map[string]string{
		"X-Goog-Visitor-Id":        utils.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": utils.CLIENT_VERSION,
	}

	data, err := fetch[types.SearchResults](method, url, body, headers)
	if err != nil {
		return data, err
	}

	return data, nil
}
