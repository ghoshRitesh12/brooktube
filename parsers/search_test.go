package parsers_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
	"github.com/ghoshRitesh12/brooktube/models/search"
	"github.com/ghoshRitesh12/brooktube/parsers"
)

func TestGetSearchResults(t *testing.T) {
	parser := brooktube.NewParser()

	d, err := parser.GetSearchResults(
		"black clover",
		parsers.SearchParserParams{
			Category: search.COMMUNITY_PLAYLIST_SEARCH_KEY,
			// ContinuationToken: "EqkDEgxibGFjayBjbG92ZXIamANFZ1dLQVFJSUFVZ1VhZ29RQ1JBRkVBb1FBeEFFZ2dFTE5qZGZaVkppVWpoWWFsV0NBUXRIUkd0MGVYQnpWbVk1V1lJQkMwVkZOak5sTWxKS2JIQmpnZ0VMVDFweWFEWlZORU41UTNlQ0FRdEZOMVZ1ZFc5dk1IWklRWUlCQzFKZmIzcEllbEJEWWxZMGdnRUxaRk56VG1neFp5MHlORldDQVF0NFFXdFRYeTFHVkVoSmM0SUJDM1pPWjFoWFUxQlJkMncwZ2dFTGFIVmpTVlZqZG1KQ2JVbUNBUXQ1U1d4Q1lqRktkbXBpYTRJQkMwMTFPSHBIUzJOWlgwTnJnZ0VMWDJRdE9FZ3RPRXhxWTFXQ0FRdHpYMkoxT1ZFd01ISTBRWUlCQzFoblMwa3hObDloWTJoVmdnRUxTMGRhU0UxTWRHTlJNVldDQVF0VFZVNUljMVowUkhSTWQ0SUJDekJNTjB0V1RYSXlZVkZuZ2dFTExUbHpNRXd0Y1Vad1JEaUNBUXRNVlRST2FVcFRRa05JUVElM0QlM0QY8erQLg%3D%3D",
		},
	)

	for _, cmPlaylist := range d.Content.CommunityPlaylists {
		spew.Dump(cmPlaylist)
	}

	spew.Dump(d.ContinuationToken)

	if err != nil {
		t.Fatal(err)
	}
}
