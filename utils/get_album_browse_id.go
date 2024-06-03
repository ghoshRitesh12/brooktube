package utils

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

/*
takes in album id that start with "OLAK5" and returns browse id
that starts with "MPREb_"
*/
func GetAlbumBrowseId(albumId string) (string, error) {
	url := HOST + PLAYLIST_PATH + albumId
	headers := map[string]string{
		"Accept":                   "*/*",
		"User-Agent":               USER_AGENT_HEADER,
		"X-Goog-Visitor-Id":        GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": CLIENT_VERSION,
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

	browseIdRegex := regexp.MustCompile(`"MPREb_.+?"`)
	match := browseIdRegex.FindStringSubmatch(string(rawHTML))
	if len(match) != 1 {
		return "", fmt.Errorf("could not get album browse id")
	}

	foundBrowseId := match[0]
	parsedBrowserId := regexp.MustCompile(`MPREb_[^\\"]+`).FindString(foundBrowseId)

	if parsedBrowserId == "" {
		return "", fmt.Errorf("could not get album browse id")
	}

	return parsedBrowserId, nil
}
