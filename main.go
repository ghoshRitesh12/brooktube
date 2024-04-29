package ytmusic

import (
	"fmt"

	"github.com/ghoshRitesh12/yt_music/parsers"
)

func NewParser() *parsers.YtParser {
	fmt.Println("hello yt")

	return parsers.New()
}
