package requests

import (
	"fmt"

	"github.com/ghoshRitesh12/yt_music/types/search"
	"github.com/ghoshRitesh12/yt_music/utils"
)

func FetchSearchResults(query string, category search.SearchCategory) (search.RespResult, error) {
	method := "POST"
	url := fmt.Sprintf("%s%s?prettyPrint=false", utils.HOST, utils.SEARCH_PATH)

	paramsId, ok := search.SEARCH_PARAMS_MAP[category]
	if !ok {
		paramsId = search.SEARCH_PARAMS_MAP[search.SONG_SEARCH_KEY]
	}

	body := map[string]any{
		"query":  query,
		"params": paramsId,
	}
	headers := map[string]string{
		"X-Goog-Visitor-Id":        utils.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": utils.CLIENT_VERSION,
	}

	data, err := fetch[search.RespResult](method, url, body, headers)
	if err != nil {
		return data, err
	}

	return data, nil
}
