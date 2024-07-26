package artist

import (
	"strings"
	"sync"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/models"
	"github.com/ghoshRitesh12/brooktube/internal/models/search"
	"github.com/ghoshRitesh12/brooktube/internal/utils"
)

type ScrapedData struct {
	Info struct {
		Name             string               `json:"name"`
		About            string               `json:"about"`
		Views            string               `json:"views"`
		Subscribers      string               `json:"subscribersCount"`
		BackgroundImages models.AppThumbnails `json:"backgroundImages,omitempty"`
		ForegroundImages models.AppThumbnails `json:"foregroundImages,omitempty"`
	} `json:"info,omitempty"`

	Songs        Songs        `json:"songs,omitempty"`        // section
	Videos       Videos       `json:"videos,omitempty"`       // section
	Albums       Albums       `json:"albums,omitempty"`       // section
	Singles      Singles      `json:"singles,omitempty"`      // section
	FeaturedOns  FeaturedOns  `json:"featuredOns,omitempty"`  // section
	AlikeArtists AlikeArtists `json:"alikeArtists,omitempty"` // section
	Playlists    Playlists    `json:"playlists,omitempty"`    // section
}

// scrapes and sets basic info of the artist
func (artist *ScrapedData) ScrapeAndSetBasicInfo(
	wg *sync.WaitGroup,
	header *apiRespHeader,
	sections *[]apiRespSectionContent,
) {
	defer wg.Done()

	// for artists
	if header.MusicImmersiveHeaderRenderer.Title.Runs.GetText() != "" {
		artist.Info.Name = header.MusicImmersiveHeaderRenderer.Title.Runs.GetText()
		artist.Info.About = header.MusicImmersiveHeaderRenderer.Description.Runs.GetText()
		artist.Info.Subscribers = header.MusicImmersiveHeaderRenderer.
			SubscriptionButton.SubscribeButtonRenderer.
			SubscriberCountText.Runs.GetText()
		artist.Info.BackgroundImages = header.MusicImmersiveHeaderRenderer.
			Thumbnail.GetAllThumbnails()
		artist.Info.Views = strings.Split(
			(*sections)[len(*sections)-1].
				MusicDescriptionShelfRenderer.Subheader.
				Runs.GetText(),
			" ",
		)[0]

		return
	} else if header.MusicImmersiveHeaderRenderer.Title.Runs.GetText() == "" { // for non artists
		artist.Info.Name = header.MusicVisualHeaderRenderer.Title.Runs.GetText()
		artist.Info.BackgroundImages = header.MusicVisualHeaderRenderer.
			Thumbnail.GetAllThumbnails()
		artist.Info.ForegroundImages = header.MusicVisualHeaderRenderer.
			ForegroundThumbnail.GetAllThumbnails()

		return
	}
}

type Songs struct {
	Contents          search.Songs `json:"contents,omitempty"`
	SeeMorePlaylistId string       `json:"seeMorePlaylistId"`
}

// scrapes songs section data and sets it
func (songs *Songs) ScrapeAndSet(wg *sync.WaitGroup, renderer *APIRespMusicShelfRenderer) {
	defer wg.Done()

	if renderer == nil {
		return
	}
	if renderer.Title.Runs.GetText() != "Songs" {
		return
	}

	_, browseId := renderer.Title.Runs.GetNavData(0)
	songs.SeeMorePlaylistId = browseId
	songs.Contents = utils.ParseArtistSongContents(&(renderer.Contents))
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
		AlbumId   string               `json:"albumId"` // browseEndpoint.browseId
		Title     string               `json:"title"`
		Subtitle  string               `json:"subtitle"`
		CoverArts models.AppThumbnails `json:"coverArts"`
	}
)

// scrapes albums section data and sets it
func (albums *Albums) ScrapeAndSet(wg *sync.WaitGroup, renderer *APIRespMusicCarouselShelfRenderer) {
	defer wg.Done()

	if renderer == nil {
		return
	}

	_, browseId, browseParams := renderer.Header.
		MusicCarouselShelfBasicHeaderRenderer.
		Title.Runs.GetNavData(0)
	albums.SeeMoreEndpoint.DiscographyId = browseId
	albums.SeeMoreEndpoint.Params = browseParams

	albums.Contents = make([]album, 0, len(renderer.Contents))

	for _, content := range renderer.Contents {
		albums.Contents = append(albums.Contents,
			album{
				AlbumId: content.MusicTwoRowItemRenderer.
					NavigationEndpoint.BrowseEndpoint.BrowseID,
				Title:     content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				Subtitle:  content.MusicTwoRowItemRenderer.Subtitle.Runs.GetText(),
				CoverArts: content.MusicTwoRowItemRenderer.ThumbnailRenderer.GetAllThumbnails(),
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
		AlbumId    string               `json:"albumId"` // browseEndpoint.browseId
		Title      string               `json:"title"`
		Subtitle   string               `json:"subtitle"`
		Thumbnails models.AppThumbnails `json:"thumbnails"`
	}
)

// scrapes singles section data and sets it
func (singles *Singles) ScrapeAndSet(wg *sync.WaitGroup, renderer *APIRespMusicCarouselShelfRenderer) {
	defer wg.Done()

	if renderer == nil {
		return
	}

	_, browseId, browseParams := renderer.Header.
		MusicCarouselShelfBasicHeaderRenderer.
		Title.Runs.GetNavData(0)
	singles.SeeMoreEndpoint.DiscographyId = browseId
	singles.SeeMoreEndpoint.Params = browseParams

	singles.Contents = make([]single, 0, len(renderer.Contents))

	for _, content := range renderer.Contents {
		singles.Contents = append(singles.Contents,
			single{
				AlbumId: content.MusicTwoRowItemRenderer.
					NavigationEndpoint.BrowseEndpoint.BrowseID,
				Title:      content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				Subtitle:   content.MusicTwoRowItemRenderer.Subtitle.Runs.GetText(),
				Thumbnails: content.MusicTwoRowItemRenderer.ThumbnailRenderer.GetAllThumbnails(),
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
		VideoId    string               `json:"videoId"` // watchEndpoint.videoId
		Title      string               `json:"title"`
		Subtitle   string               `json:"subtitle"`
		Thumbnails models.AppThumbnails `json:"thumbnails"`
	}
)

// scrapes videos section data and sets it
func (videos *Videos) ScrapeAndSet(wg *sync.WaitGroup, renderer *APIRespMusicCarouselShelfRenderer) {
	defer wg.Done()

	if renderer == nil {
		return
	}

	_, browseId, _ := renderer.Header.
		MusicCarouselShelfBasicHeaderRenderer.
		Title.Runs.GetNavData(0)
	videos.SeeMorePlaylistId = browseId

	videos.Contents = make([]video, 0, len(renderer.Contents))

	for _, content := range renderer.Contents {
		videos.Contents = append(videos.Contents,
			video{
				VideoId: content.MusicTwoRowItemRenderer.
					NavigationEndpoint.WatchEndpoint.VideoID,
				Title:      content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				Subtitle:   content.MusicTwoRowItemRenderer.Subtitle.Runs.GetText(),
				Thumbnails: content.MusicTwoRowItemRenderer.ThumbnailRenderer.GetAllThumbnails(),
			},
		)
	}
}

type (
	FeaturedOns struct {
		Contents []featuredOn `json:"contents,omitempty"`
	}
	featuredOn struct {
		Title      string               `json:"title"`
		PlaylistId string               `json:"playlistId"`
		Thumbnails models.AppThumbnails `json:"thumbnails"`
	}
)

// scrapes featured on section data and sets it
func (featuredOns *FeaturedOns) ScrapeAndSet(wg *sync.WaitGroup, renderer *APIRespMusicCarouselShelfRenderer) {
	defer wg.Done()

	if renderer == nil {
		return
	}

	featuredOns.Contents = make([]featuredOn, 0, len(renderer.Contents))

	for _, content := range renderer.Contents {
		browseEndpoint := content.MusicTwoRowItemRenderer.
			NavigationEndpoint.BrowseEndpoint

		featuredOns.Contents = append(featuredOns.Contents,
			featuredOn{
				Title:      content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				PlaylistId: browseEndpoint.BrowseID,
				Thumbnails: content.MusicTwoRowItemRenderer.ThumbnailRenderer.GetAllThumbnails(),
			},
		)
	}
}

type (
	AlikeArtists struct {
		Contents []alikeArtist `json:"contents,omitempty"`
	}
	alikeArtist struct {
		Name        string               `json:"name"`
		ChannelId   string               `json:"channelId"`
		Subscribers string               `json:"subscribers"`
		Images      models.AppThumbnails `json:"images"`
	}
)

// scrapes alike artists section data and sets it
func (alikeArtists *AlikeArtists) ScrapeAndSet(wg *sync.WaitGroup, renderer *APIRespMusicCarouselShelfRenderer) {
	defer wg.Done()

	if renderer == nil {
		return
	}

	// to pre-allocate memory
	alikeArtists.Contents = make([]alikeArtist, 0, len(renderer.Contents))

	for _, content := range renderer.Contents {
		browseEndpoint := content.MusicTwoRowItemRenderer.
			NavigationEndpoint.BrowseEndpoint

		if browseEndpoint.
			BrowseEndpointContextSupportedConfigs.
			BrowseEndpointContextMusicConfig.
			PageType != constants.MUSIC_PAGE_TYPE_ARTIST {
			continue
		}
		content.MusicTwoRowItemRenderer.ThumbnailRenderer.GetAllThumbnails()

		alikeArtists.Contents = append(alikeArtists.Contents,
			alikeArtist{
				Name:      content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				ChannelId: browseEndpoint.BrowseID,
				Subscribers: strings.Split(
					content.MusicTwoRowItemRenderer.Subtitle.Runs.GetText(),
					" ",
				)[0],
				Images: content.MusicTwoRowItemRenderer.ThumbnailRenderer.GetAllThumbnails(),
			},
		)
	}
}

type (
	Playlists struct {
		Contents        []playlist `json:"contents,omitempty"`
		SeeMoreEndpoint struct {
			Params        string `json:"params"`
			DiscographyId string `json:"discographyId"` // somewhat like playlistId
		} `json:"seeMoreEndpoint"`
	}
	playlist struct {
		PlaylistId string               `json:"albumId"` // browseEndpoint.browseId
		Title      string               `json:"title"`
		Subtitle   string               `json:"subtitle"`
		CoverArts  models.AppThumbnails `json:"coverArts"`
	}
)

// scrapes alike artists section data and sets it
func (playlists *Playlists) ScrapeAndSet(wg *sync.WaitGroup, renderer *APIRespMusicCarouselShelfRenderer) {
	defer wg.Done()

	if renderer == nil {
		return
	}

	_, browseId, browseParams := renderer.Header.
		MusicCarouselShelfBasicHeaderRenderer.
		Title.Runs.GetNavData(0)
	playlists.SeeMoreEndpoint.DiscographyId = browseId
	playlists.SeeMoreEndpoint.Params = browseParams

	playlists.Contents = make([]playlist, 0, len(renderer.Contents))

	for _, content := range renderer.Contents {
		playlists.Contents = append(playlists.Contents,
			playlist{
				PlaylistId: content.MusicTwoRowItemRenderer.
					NavigationEndpoint.BrowseEndpoint.BrowseID,
				Title:     content.MusicTwoRowItemRenderer.Title.Runs.GetText(),
				Subtitle:  content.MusicTwoRowItemRenderer.Subtitle.Runs.GetText(),
				CoverArts: content.MusicTwoRowItemRenderer.ThumbnailRenderer.GetAllThumbnails(),
			},
		)
	}
}
