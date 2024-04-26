package ytmusic_test

import (
	"fmt"
	"testing"

	ytmusic "github.com/ghoshRitesh12/yt_music"
)

func TestMain(t *testing.T) {
	parser := ytmusic.New()

	// data, err := parser.GetSearchSuggestions("black clover")
	// if err != nil {
	// 	t.Fatal(err)
	// 	return
	// }

	// for _, c := range data.Contents {
	// 	fmt.Printf("%+v\n", c.SearchSuggestionsSectionRenderer.Contents)
	// 	fmt.Println()
	// 	for _, cc := range c.SearchSuggestionsSectionRenderer.Contents {
	// 		fmt.Println(cc.SearchSuggestionRenderer.NavigationEndpoint.SearchEndpoint.Query)
	// 	}
	// }

	result, err := parser.GetSearchResults("black+clover")

	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Printf("%+v\n", result.Contents)
}
