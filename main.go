package brooktube

import "github.com/ghoshRitesh12/brooktube/internal/parsers"

type scraper struct {
	parsers.Scraper
}

func New() *scraper {
	return &scraper{}
}
