package brooktube

import (
	"fmt"

	"github.com/ghoshRitesh12/brooktube/parsers"
)

func NewParser() *parsers.YtParser {
	fmt.Println("hello yt")

	return parsers.New()
}
