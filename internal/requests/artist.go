package requests

import (
	"net/url"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/models/artist"
	"github.com/ghoshRitesh12/brooktube/internal/utils"
)

func FetchArtist(artistChannelID string) (*artist.APIResp, error) {
	method := "POST"
	reqURL, err := url.Parse(constants.HOST + constants.BROWSE_PATH)
	if err != nil {
		return nil, err
	}

	body := utils.NewBrowserEndpointContext(constants.MUSIC_PAGE_TYPE_ARTIST, artistChannelID)
	queryParams := reqURL.Query()

	queryParams.Set("prettyPrint", "false")
	reqURL.RawQuery = queryParams.Encode()

	headers := map[string]string{
		"X-Goog-Visitor-Id":        constants.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": constants.CLIENT_VERSION,
	}

	data, err := fetch[artist.APIResp](method, reqURL.String(), body, headers)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
