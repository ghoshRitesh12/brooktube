package requests

import (
	"net/url"

	"github.com/ghoshRitesh12/brooktube/models"
	"github.com/ghoshRitesh12/brooktube/utils"
)

func FetchSearchSuggestions(query string) (models.SearchSuggestions, error) {
	method := "POST"
	reqURL, err := url.Parse(utils.HOST + utils.SEARCH_PATH)
	if err != nil {
		return models.SearchSuggestions{}, err
	}

	body := map[string]any{
		"input": query,
	}

	queryParams := reqURL.Query()
	queryParams.Set("prettyPrint", "false")

	reqURL.RawQuery = queryParams.Encode()

	headers := map[string]string{
		"X-Goog-Visitor-Id":        utils.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": utils.CLIENT_VERSION,
	}

	data, err := fetch[models.SearchSuggestions](method, reqURL.String(), body, headers)
	if err != nil {
		return data, err
	}

	return data, nil
}
