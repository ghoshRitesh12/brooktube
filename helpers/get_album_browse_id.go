package helpers

import (
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/ghoshRitesh12/brooktube/utils"
)

const RAW_ALBUM_ID_PREFIX string = "OLAK5"

/*
takes in album id that start with "OLAK5" and returns browse id
that starts with "MPREb_"
*/
func GetAlbumBrowseId(albumId string) (string, error) {
	if !strings.HasPrefix(albumId, RAW_ALBUM_ID_PREFIX) {
		return "", utils.ErrInvalidAlbumId
	}

	url := utils.HOST + utils.PLAYLIST_PATH + albumId
	headers := map[string]string{
		"Accept":                   "*/*",
		"User-Agent":               utils.USER_AGENT_HEADER,
		"X-Goog-Visitor-Id":        utils.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": utils.CLIENT_VERSION,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	rawHTML, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	browseIdPattern := `"MPREb_.+?"`
	parseBrowseIdPattern := `MPREb_[^\\"]+`

	foundBrowseId := regexp.MustCompile(browseIdPattern).FindString(string(rawHTML))
	if foundBrowseId == "" {
		return "", utils.ErrCouldntGetAlbumBrowseId
	}

	parsedBrowserId := regexp.MustCompile(parseBrowseIdPattern).FindString(foundBrowseId)
	if parsedBrowserId == "" {
		return "", utils.ErrCouldntGetAlbumBrowseId
	}

	return parsedBrowserId, nil
}
