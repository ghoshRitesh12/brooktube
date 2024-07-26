package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/utils"
)

var defaultHeaders = map[string]string{
	"Accept":       "*/*",
	"Origin":       constants.HOST,
	"Referer":      constants.HOST + "/",
	"Content-Type": "application/json",
	"User-Agent":   constants.USER_AGENT_HEADER,
}

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

	req, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return respBody, err
	}

	// setting default headers
	for key, value := range defaultHeaders {
		req.Header.Set(key, value)
	}

	queryParams := req.URL.Query()
	queryParams.Set("key", constants.GOOG_API_KEY)
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
