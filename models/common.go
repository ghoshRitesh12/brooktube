package models

import "strings"

type BasicRuns []struct {
	Text string
}

func (runs BasicRuns) GetTextField() string {
	var str strings.Builder

	if len(runs) > 0 {
		for _, val := range runs {
			str.WriteString(val.Text)
		}
		return str.String()
	}

	return str.String()
}

type NavigationEndpointRuns []struct {
	Text string

	NavigationEndpoint struct {
		BrowseEndpoint struct {
			BrowseID string `json:"browseId,omitempty"`

			BrowseEndpointContextSupportedConfigs struct {
				BrowseEndpointContextMusicConfig struct {
					PageType string `json:"pageType,omitempty"`
				} `json:"browseEndpointContextMusicConfig,omitempty"`
			} `json:"browseEndpointContextSupportedConfigs,omitempty"`
		} `json:"browseEndpoint,omitempty"`
	} `json:"navigationEndpoint,omitempty"`
}

func (runs *NavigationEndpointRuns) GetTextField() string {
	var str strings.Builder

	if len((*runs)) > 0 {
		for _, val := range *runs {
			str.WriteString(val.Text)
		}
		return str.String()
	}

	return str.String()
}

func (runs *NavigationEndpointRuns) GetNavData(index uint8) (pageType, browseId string) {
	if len((*runs)) == 0 || int(index) >= len((*runs)) {
		return
	}

	run := (*runs)[index]
	browseId = run.NavigationEndpoint.BrowseEndpoint.BrowseID
	pageType = run.NavigationEndpoint.BrowseEndpoint.BrowseEndpointContextSupportedConfigs.BrowseEndpointContextMusicConfig.PageType
	return
}

type NavigationEndpointParamsRuns []struct {
	Text string

	NavigationEndpoint struct {
		BrowseEndpoint struct {
			BrowseID string `json:"browseId,omitempty"`

			Params string `json:"params,omitempty"`

			BrowseEndpointContextSupportedConfigs struct {
				BrowseEndpointContextMusicConfig struct {
					PageType string `json:"pageType,omitempty"`
				} `json:"browseEndpointContextMusicConfig,omitempty"`
			} `json:"browseEndpointContextSupportedConfigs,omitempty"`
		} `json:"browseEndpoint,omitempty"`
	} `json:"navigationEndpoint,omitempty"`
}

func (runs *NavigationEndpointParamsRuns) GetTextField() string {
	var str strings.Builder

	if len((*runs)) > 0 {
		for _, val := range *runs {
			str.WriteString(val.Text)
		}
		return str.String()
	}

	return str.String()
}

func (runs *NavigationEndpointParamsRuns) GetNavData(index uint8) (pageType, browseId, browseParams string) {
	if len((*runs)) == 0 || int(index) >= len((*runs)) {
		return
	}

	run := (*runs)[index]
	browseId = run.NavigationEndpoint.BrowseEndpoint.BrowseID
	pageType = run.NavigationEndpoint.BrowseEndpoint.BrowseEndpointContextSupportedConfigs.BrowseEndpointContextMusicConfig.PageType
	browseParams = run.NavigationEndpoint.BrowseEndpoint.Params
	return
}

type NavigationAndWatchEndpointRuns struct {
	BrowseEndpoint struct {
		BrowseID string `json:"browseId,omitempty"`

		BrowseEndpointContextSupportedConfigs struct {
			BrowseEndpointContextMusicConfig struct {
				PageType string `json:"pageType,omitempty"`
			} `json:"browseEndpointContextMusicConfig,omitempty"`
		} `json:"browseEndpointContextSupportedConfigs,omitempty"`
	} `json:"browseEndpoint,omitempty"`

	WatchEndpoint struct {
		VideoID string `json:"videoId,omitempty"`

		WatchEndpointMusicSupportedConfigs struct {
			WatchEndpointMusicConfig struct {
				MusicVideoType string `json:"musicVideoType,omitempty"`
			} `json:"watchEndpointMusicConfig,omitempty"`
		} `json:"watchEndpointMusicSupportedConfigs,omitempty"`
	} `json:"watchEndpoint,omitempty"`
}
