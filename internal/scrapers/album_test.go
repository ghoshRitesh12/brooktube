package scrapers_test

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube"
	"github.com/stretchr/testify/assert"
)

// go test -run TestGetAlbum ./internal/scrapers -v -count=1
func TestGetAlbum(t *testing.T) {
	testsTable := []struct {
		albumName string
		albumId   string
	}{
		{"kamikaze", "OLAK5uy_kRVaDLvDemKrwYjkdUTryKHIyQa_RiiPo"},
		{"the eminem show", "OLAK5uy_kkypLq7TlpT3uYdH3MbuHDiF2J3u-BRjc"},
		{"the eminem show expanded version", "OLAK5uy_lqWe7SUa0zi9eDcuCSCi1eeiakfPi2skg"},
	}

	btube := brooktube.New()

	for _, test := range testsTable {
		t.Run(test.albumName, func(t *testing.T) {
			album, err := btube.GetAlbum(test.albumId)

			assert.NoError(t, err)
			assert.NotNil(t, album)

			spew.Dump(album, len(album.Tracks))
			fmt.Println()
		})
	}
}
