package search

import (
	"strings"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
	"github.com/ghoshRitesh12/brooktube/internal/models"
)

type SearchResult interface { // result
	SetBasicInfo(shelfRenderer *apiRespSection)
	SetContinuationToken(contents *continuationContents)
}

type SearchContent interface { // content
	ScrapeAndSet(shelfContents *[]APIRespSectionContent)
}

type SearchInfo struct {
	Title string `json:"title"`
	// used as ctoken and continuation query params for getting paginated data
	ContinuationToken string `json:"continuationToken,omitempty"`
}

type (
	ScrapedSongResult struct {
		Info     SearchInfo `json:"info"`
		Contents Songs      `json:"contents"`
	}
	Song struct {
		VideoId      string               `json:"videoId"`
		Name         string               `json:"name"`
		AlbumName    string               `json:"albumName"`
		AlbumId      string               `json:"albumId"`
		Duration     string               `json:"duration"`
		Interactions string               `json:"interactions"`
		Thumbnails   models.AppThumbnails `json:"thumbnails"`

		ChannelName string `json:"channelName"`
		ChannelId   string `json:"channelId"`
	}
	Songs []Song
)

func (data *ScrapedSongResult) SetBasicInfo(shelfRenderer *apiRespSection) {
	if shelfRenderer == nil {
		return
	}
	data.Info.Title = shelfRenderer.MusicShelfRenderer.Title.Runs.GetText()
	data.Info.ContinuationToken = shelfRenderer.MusicShelfRenderer.
		Continuations.GetContinuationToken()
}

// for continuation token funcs only
func (data *ScrapedSongResult) SetContinuationToken(contents *continuationContents) {
	if contents == nil {
		return
	}
	data.Info.ContinuationToken = contents.MusicShelfContinuation.
		Continuations.GetContinuationToken()
}

func (songs *Songs) ScrapeAndSet(shelfContents *[]APIRespSectionContent) {
	*songs = make(Songs, 0, len(*shelfContents))

	for _, content := range *shelfContents {
		song := Song{
			VideoId:    content.MusicResponsiveListItemRenderer.PlaylistItemData.VideoId,
			Thumbnails: content.MusicResponsiveListItemRenderer.Thumbnail.GetAllThumbnails(),
		}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs

			switch i {
			case 0:
				song.Name = textRuns.GetText(0)
			case 1:
				song.ChannelName = textRuns.GetText(0)
				song.AlbumName = textRuns.GetText(2)

				pageType, browseId, _ := flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs.GetNavData(0)

				if (pageType == constants.MUSIC_PAGE_TYPE_ARTIST) ||
					(pageType == constants.MUSIC_PAGE_TYPE_USER_CHANNEL) {
					song.ChannelId = browseId
				}

				// re-using variables
				pageType, browseId, _ = flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs.GetNavData(2)

				if pageType == constants.MUSIC_PAGE_TYPE_ALBUM {
					song.AlbumId = browseId
				}

				song.Duration = textRuns.GetText(uint8(len(textRuns) - 1))
			case 2:
				song.Interactions = textRuns.GetText(0)
			}
		}

		*songs = append(*songs, song)
	}
}

type (
	ScrapedVideoResult struct {
		Info     SearchInfo `json:"info"`
		Contents Videos     `json:"contents"`
	}
	Video struct {
		VideoId      string `json:"videoId"`
		Name         string `json:"name"`
		Duration     string `json:"duration"`
		Interactions string `json:"interactions"`

		ChannelName string `json:"channelName"`
		ChannelId   string `json:"channelId"`

		Thumbnails models.AppThumbnails `json:"thumbnails"`
	}
	Videos []Video
)

func (data *ScrapedVideoResult) SetBasicInfo(shelfRenderer *apiRespSection) {
	if shelfRenderer == nil {
		return
	}
	data.Info.Title = shelfRenderer.MusicShelfRenderer.Title.Runs.GetText()
	data.Info.ContinuationToken = shelfRenderer.MusicShelfRenderer.
		Continuations.GetContinuationToken()
}

// for continuation token funcs only
func (data *ScrapedVideoResult) SetContinuationToken(contents *continuationContents) {
	if contents == nil {
		return
	}
	data.Info.ContinuationToken = contents.MusicShelfContinuation.
		Continuations.GetContinuationToken()
}

func (videos *Videos) ScrapeAndSet(shelfContents *[]APIRespSectionContent) {
	*videos = make(Videos, 0, len(*shelfContents))

	for _, content := range *shelfContents {
		video := Video{
			VideoId:    content.MusicResponsiveListItemRenderer.PlaylistItemData.VideoId,
			Thumbnails: content.MusicResponsiveListItemRenderer.Thumbnail.GetAllThumbnails(),
		}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.Text.Runs

			switch i {
			case 0:
				video.Name = textRuns.GetText(0)
			case 1:
				video.ChannelName = textRuns.GetText(0)

				pageType, browseId, _ := flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs.GetNavData(0)

				if (pageType == constants.MUSIC_PAGE_TYPE_ARTIST) ||
					(pageType == constants.MUSIC_PAGE_TYPE_USER_CHANNEL) {
					video.ChannelId = browseId
				}

				video.Duration = textRuns.GetText(uint8(len(textRuns) - 1))
				video.Interactions = textRuns.GetText(2)
			case 2:
				video.Interactions = textRuns.GetText(0)
			}
		}

		*videos = append(*videos, video)
	}
}

type (
	ScrapedArtistResult struct {
		Info     SearchInfo `json:"info"`
		Contents Artists    `json:"contents"`
	}
	Artist struct {
		Name        string `json:"name"`
		Subscribers string `json:"subscribers"`
		ChannelId   string `json:"channelId"`
	}
	Artists []Artist
)

func (data *ScrapedArtistResult) SetBasicInfo(shelfRenderer *apiRespSection) {
	if shelfRenderer == nil {
		return
	}
	data.Info.Title = shelfRenderer.MusicShelfRenderer.Title.Runs.GetText()
	data.Info.ContinuationToken = shelfRenderer.MusicShelfRenderer.
		Continuations.GetContinuationToken()
}

// for continuation token funcs only
func (data *ScrapedArtistResult) SetContinuationToken(contents *continuationContents) {
	if contents == nil {
		return
	}
	data.Info.ContinuationToken = contents.MusicShelfContinuation.
		Continuations.GetContinuationToken()
}

func (artists *Artists) ScrapeAndSet(shelfContents *[]APIRespSectionContent) {
	*artists = make(Artists, 0, len(*shelfContents))

	for _, content := range *shelfContents {
		artist := Artist{}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.
				Text.Runs

			switch i {
			case 0:
				artist.Name = flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs.GetText()
			case 1:
				artist.Subscribers = strings.Split(
					textRuns.GetText(2),
					" ",
				)[0]
			}
		}

		browseEndpoint := content.MusicResponsiveListItemRenderer.
			NavigationEndpoint.BrowseEndpoint

		if browseEndpoint.
			BrowseEndpointContextSupportedConfigs.
			BrowseEndpointContextMusicConfig.
			PageType == constants.MUSIC_PAGE_TYPE_ARTIST {
			artist.ChannelId = browseEndpoint.BrowseID
		}

		*artists = append(*artists, artist)
	}
}

type (
	ScrapedAlbumResult struct {
		Info     SearchInfo `json:"info"`
		Contents Albums     `json:"contents"`
	}
	Album struct {
		Name            string `json:"name"`
		OtherInfo       string `json:"otherInfo"`
		ArtistName      string `json:"artistName"`
		ArtistChannelId string `json:"artistChannelId"`
		YearOfRelease   string `json:"yearOfRelease"`
	}
	Albums []Album
)

func (data *ScrapedAlbumResult) SetBasicInfo(shelfRenderer *apiRespSection) {
	if shelfRenderer == nil {
		return
	}
	data.Info.Title = shelfRenderer.MusicShelfRenderer.Title.Runs.GetText()
	data.Info.ContinuationToken = shelfRenderer.MusicShelfRenderer.
		Continuations.GetContinuationToken()
}

// for continuation token funcs only
func (data *ScrapedAlbumResult) SetContinuationToken(contents *continuationContents) {
	if contents == nil {
		return
	}
	data.Info.ContinuationToken = contents.MusicShelfContinuation.
		Continuations.GetContinuationToken()
}

func (albums *Albums) ScrapeAndSet(shelfContents *[]APIRespSectionContent) {
	*albums = make(Albums, 0, len(*shelfContents))
	otherInfoBuilder := strings.Builder{}

	for _, content := range *shelfContents {
		album := Album{}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs

			switch i {
			case 0:
				album.Name = textRuns.GetText()
			case 1:
				otherInfoBuilder.WriteString(textRuns.GetText())

				pageType, browseId, _ := flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs.GetNavData(2)

				if pageType == constants.MUSIC_PAGE_TYPE_ARTIST {
					album.ArtistChannelId = browseId
				}

				album.YearOfRelease = textRuns.GetText(uint8(len(textRuns) - 1))
			}
		}

		album.OtherInfo = otherInfoBuilder.String()
		album.ArtistName = strings.Split(
			album.OtherInfo,
			constants.OTHER_INFO_SEPARATOR,
		)[1]

		*albums = append(*albums, album)
		otherInfoBuilder.Reset()
	}
}

type (
	ScrapedCommunityPlaylistResult struct {
		Info     SearchInfo         `json:"info"`
		Contents CommunityPlaylists `json:"contents"`
	}
	CommunityPlaylist struct {
		Name            string `json:"name"`
		OtherInfo       string `json:"otherInfo"`
		PlaylistId      string `json:"playlistId"`
		ArtistChannelId string `json:"artistChannelId"`
		Interactions    string `json:"interactions"`
	}
	CommunityPlaylists []CommunityPlaylist
)

func (data *ScrapedCommunityPlaylistResult) SetBasicInfo(shelfRenderer *apiRespSection) {
	if shelfRenderer == nil {
		return
	}
	data.Info.Title = shelfRenderer.MusicShelfRenderer.Title.Runs.GetText()
	data.Info.ContinuationToken = shelfRenderer.MusicShelfRenderer.
		Continuations.GetContinuationToken()
}

// for continuation token funcs only
func (data *ScrapedCommunityPlaylistResult) SetContinuationToken(contents *continuationContents) {
	if contents == nil {
		return
	}
	data.Info.ContinuationToken = contents.MusicShelfContinuation.
		Continuations.GetContinuationToken()
}

func (communityPlaylists *CommunityPlaylists) ScrapeAndSet(shelfContents *[]APIRespSectionContent) {
	*communityPlaylists = make(CommunityPlaylists, 0, len(*shelfContents))
	otherInfoBuilder := strings.Builder{}

	for _, content := range *shelfContents {
		communityPlaylist := CommunityPlaylist{}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.
				Text.Runs

			switch i {
			case 0:
				communityPlaylist.Name = textRuns.GetText()
			case 1:
				otherInfoBuilder.WriteString(textRuns.GetText())

				_, browseId, _ := textRuns.GetNavData(0)
				communityPlaylist.ArtistChannelId = browseId
				communityPlaylist.Interactions = textRuns.GetText(uint8(len(textRuns) - 1))
			}
		}

		playlistBrowseEndpoint := content.
			MusicResponsiveListItemRenderer.
			NavigationEndpoint.BrowseEndpoint

		if playlistBrowseEndpoint.
			BrowseEndpointContextSupportedConfigs.
			BrowseEndpointContextMusicConfig.
			PageType == constants.MUSIC_PAGE_TYPE_PLAYLIST {
			communityPlaylist.PlaylistId = playlistBrowseEndpoint.BrowseID
		}

		communityPlaylist.OtherInfo = otherInfoBuilder.String()
		*communityPlaylists = append(*communityPlaylists, communityPlaylist)
		otherInfoBuilder.Reset()
	}
}

type (
	ScrapedFeaturedPlaylistResult struct {
		Info     SearchInfo        `json:"info"`
		Contents FeaturedPlaylists `json:"contents"`
	}
	FeaturedPlaylist struct {
		Name      string `json:"name"`
		OtherInfo string `json:"otherInfo"`
	}
	FeaturedPlaylists []FeaturedPlaylist
)

func (data *ScrapedFeaturedPlaylistResult) SetBasicInfo(shelfRenderer *apiRespSection) {
	if shelfRenderer == nil {
		return
	}
	data.Info.Title = shelfRenderer.MusicShelfRenderer.Title.Runs.GetText()
	data.Info.ContinuationToken = shelfRenderer.MusicShelfRenderer.
		Continuations.GetContinuationToken()
}

// for continuation token funcs only
func (data *ScrapedFeaturedPlaylistResult) SetContinuationToken(contents *continuationContents) {
	if contents == nil {
		return
	}
	data.Info.ContinuationToken = contents.MusicShelfContinuation.
		Continuations.GetContinuationToken()
}

func (featuredPlaylists *FeaturedPlaylists) ScrapeAndSet(shelfContents *[]APIRespSectionContent) {
	*featuredPlaylists = make(FeaturedPlaylists, 0, len(*shelfContents))
	otherInfoBuilder := strings.Builder{}

	for _, content := range *shelfContents {
		featuredPlaylist := FeaturedPlaylist{}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.
				Text.Runs

			switch i {
			case 0:
				featuredPlaylist.Name = textRuns.GetText()
			case 1:
				otherInfoBuilder.WriteString(textRuns.GetText())
			}
		}

		featuredPlaylist.OtherInfo = otherInfoBuilder.String()
		*featuredPlaylists = append(*featuredPlaylists, featuredPlaylist)
		otherInfoBuilder.Reset()
	}
}
