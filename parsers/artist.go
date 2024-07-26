package parsers

import (
	"fmt"
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghoshRitesh12/brooktube/internal/errors"
	"github.com/ghoshRitesh12/brooktube/internal/models/artist"
	"github.com/ghoshRitesh12/brooktube/internal/requests"
)

var ARTIST_SCRAPE_OPERATIONS int = 2

func (p *Scraper) GetArtist(artistChannelID string) (*artist.ScrapedData, error) {
	wg := &sync.WaitGroup{}
	result := &artist.ScrapedData{}

	data, err := requests.FetchArtist(artistChannelID)
	if err != nil {
		return nil, err
	}

	tabs := &data.Contents.SingleColumnBrowseResultsRenderer.Tabs
	if len(*tabs) == 0 {
		return nil, errors.ErrArtistContentNotFound
	}

	sections := (*tabs)[0].TabRenderer.Content.SectionListRenderer.Contents
	if len(sections) == 0 {
		return nil, errors.ErrArtistContentNotFound
	}

	for _, section := range sections {
		sectionName := artist.SectionName(
			section.MusicCarouselShelfRenderer.
				Header.MusicCarouselShelfBasicHeaderRenderer.
				Title.Runs.GetText(),
		)
		if _, ok := artist.VALID_ARTIST_SECTIONS[sectionName]; ok {
			fmt.Println(sectionName)
			ARTIST_SCRAPE_OPERATIONS += 1
		}
	}

	wg.Add(ARTIST_SCRAPE_OPERATIONS)

	spew.Dump(ARTIST_SCRAPE_OPERATIONS)

	go result.ScrapeAndSetBasicInfo(wg, &data.Header, &sections)
	go result.Songs.ScrapeAndSet(wg, &sections)

	for _, section := range sections {
		sectionName := artist.SectionName(
			section.MusicCarouselShelfRenderer.
				Header.MusicCarouselShelfBasicHeaderRenderer.
				Title.Runs.GetText(),
		)

		switch sectionName {
		case artist.SECTION_ALBUMS:
			go result.Albums.ScrapeAndSet(wg, &section)
		case artist.SECTION_SINGLES:
			go result.Singles.ScrapeAndSet(wg, &section)
		case artist.SECTION_VIDEOS:
			go result.Videos.ScrapeAndSet(wg, &section)
		case artist.SECTION_FEATURED_ON:
			go result.FeaturedOn.ScrapeAndSet(wg, &section)
		case artist.SECTION_ALIKE_ARTISTS:
			go result.AlikeArtists.ScrapeAndSet(wg, &section)
		default:
			continue
		}
	}

	wg.Wait()

	return result, nil
}
