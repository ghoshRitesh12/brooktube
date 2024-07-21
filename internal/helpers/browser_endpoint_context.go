package helpers

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

// func NewBrowserEndpointContext(musicPageType, browseId string) *BrowserEndpointContext {
func NewBrowserEndpointContext(musicPageType, browseId string) map[string]any {
	return map[string]any{
		"browseEndpointContextSupportedConfigs": map[string]any{
			"browseEndpointContextMusicConfig": map[string]string{
				"pageType": musicPageType,
			},
		},
		"browseId": browseId,
	}
}
