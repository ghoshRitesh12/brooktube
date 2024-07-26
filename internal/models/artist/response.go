package artist

import (
	"strings"
	"sync"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/models/search"
	"github.com/ghoshRitesh12/brooktube/internal/utils"
)

type ScrapedData struct {
	Name        string `json:"name"`
	About       string `json:"about"`
	Views       string `json:"views"`
	Subscribers string `json:"subscribersCount"`

	Songs        Songs        `json:"songs,omitempty"`               // section
	Videos       Videos       `json:"videos,omitempty"`              // section
	Albums       Albums       `json:"albums,omitempty"`              // section
	Singles      Singles      `json:"singles,omitempty"`             // section
	FeaturedOn   FeaturedOn   `json:"featuredOn,omitempty"`          // section
	AlikeArtists AlikeArtists `json:"alikeArtistsSection,omitempty"` // section
}

// scrapes and sets basic info of the artist
func (artist *ScrapedData) ScrapeAndSetBasicInfo(
	wg *sync.WaitGroup,
	header *apiRespHeader,
	sections *[]apiRespSectionContent,
) {
	defer wg.Done()

	artist.Name = header.MusicImmersiveHeaderRenderer.Title.Runs.GetText()
	artist.About = header.MusicImmersiveHeaderRenderer.Description.Runs.GetText()
	artist.Subscribers = header.MusicImmersiveHeaderRenderer.
		SubscriptionButton.SubscribeButtonRenderer.
		SubscriberCountText.Runs.GetText()

	artist.Views = strings.Split(
		(*sections)[len(*sections)-1].
			MusicDescriptionShelfRenderer.Subheader.
			Runs.GetText(),
		" ",
	)[0]
}

type Songs struct {
	Contents          search.Songs `json:"contents,omitempty"`
	SeeMorePlaylistId string       `json:"seeMorePlaylistId"`
}

// scrapes songs section data and sets it
func (songsSection *Songs) ScrapeAndSet(wg *sync.WaitGroup, sections *[]apiRespSectionContent) {
	defer wg.Done()
	section := (*sections)[0].MusicShelfRenderer

	if section.Title.Runs.GetText() != "Songs" {
		return
	}

	_, browseId := section.Title.Runs.GetNavData(0)
	songsSection.SeeMorePlaylistId = browseId
	songsSection.Contents = utils.ParseArtistSongContents(&(section.Contents))
}

type (
	Albums struct {
		Contents        []album `json:"contents,omitempty"`
		SeeMoreEndpoint struct {
			Params        string `json:"params"`
			DiscographyId string `json:"discographyId"` // somewhat like playlistId
		} `json:"seeMoreEndpoint"`
	}
	album struct {
		AlbumId  string `json:"albumId"` // browseEndpoint.browseId
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
	}
)

// scrapes albums section data and sets it
func (albumsSection *Albums) ScrapeAndSet(wg *sync.WaitGroup, section *apiRespSectionContent) {
	defer wg.Done()
	if section == nil {
		return
	}

	_, browseId, browseParams := section.MusicCarouselShelfRenderer.
		Header.MusicCarouselShelfBasicHeaderRenderer.
		Title.Runs.GetNavData(0)
	albumsSection.SeeMoreEndpoint.DiscographyId = browseId
	albumsSection.SeeMoreEndpoint.Params = browseParams

	albumsSection.Contents = make([]album, 0, len(section.MusicCarouselShelfRenderer.Contents))

	for _, content := range section.MusicCarouselShelfRenderer.Contents {
		albumsSection.Contents = append(
			albumsSection.Contents,
			album{
				AlbumId: content.MusicTwoRowItemRenderer.
					NavigationEndpoint.BrowseEndpoint.BrowseID,
				Title:    content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				Subtitle: content.MusicTwoRowItemRenderer.Subtitle.Runs.GetText(),
			},
		)
	}
}

type (
	Singles struct {
		Contents []single `json:"contents,omitempty"`

		SeeMoreEndpoint struct {
			Params        string `json:"params"`
			DiscographyId string `json:"discographyId"` // somewhat like playlistId
		} `json:"seeMoreEndpoint"`
	}
	single struct {
		AlbumId  string `json:"albumId"` // browseEndpoint.browseId
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
	}
)

// scrapes singles section data and sets it
func (singlesSection *Singles) ScrapeAndSet(wg *sync.WaitGroup, section *apiRespSectionContent) {
	defer wg.Done()
	if section == nil {
		return
	}

	_, browseId, browseParams := section.MusicCarouselShelfRenderer.
		Header.MusicCarouselShelfBasicHeaderRenderer.
		Title.Runs.GetNavData(0)
	singlesSection.SeeMoreEndpoint.DiscographyId = browseId
	singlesSection.SeeMoreEndpoint.Params = browseParams

	singlesSection.Contents = make([]single, 0, len(section.MusicCarouselShelfRenderer.Contents))

	for _, content := range section.MusicCarouselShelfRenderer.Contents {
		singlesSection.Contents = append(
			singlesSection.Contents,
			single{
				AlbumId: content.MusicTwoRowItemRenderer.
					NavigationEndpoint.BrowseEndpoint.BrowseID,
				Title:    content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				Subtitle: content.MusicTwoRowItemRenderer.Subtitle.Runs.GetText(),
			},
		)
	}
}

type (
	Videos struct {
		Contents          []video `json:"contents,omitempty"`
		SeeMorePlaylistId string  `json:"seeMorePlaylistId"`
	}
	video struct {
		VideoId  string `json:"videoId"` // watchEndpoint.videoId
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
	}
)

// scrapes videos section data and sets it
func (videosSection *Videos) ScrapeAndSet(wg *sync.WaitGroup, section *apiRespSectionContent) {
	defer wg.Done()
	if section == nil {
		return
	}

	_, browseId, _ := section.MusicCarouselShelfRenderer.
		Header.MusicCarouselShelfBasicHeaderRenderer.
		Title.Runs.GetNavData(0)
	videosSection.SeeMorePlaylistId = browseId

	videosSection.Contents = make([]video, 0, len(section.MusicCarouselShelfRenderer.Contents))

	for _, content := range section.MusicCarouselShelfRenderer.Contents {
		videosSection.Contents = append(
			videosSection.Contents,
			video{
				VideoId: content.MusicTwoRowItemRenderer.
					NavigationEndpoint.WatchEndpoint.VideoID,
				Title:    content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				Subtitle: content.MusicTwoRowItemRenderer.Subtitle.Runs.GetText(),
			},
		)
	}
}

type (
	FeaturedOn struct {
		Contents []featuredOn `json:"contents,omitempty"`
	}
	featuredOn struct {
		Title      string `json:"title"`
		PlaylistId string `json:"playlistId"`
	}
)

// scrapes featured on section data and sets it
func (featuredOnSection *FeaturedOn) ScrapeAndSet(wg *sync.WaitGroup, section *apiRespSectionContent) {
	defer wg.Done()
	if section == nil {
		return
	}

	featuredOnSection.Contents = make([]featuredOn, 0, len(section.MusicCarouselShelfRenderer.Contents))

	for _, content := range section.MusicCarouselShelfRenderer.Contents {
		browseEndpoint := content.MusicTwoRowItemRenderer.
			NavigationEndpoint.BrowseEndpoint

		featuredOnSection.Contents = append(
			featuredOnSection.Contents,
			featuredOn{
				Title:      content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				PlaylistId: browseEndpoint.BrowseID,
			},
		)
	}
}

type (
	AlikeArtists struct {
		Contents []alikeArtist `json:"contents,omitempty"`
	}
	alikeArtist struct {
		Name        string `json:"name"`
		ChannelId   string `json:"channelId"`
		Subscribers string `json:"subscribers"`
	}
)

// scrapes alike artists section data and sets it
func (alikeArtistSection *AlikeArtists) ScrapeAndSet(wg *sync.WaitGroup, section *apiRespSectionContent) {
	defer wg.Done()
	if section == nil {
		return
	}

	// to pre-allocate memory
	alikeArtistSection.Contents = make([]alikeArtist, 0, len(section.MusicCarouselShelfRenderer.Contents))

	for _, content := range section.MusicCarouselShelfRenderer.Contents {
		browseEndpoint := content.MusicTwoRowItemRenderer.
			NavigationEndpoint.BrowseEndpoint

		if browseEndpoint.
			BrowseEndpointContextSupportedConfigs.
			BrowseEndpointContextMusicConfig.
			PageType != constants.MUSIC_PAGE_TYPE_ARTIST {
			continue
		}

		alikeArtistSection.Contents = append(
			alikeArtistSection.Contents,
			alikeArtist{
				Name:      content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				ChannelId: browseEndpoint.BrowseID,
				Subscribers: strings.Split(
					content.MusicTwoRowItemRenderer.Subtitle.Runs.GetText(),
					" ",
				)[0],
			},
		)
	}
}
