package utils

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
			HL:            HL,
			Platform:      PLATFORM,
			TimeZone:      TIME_ZONE,
			ClientName:    CLIENT_NAME,
			ClientVersion: CLIENT_VERSION,
			VisitorData:   GOOG_VISITOR_ID,
			UserAgent:     USER_AGENT_HEADER,
		},
	}
}
