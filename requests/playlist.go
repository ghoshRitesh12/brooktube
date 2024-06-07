package requests

import (
	"net/url"

	"github.com/ghoshRitesh12/brooktube/helpers"
	"github.com/ghoshRitesh12/brooktube/models/playlist"
	"github.com/ghoshRitesh12/brooktube/utils"
)

func FetchPlaylist(playlistId, continuationToken string) (*playlist.APIResp, error) {
	method := "POST"
	reqURL, err := url.Parse(utils.HOST + utils.BROWSE_PATH)
	if err != nil {
		return nil, err
	}

	body := helpers.NewBrowserEndpointContext(utils.MUSIC_PAGE_TYPE_PLAYLIST, playlistId)
	queryParams := reqURL.Query()

	if continuationToken != "" {
		body = map[string]any{}

		queryParams.Set("type", "next")
		queryParams.Set("ctoken", continuationToken)
		queryParams.Set("continuation", continuationToken)
	}

	queryParams.Set("prettyPrint", "false")
	reqURL.RawQuery = queryParams.Encode()

	headers := map[string]string{
		"X-Goog-Visitor-Id":        utils.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": utils.CLIENT_VERSION,
	}

	data, err := fetch[playlist.APIResp](method, reqURL.String(), body, headers)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
