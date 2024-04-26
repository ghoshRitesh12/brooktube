package types

// SSg stands for SearchSuggestions

// main
type SearchSuggestions struct {
	Contents []SSgContents `json:"contents"`
}

type SSgRun struct {
	Text string `json:"text"`
	Bold bool   `json:"bold"`
}

type Suggestion struct {
	Runs []SSgRun `json:"runs"`
}

type SSgSearchEndpoint struct {
	Query string `json:"query"`
}

type SSgNavigationEndpoint struct {
	ClickTrackingParams string            `json:"clickTrackingParams"`
	SearchEndpoint      SSgSearchEndpoint `json:"searchEndpoint"`
}

type SSgIcon struct {
	IconType string `json:"iconType"`
}

type SearchSuggestionRenderer struct {
	Suggestion         Suggestion            `json:"suggestion"`
	NavigationEndpoint SSgNavigationEndpoint `json:"navigationEndpoint"`
	TrackingParams     string                `json:"trackingParams"`
	Icon               SSgIcon               `json:"icon"`
}

type SSgContents struct {
	SearchSuggestionsSectionRenderer SearchSuggestionsSectionRenderer `json:"searchSuggestionsSectionRenderer"`
}

type SearchSuggestionsSectionRenderer struct {
	Contents []SubContent `json:"contents"`
}

type SubContent struct {
	SearchSuggestionRenderer SearchSuggestionRenderer `json:"searchSuggestionRenderer"`
}
