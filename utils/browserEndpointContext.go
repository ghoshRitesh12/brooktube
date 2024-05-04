package utils

type BrowserEndpointContext struct {
	BrowseEndpointContextSupportedConfigs browseEndpointContextConfigs `json:"browseEndpointContextSupportedConfigs"`

	BrowseId string `json:"browseId"`
}

type browseEndpointContextConfigs struct {
	BrowseEndpointContextMusicConfig browseEndpointContextMusicConfig `json:"browseEndpointContextMusicConfig"`
}

type browseEndpointContextMusicConfig struct {
	PageType string `json:"pageType"`
}

func NewBrowserEndpointContext(pageType, browseId string) *BrowserEndpointContext {
	return &BrowserEndpointContext{
		BrowseEndpointContextSupportedConfigs: browseEndpointContextConfigs{
			BrowseEndpointContextMusicConfig: browseEndpointContextMusicConfig{
				PageType: "MUSIC_PAGE_TYPE_" + pageType,
			},
		},
		BrowseId: browseId,
	}
}
