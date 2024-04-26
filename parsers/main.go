package parsers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ghoshRitesh12/yt_music/utils"
)

type YtParser struct{}

func NewParser() *YtParser {
	return &YtParser{}
}

// utility function to make http requests
func fetch[T any](method string, url string, body map[string]any, headers map[string]string) (T, error) {
	var respBody T

	ytContext := utils.NewYtMusicContext()
	payload := map[string]any{
		"context": ytContext,
		// "user": map[string]any{
		// 	"lockedSafetyMode": false,
		// },
	}

	for key, val := range body {
		payload[key] = val
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return respBody, err
	}

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", utils.HOST)
	req.Header.Set("Referer", utils.HOST+"/")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", utils.USER_AGENT_HEADER)

	for key, val := range headers {
		req.Header.Set(key, val)
	}

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return respBody, err
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&respBody); err != nil {
		return respBody, err
	}

	return respBody, nil
}
