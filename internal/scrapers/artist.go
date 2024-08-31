package scrapers

import (
	"github.com/ghoshRitesh12/brooktube/internal/errors"
	"github.com/ghoshRitesh12/brooktube/internal/models/artist"
	"github.com/ghoshRitesh12/brooktube/internal/requests"
)

func (p *Scraper) GetArtist(artistChannelID string) (*artist.ScrapedData, error) {
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

	result.ScrapeAndSetBasicInfo(&data.Header, &sections)

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
			result.Songs.ScrapeAndSet(&section.MusicShelfRenderer)
		case artist.SECTION_ALBUMS:
			result.Albums.ScrapeAndSet(&section.MusicCarouselShelfRenderer)
		case artist.SECTION_SINGLES:
			result.Singles.ScrapeAndSet(&section.MusicCarouselShelfRenderer)
		case artist.SECTION_VIDEOS:
			result.Videos.ScrapeAndSet(&section.MusicCarouselShelfRenderer)
		case artist.SECTION_FEATURED_ON:
			result.FeaturedOns.ScrapeAndSet(&section.MusicCarouselShelfRenderer)
		case artist.SECTION_ALIKE_ARTISTS:
			result.AlikeArtists.ScrapeAndSet(&section.MusicCarouselShelfRenderer)
		case artist.SECTION_PLAYLISTS:
			result.Playlists.ScrapeAndSet(&section.MusicCarouselShelfRenderer)
		default:
			continue
		}
	}

	return result, nil
}
