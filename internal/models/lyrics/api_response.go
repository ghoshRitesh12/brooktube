package lyrics

import "github.com/ghoshRitesh12/brooktube/internal/models"

type NextAPIResp struct {
	Contents struct {
		SingleColumnMusicWatchNextResultsRenderer struct {
			TabbedRenderer struct {
				WatchNextTabbedResultsRenderer struct {
					Tabs NextAPIRespTabs `json:"tabs,omitempty"`
				} `json:"watchNextTabbedResultsRenderer,omitempty"`
			} `json:"tabbedRenderer,omitempty"`
		} `json:"singleColumnMusicWatchNextResultsRenderer,omitempty"`
	} `json:"contents,omitempty"`

	QueueContextParams string `json:"queueContextParams,omitempty"`
}

type NextAPIRespTabs [2]struct {
	TabRenderer models.TabRenderer `json:"tabRenderer,omitempty"`
}

type APIResp struct {
	Contents struct {
		SectionListRenderer struct {
			Contents [1]struct {
				MusicDescriptionShelfRenderer struct {
					Description struct {
						Runs models.BasicRuns `json:"runs,omitempty"`
					} `json:"description,omitempty"`

					MaxCollapsedLines int `json:"maxCollapsedLines,omitempty"`
					MaxExpandedLines  int `json:"maxExpandedLines,omitempty"`

					Footer struct {
						Runs models.BasicRuns `json:"runs,omitempty"`
					} `json:"footer,omitempty"`
				} `json:"musicDescriptionShelfRenderer,omitempty"`
			} `json:"contents,omitempty"`
		} `json:"sectionListRenderer,omitempty"`
	} `json:"contents,omitempty"`
}
