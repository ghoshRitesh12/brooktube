package parsers_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
	"github.com/stretchr/testify/assert"
)

// go test -run TestGetLyrics ./internal/parsers -v -count=1
func TestGetLyrics(t *testing.T) {
	testsTable := []struct {
		mediaName string
		mediaId   string
	}{
		{"borderline by tape impala", "rymYToIEL9o"},  // should pass
		{"yellow by yoh kamiyama", "waYhTInn7GU"},     // should pass
		{"renegade by jayz ft eminem", "3JTLkJVDDE4"}, // should fail
	}

	btube := brooktube.New()

	for _, test := range testsTable {
		t.Run(test.mediaName, func(t *testing.T) {
			lyricsData, err := btube.GetLyrics(test.mediaId)
			assert.NoError(t, err)
			assert.NotEmpty(t, lyricsData)

			spew.Dump(lyricsData)
		})
	}
}
