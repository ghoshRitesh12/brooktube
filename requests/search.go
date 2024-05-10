package requests

import (
	"net/url"

	"github.com/ghoshRitesh12/brooktube/models/search"
	"github.com/ghoshRitesh12/brooktube/utils"
)

func FetchSearchResults(query string, category search.SearchCategory, continuationToken string) (search.APIResp, error) {
	method := "POST"
	reqURL, err := url.Parse(utils.HOST + utils.SEARCH_PATH)
	if err != nil {
		return search.APIResp{}, err
	}

	body := map[string]any{}
	queryParams := reqURL.Query()

	if continuationToken != "" {
		queryParams.Set("type", "next")
		queryParams.Set("ctoken", continuationToken)
		queryParams.Set("continuation", continuationToken)
	} else {
		searchParamsId, ok := search.SEARCH_PARAMS_MAP[category]
		if !ok {
			searchParamsId = search.SEARCH_PARAMS_MAP[search.SONG_SEARCH_KEY]
		}

		body["query"] = query
		body["params"] = searchParamsId
	}

	queryParams.Set("prettyPrint", "false")
	reqURL.RawQuery = queryParams.Encode()

	headers := map[string]string{
		"X-Goog-Visitor-Id":        utils.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": utils.CLIENT_VERSION,
	}

	data, err := fetch[search.APIResp](method, reqURL.String(), body, headers)
	if err != nil {
		return data, err
	}

	return data, nil
}
