package utils

import (
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/errors"
)

const RAW_ALBUM_ID_PREFIX string = "OLAK5"

/*
takes in album id that start with "OLAK5" and returns browse id
that starts with "MPREb_"
*/
func GetAlbumBrowseId(albumId string) (string, error) {
	if !strings.HasPrefix(albumId, RAW_ALBUM_ID_PREFIX) {
		return "", errors.ErrInvalidAlbumId
	}

	url := constants.HOST + constants.PLAYLIST_PATH + albumId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", constants.USER_AGENT_HEADER)
	req.Header.Set("X-Goog-Visitor-Id", constants.GOOG_VISITOR_ID)
	req.Header.Set("X-Youtube-Client-Version", constants.CLIENT_VERSION)

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
		return "", errors.ErrCouldntGetAlbumBrowseId
	}

	parsedBrowserId := regexp.MustCompile(parseBrowseIdPattern).FindString(foundBrowseId)
	if parsedBrowserId == "" {
		return "", errors.ErrCouldntGetAlbumBrowseId
	}

	return parsedBrowserId, nil
}
