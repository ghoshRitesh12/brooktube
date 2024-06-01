package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ghoshRitesh12/brooktube/utils"
)

func fetch[T any](method string, reqUrl string, reqBody map[string]any, reqHeaders map[string]string) (T, error) {
	var respBody T

	ytContext := utils.NewYtMusicContext()
	payload := map[string]any{
		"context": ytContext,
	}

	for key, val := range reqBody {
		switch val.(type) {
		case string:
			if val == "" {
				continue
			}
		}

		payload[key] = val
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return respBody, err
	}

	req, _ := http.NewRequest(method, reqUrl, bytes.NewBuffer(jsonPayload))

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", utils.HOST)
	req.Header.Set("Referer", utils.HOST+"/")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", utils.USER_AGENT_HEADER)

	queryParams := req.URL.Query()
	queryParams.Set("key", utils.GOOG_API_KEY)
	queryParams.Set("alt", "json")

	req.URL.RawQuery = queryParams.Encode()

	for key, val := range reqHeaders {
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
