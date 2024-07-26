package utils

import "github.com/ghoshRitesh12/brooktube/internal/constants"

type YtMusicClient struct {
	HL            string `json:"hl"`
	Platform      string `json:"platform"`
	TimeZone      string `json:"timeZone"`
	UserAgent     string `json:"userAgent"`
	ClientName    string `json:"clientName"`
	VisitorData   string `json:"visitorData"`
	ClientVersion string `json:"clientVersion"`
}

type YtMusicContext struct {
	Client YtMusicClient `json:"client"`
}

func NewYtMusicContext() *YtMusicContext {
	return &YtMusicContext{
		Client: YtMusicClient{
			HL:            constants.HL,
			Platform:      constants.PLATFORM,
			TimeZone:      constants.TIME_ZONE,
			ClientName:    constants.CLIENT_NAME,
			ClientVersion: constants.CLIENT_VERSION,
			VisitorData:   constants.GOOG_VISITOR_ID,
			UserAgent:     constants.USER_AGENT_HEADER,
		},
	}
}
