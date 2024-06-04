package brooktube

import (
	parser "github.com/ghoshRitesh12/brooktube/parsers"
)

const ASCIIArt string = `
_                     _    _         _
| |                   | |  | |       | |
| |__  _ __ ___   ___ | | _| |_ _   _| |__   ___
| '_ \| '__/ _ \ / _ \| |/ / __| | | | '_ \ / _ \
| |_) | | | (_) | (_) |   <| |_| |_| | |_) |  __/
|_.__/|_|  \___/ \___/|_|\_\\__|\__,_|_.__/ \___|
`

func New() *parser.YTMusicAPI {
	return parser.New()
}
