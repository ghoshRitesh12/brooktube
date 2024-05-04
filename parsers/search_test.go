package parsers_test

import (
	"fmt"
	"testing"

	"github.com/ghoshRitesh12/brooktube"
	"github.com/ghoshRitesh12/brooktube/models/search"
	"github.com/ghoshRitesh12/brooktube/parsers"
)

func TestGetSearchResults(t *testing.T) {
	parser := brooktube.NewParser()

	d, err := parser.GetSearchResults(
		"black clover",
		parsers.SearchParserParams{
			Category:          search.SONG_SEARCH_KEY,
			ContinuationToken: "ErEDEgxibGFjayBjbG92ZXIaoANFZ1dLQVFJSUFVZ29haEFRQ2hBREVBUVFDUkFRRUFVUUVSQVZnZ0VMVTFWT1NITldkRVIwVEhlQ0FRdFlaMHRKTVRaZllXTm9WWUlCQzJwcFQxVmFaVzFhYTFNNGdnRUxMVGx6TUV3dGNVWndSRGlDQVF0YVdVWlJTRFJyYldJNFZZSUJDMU5MWTBNM2RWVlNWRGRCZ2dFTGRuVnJYM1YzZVdVd1MxV0NBUXROZFdWMmJXWmFOVkZuUVlJQkMxQnVSWGRvVms1TVJXRmpnZ0VMZDJsclgwaHNXVzUwU1VHQ0FRdEhaVEowYlZOSGRuSlJhNElCQzNvM1h6aFFkV3BZYlhacmdnRUxPWFpYVG1GMVlWcEJaMmVDQVFzeE56QnpZMlZQVjFkWVk0SUJDM0ZVUzIxcmVIcFNNMGhyZ2dFTFYyZFFWMTlsYjFad05FMkNBUXN3ZGtaRVVHNXJiMWwxVVlJQkMxQTFXa0pKVTI4eVEyRkZnZ0VMT1RWNGJrbzRVVXhoTm5PQ0FRdFRjekJHTUhNelVDMXJNQSUzRCUzRBjx6tAu",
		},
	)

	fmt.Printf("SEARCH DATA:  %+v\n\n", d)
	if err != nil {
		t.Fatal(err)
	}
}
