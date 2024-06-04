package artist

import (
	"strings"
	"sync"

	"github.com/ghoshRitesh12/brooktube/helpers"
	"github.com/ghoshRitesh12/brooktube/models/search"
	"github.com/ghoshRitesh12/brooktube/utils"
)

type ScrapedData struct {
	Name        string `json:"name"`
	About       string `json:"about"`
	Views       string `json:"views"`
	Subscribers string `json:"subscribersCount"`

	SongsSection        ArtistSongsSection      `json:"songsSection,omitempty"`
	VideosSection       ArtistVideosSection     `json:"videosSection,omitempty"`
	AlbumsSection       ArtistAlbumsSection     `json:"albumsSection,omitempty"`
	SinglesSection      ArtistSinglesSection    `json:"singlesSection,omitempty"`
	FeaturedOnSection   ArtistFeaturedOnSection `json:"featuredOnSection,omitempty"`
	AlikeArtistsSection AlikeArtistsSection     `json:"alikeArtistsSection,omitempty"`
}

// scrapes and sets basic info of the artist
func (artist *ScrapedData) ScrapeAndSetBasicInfo(
	wg *sync.WaitGroup,
	header *apiRespHeader,
	sections *[]apiRespSectionContent,
) {
	defer wg.Done()

	artist.Name = header.MusicImmersiveHeaderRenderer.Title.Runs.GetTextField()
	artist.About = header.MusicImmersiveHeaderRenderer.Description.Runs.GetTextField()
	artist.Subscribers = header.MusicImmersiveHeaderRenderer.
		SubscriptionButton.SubscribeButtonRenderer.
		SubscriberCountText.Runs.GetTextField()

	artist.Views = strings.Split(
		(*sections)[len(*sections)-1].
			MusicDescriptionShelfRenderer.Subheader.
			Runs.GetTextField(),
		" ",
	)[0]
}

type ArtistSongsSection struct {
	Contents          []search.SongOrVideo `json:"contents,omitempty"`
	SeeMorePlaylistId string               `json:"seeMorePlaylistId"`
}

// scrapes songs section data and sets it
func (songsSection *ArtistSongsSection) ScrapeAndSet(
	wg *sync.WaitGroup,
	sections *[]apiRespSectionContent,
) {
	defer wg.Done()
	section := (*sections)[0].MusicShelfRenderer

	if section.Title.Runs.GetTextField() != "Songs" {
		return
	}

	_, browseId := section.Title.Runs.GetNavData(0)
	songsSection.SeeMorePlaylistId = browseId

	// spew.Dump("Before parsing song contents", len(section.Contents))
	songsSection.Contents = helpers.ParseArtistSongContents(&(section.Contents))
}

type (
	ArtistAlbumsSection struct {
		Contents        []artistAlbum `json:"contents,omitempty"`
		SeeMoreEndpoint struct {
			Params        string `json:"params"`
			DiscographyId string `json:"discographyId"` // somewhat like playlistId
		} `json:"seeMoreEndpoint"`
	}
	artistAlbum struct {
		AlbumId  string `json:"albumId"` // browseEndpoint.browseId
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
	}
)

// scrapes albums section data and sets it
func (albumsSection *ArtistAlbumsSection) ScrapeAndSet(
	wg *sync.WaitGroup,
	section *apiRespSectionContent,
) {
	defer wg.Done()
	if section == nil {
		return
	}

	_, browseId, browseParams := section.MusicCarouselShelfRenderer.
		Header.MusicCarouselShelfBasicHeaderRenderer.
		Title.Runs.GetNavData(0)
	albumsSection.SeeMoreEndpoint.DiscographyId = browseId
	albumsSection.SeeMoreEndpoint.Params = browseParams

	albumsSection.Contents = make([]artistAlbum, 0, len(section.MusicCarouselShelfRenderer.Contents))

	for _, content := range section.MusicCarouselShelfRenderer.Contents {
		albumsSection.Contents = append(
			albumsSection.Contents,
			artistAlbum{
				AlbumId: content.MusicTwoRowItemRenderer.
					NavigationEndpoint.BrowseEndpoint.BrowseID,
				Title:    content.MusicTwoRowItemRenderer.Title.Runs.GetTextField(),
				Subtitle: content.MusicTwoRowItemRenderer.Subtitle.Runs.GetTextField(),
			},
		)
	}
}

type (
	ArtistSinglesSection struct {
		Contents []artistSingle `json:"contents,omitempty"`

		SeeMoreEndpoint struct {
			Params        string `json:"params"`
			DiscographyId string `json:"discographyId"` // somewhat like playlistId
		} `json:"seeMoreEndpoint"`
	}
	artistSingle struct {
		AlbumId  string `json:"albumId"` // browseEndpoint.browseId
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
	}
)

// scrapes singles section data and sets it
func (singlesSection *ArtistSinglesSection) ScrapeAndSet(
	wg *sync.WaitGroup,
	section *apiRespSectionContent,
) {
	defer wg.Done()
	if section == nil {
		return
	}

	_, browseId, browseParams := section.MusicCarouselShelfRenderer.
		Header.MusicCarouselShelfBasicHeaderRenderer.
		Title.Runs.GetNavData(0)
	singlesSection.SeeMoreEndpoint.DiscographyId = browseId
	singlesSection.SeeMoreEndpoint.Params = browseParams

	singlesSection.Contents = make([]artistSingle, 0, len(section.MusicCarouselShelfRenderer.Contents))

	for _, content := range section.MusicCarouselShelfRenderer.Contents {
		singlesSection.Contents = append(
			singlesSection.Contents,
			artistSingle{
				AlbumId: content.MusicTwoRowItemRenderer.
					NavigationEndpoint.BrowseEndpoint.BrowseID,
				Title:    content.MusicTwoRowItemRenderer.Title.Runs.GetTextField(),
				Subtitle: content.MusicTwoRowItemRenderer.Subtitle.Runs.GetTextField(),
			},
		)
	}
}

type (
	ArtistVideosSection struct {
		Contents          []artistVideo `json:"contents,omitempty"`
		SeeMorePlaylistId string        `json:"seeMorePlaylistId"`
	}
	artistVideo struct {
		VideoId  string `json:"videoId"` // watchEndpoint.videoId
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
	}
)

// scrapes videos section data and sets it
func (videosSection *ArtistVideosSection) ScrapeAndSet(
	wg *sync.WaitGroup,
	section *apiRespSectionContent,
) {
	defer wg.Done()
	if section == nil {
		return
	}

	_, browseId, _ := section.MusicCarouselShelfRenderer.
		Header.MusicCarouselShelfBasicHeaderRenderer.
		Title.Runs.GetNavData(0)
	videosSection.SeeMorePlaylistId = browseId

	videosSection.Contents = make([]artistVideo, 0, len(section.MusicCarouselShelfRenderer.Contents))

	for _, content := range section.MusicCarouselShelfRenderer.Contents {
		videosSection.Contents = append(
			videosSection.Contents,
			artistVideo{
				VideoId: content.MusicTwoRowItemRenderer.
					NavigationEndpoint.WatchEndpoint.VideoID,
				Title:    content.MusicTwoRowItemRenderer.Title.Runs.GetTextField(),
				Subtitle: content.MusicTwoRowItemRenderer.Subtitle.Runs.GetTextField(),
			},
		)
	}
}

type (
	ArtistFeaturedOnSection struct {
		Contents []artistFeaturedOn `json:"contents,omitempty"`
	}
	artistFeaturedOn struct {
		Title      string `json:"title"`
		PlaylistId string `json:"playlistId"`
	}
)

// scrapes featured on section data and sets it
func (featuredOnSection *ArtistFeaturedOnSection) ScrapeAndSet(
	wg *sync.WaitGroup,
	section *apiRespSectionContent,
) {
	defer wg.Done()
	if section == nil {
		return
	}

	featuredOnSection.Contents = make([]artistFeaturedOn, 0, len(section.MusicCarouselShelfRenderer.Contents))

	for _, content := range section.MusicCarouselShelfRenderer.Contents {
		browseEndpoint := content.MusicTwoRowItemRenderer.
			NavigationEndpoint.BrowseEndpoint

		featuredOnSection.Contents = append(
			featuredOnSection.Contents,
			artistFeaturedOn{
				Title:      content.MusicTwoRowItemRenderer.Title.Runs.GetTextField(),
				PlaylistId: browseEndpoint.BrowseID,
			},
		)
	}
}

type (
	AlikeArtistsSection struct {
		Contents []alikeArtist `json:"contents,omitempty"`
	}
	alikeArtist struct {
		Name        string `json:"name"`
		ChannelId   string `json:"channelId"`
		Subscribers string `json:"subscribers"`
	}
)

// scrapes alike artists section data and sets it
func (alikeArtistSection *AlikeArtistsSection) ScrapeAndSet(
	wg *sync.WaitGroup,
	section *apiRespSectionContent,
) {
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
			PageType != utils.MUSIC_PAGE_TYPE_ARTIST {
			continue
		}

		alikeArtistSection.Contents = append(
			alikeArtistSection.Contents,
			alikeArtist{
				Name:      content.MusicTwoRowItemRenderer.Title.Runs.GetTextField(),
				ChannelId: browseEndpoint.BrowseID,
				Subscribers: strings.Split(
					content.MusicTwoRowItemRenderer.Subtitle.Runs.GetTextField(),
					" ",
				)[0],
			},
		)
	}
}
