package brooktube

import (
	parser "github.com/ghoshRitesh12/brooktube/parsers"
)

func NewParser() *parser.YtMusicParser {
	return parser.New()
}
