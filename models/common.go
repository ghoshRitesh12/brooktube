package models

import "strings"

type BasicRuns []struct {
	Text string
}

func (runs BasicRuns) GetText() string {
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

func (runs *NavigationEndpointRuns) GetText() string {
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

func (runs *NavigationEndpointParamsRuns) GetText() string {
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

type NavigationAndWatchEndpointRuns []struct {
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

		WatchEndpoint struct {
			VideoID string `json:"videoId,omitempty"`

			WatchEndpointMusicSupportedConfigs struct {
				WatchEndpointMusicConfig struct {
					MusicVideoType string `json:"musicVideoType,omitempty"`
				} `json:"watchEndpointMusicConfig,omitempty"`
			} `json:"watchEndpointMusicSupportedConfigs,omitempty"`
		} `json:"watchEndpoint,omitempty"`
	} `json:"navigationEndpoint,omitempty"`
}

// only the 0th element is considered as *index* if provided
func (runs *NavigationAndWatchEndpointRuns) GetText(index ...uint8) string {
	runsLength := len((*runs))

	if len(index) > 0 {
		idx := int(index[0])

		if idx >= runsLength {
			return ""
		}

		return (*runs)[idx].Text
	}

	var str strings.Builder

	if runsLength > 0 {
		for _, run := range *runs {
			str.WriteString(run.Text)
		}
		return str.String()
	}

	return str.String()
}

func (runs *NavigationAndWatchEndpointRuns) GetNavData(
	index uint8,
) (pageType, browseId, videoId string) {
	if len((*runs)) == 0 || int(index) >= len((*runs)) {
		return
	}

	run := (*runs)[index]
	pageType = run.NavigationEndpoint.BrowseEndpoint.BrowseEndpointContextSupportedConfigs.BrowseEndpointContextMusicConfig.PageType
	browseId = run.NavigationEndpoint.BrowseEndpoint.BrowseID
	videoId = run.NavigationEndpoint.WatchEndpoint.VideoID
	return
}

type Continuations []struct {
	NextContinuationData struct {
		Continuation string `json:"continuation"`
	} `json:"nextContinuationData"`
}

func (continuations *Continuations) GetContinuationToken() string {
	if len(*continuations) > 0 {
		return (*continuations)[0].NextContinuationData.Continuation
	}
	return ""
}
