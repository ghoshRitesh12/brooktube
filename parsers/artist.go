package parsers

import (
	"sync"

	"github.com/ghoshRitesh12/brooktube/models/artist"
	"github.com/ghoshRitesh12/brooktube/requests"
)

const SCRAPE_OPERATIONS int = 7

func (p *YTMusicAPI) GetArtist(artistChannelID string) (artist.ScrapedData, error) {
	wg := &sync.WaitGroup{}
	result := artist.ScrapedData{}

	data, err := requests.FetchArtist(artistChannelID)
	if err != nil {
		return result, err
	}

	sections := &data.Contents.SingleColumnBrowseResultsRenderer.
		Tabs[0].TabRenderer.Content.SectionListRenderer.Contents

	if len(*sections) == 0 {
		return result, nil
	}

	wg.Add(SCRAPE_OPERATIONS)

	go result.ScrapeAndSetBasicInfo(wg, &data.Header, sections)
	go result.SongsSection.ScrapeAndSet(wg, sections)

	for _, section := range *sections {
		sectionName := section.MusicCarouselShelfRenderer.
			Header.MusicCarouselShelfBasicHeaderRenderer.
			Title.Runs.GetText()

		switch sectionName {
		case "Albums":
			go result.AlbumsSection.ScrapeAndSet(wg, &section)
		case "Singles":
			go result.SinglesSection.ScrapeAndSet(wg, &section)
		case "Videos":
			go result.VideosSection.ScrapeAndSet(wg, &section)
		case "Featured on":
			go result.FeaturedOnSection.ScrapeAndSet(wg, &section)
		case "Fans might also like":
			go result.AlikeArtistsSection.ScrapeAndSet(wg, &section)
		default:
			continue
		}
	}

	wg.Wait()

	return result, nil
}
