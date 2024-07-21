package requests

import (
	"net/url"

	"github.com/ghoshRitesh12/brooktube/internal/helpers"
	"github.com/ghoshRitesh12/brooktube/internal/models/album"
	"github.com/ghoshRitesh12/brooktube/internal/utils"
)

func FetchAlbum(albumBrowseId string) (*album.APIResp, error) {
	method := "POST"
	reqURL, err := url.Parse(utils.HOST + utils.BROWSE_PATH)
	if err != nil {
		return nil, err
	}

	body := helpers.NewBrowserEndpointContext(utils.MUSIC_PAGE_TYPE_ALBUM, albumBrowseId)
	queryParams := reqURL.Query()

	queryParams.Set("prettyPrint", "false")
	reqURL.RawQuery = queryParams.Encode()

	headers := map[string]string{
		"X-Goog-Visitor-Id":        utils.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": utils.CLIENT_VERSION,
	}

	data, err := fetch[album.APIResp](method, reqURL.String(), body, headers)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
