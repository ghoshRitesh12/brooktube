package requests

import (
	"net/url"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/errors"
	"github.com/ghoshRitesh12/brooktube/internal/models/lyrics"
)

const NEXT_ENDPOINT_PARAMS = "8gEAmgMDCNgE"
const NEXT_ENDPOINT_PLAYER_PARAMS = "igMDCNgE"

func fetchNextContent(videoId string) (*lyrics.NextAPIResp, error) {
	method := "POST"
	reqURL, err := url.Parse(constants.HOST + constants.NEXT_PATH + "?prettyPrint=false")
	if err != nil {
		return nil, err
	}

	reqHeaders := map[string]string{
		"X-Goog-Visitor-Id":        constants.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": constants.CLIENT_VERSION,
	}
	reqBody := map[string]any{
		"isAudioOnly":  true,
		"videoId":      videoId,
		"params":       NEXT_ENDPOINT_PARAMS,
		"playerParams": NEXT_ENDPOINT_PLAYER_PARAMS,
	}

	data, err := fetch[lyrics.NextAPIResp](method, reqURL.String(), reqBody, reqHeaders)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func FetchLyricsData(videoId string) (*lyrics.APIResp, error) {
	nextContent, err := fetchNextContent(videoId)
	if err != nil {
		return nil, err
	}
	tabs := nextContent.Contents.
		SingleColumnMusicWatchNextResultsRenderer.
		TabbedRenderer.WatchNextTabbedResultsRenderer.Tabs

	if len(tabs) == 0 {
		return nil, errors.ErrLyricsContentNotFound
	}

	browseId := ""

	if tabs[1].TabRenderer.IsTabUnselectable() {
		return nil, errors.ErrLyricsNotFound
	}

	pageType, brwseId := tabs[1].TabRenderer.GetTabNavData()
	if pageType == constants.MUSIC_PAGE_TYPE_TRACK_LYRICS {
		browseId = brwseId
	}
	if browseId == "" {
		return nil, errors.ErrLyricsContentNotFound
	}

	method := "POST"
	reqURL, err := url.Parse(constants.HOST + constants.BROWSE_PATH + "?prettyPrint=false")
	if err != nil {
		return nil, err
	}

	reqBody := map[string]any{
		"browseId": browseId,
	}
	reqHeaders := map[string]string{
		"X-Goog-Visitor-Id":        constants.GOOG_VISITOR_ID,
		"X-Youtube-Client-Version": constants.CLIENT_VERSION,
	}

	data, err := fetch[lyrics.APIResp](method, reqURL.String(), reqBody, reqHeaders)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
