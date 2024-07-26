package parsers

import (
	"sync"

	"github.com/ghoshRitesh12/brooktube/internal/errors"
	"github.com/ghoshRitesh12/brooktube/internal/models/artist"
	"github.com/ghoshRitesh12/brooktube/internal/requests"
)

var ARTIST_SCRAPE_OPERATIONS int = 1

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
		if len(section.MusicCarouselShelfRenderer.Contents) > 0 || len(section.MusicShelfRenderer.Contents) > 0 {
			ARTIST_SCRAPE_OPERATIONS += 1
		}
	}

	wg.Add(ARTIST_SCRAPE_OPERATIONS)

	go result.ScrapeAndSetBasicInfo(wg, &data.Header, &sections)

	for _, section := range sections {
		sectionName := artist.SectionName("")

		if len(section.MusicCarouselShelfRenderer.Contents) > 0 {
			sectionName = artist.SectionName(
				section.MusicCarouselShelfRenderer.
					Header.MusicCarouselShelfBasicHeaderRenderer.
					Title.Runs.GetText(),
			)
		} else if len(section.MusicShelfRenderer.Contents) > 0 {
			sectionName = artist.SectionName(
				section.MusicShelfRenderer.Title.Runs.GetText(),
			)
		}

		switch sectionName {
		case artist.SECTION_SONGS:
			go result.Songs.ScrapeAndSet(wg, &section.MusicShelfRenderer)
		case artist.SECTION_ALBUMS:
			go result.Albums.ScrapeAndSet(wg, &section.MusicCarouselShelfRenderer)
		case artist.SECTION_SINGLES:
			go result.Singles.ScrapeAndSet(wg, &section.MusicCarouselShelfRenderer)
		case artist.SECTION_VIDEOS:
			go result.Videos.ScrapeAndSet(wg, &section.MusicCarouselShelfRenderer)
		case artist.SECTION_FEATURED_ON:
			go result.FeaturedOns.ScrapeAndSet(wg, &section.MusicCarouselShelfRenderer)
		case artist.SECTION_ALIKE_ARTISTS:
			go result.AlikeArtists.ScrapeAndSet(wg, &section.MusicCarouselShelfRenderer)
		case artist.SECTION_PLAYLISTS:
			go result.Playlists.ScrapeAndSet(wg, &section.MusicCarouselShelfRenderer)
		default:
			continue
		}
	}

	wg.Wait()

	return result, nil
}
