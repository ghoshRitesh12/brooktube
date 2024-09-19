package requests

import (
	"net/url"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/errors"
	"github.com/ghoshRitesh12/brooktube/internal/models/search"
)

func FetchSearchResults(query string, category search.SearchCategory) (*search.APIResp, error) {
	searchParamsId, found := search.SEARCH_PARAMS_MAP[category]
	if !found {
		return nil, errors.ErrInvalidSearchCategory
	}

	method := "POST"
	reqURL, err := url.Parse(constants.HOST + constants.SEARCH_PATH)
	if err != nil {
		return nil, err
	}

	reqBody := map[string]any{
		"query":  query,
		"params": searchParamsId,
	}

	queryParams := reqURL.Query()
	queryParams.Set("prettyPrint", "false")
	reqURL.RawQuery = queryParams.Encode()

	reqHeaders := map[string]string{
		"X-Goog-Visitor-Id":        constants.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": constants.CLIENT_VERSION,
	}

	data, err := fetch[search.APIResp](method, reqURL.String(), reqBody, reqHeaders)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func FetchNextSearchResults(query, continuationToken string) (*search.APIRespContinuation, error) {
	method := "POST"
	reqBody := map[string]any{}
	reqURL, err := url.Parse(constants.HOST + constants.SEARCH_PATH)
	if err != nil {
		return nil, err
	}

	queryParams := reqURL.Query()
	queryParams.Set("type", "next")
	queryParams.Set("ctoken", continuationToken)
	queryParams.Set("continuation", continuationToken)
	queryParams.Set("prettyPrint", "false")
	reqURL.RawQuery = queryParams.Encode()

	reqHeaders := map[string]string{
		"X-Goog-Visitor-Id":        constants.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": constants.CLIENT_VERSION,
	}

	data, err := fetch[search.APIRespContinuation](method, reqURL.String(), reqBody, reqHeaders)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
