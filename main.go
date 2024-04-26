package ytmusic

import (
	"fmt"

	"github.com/ghoshRitesh12/yt_music/parsers"
)

func New() *parsers.YtParser {
	fmt.Println("hello yt")

	return parsers.NewParser()
}
