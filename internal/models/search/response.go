package search

import (
	"strings"

	"github.com/ghoshRitesh12/brooktube/internal/constants"
)

type ScrapedData struct {
	Title string `json:"title"`
	// for community playlists, songs, albums, videos, etc.
	Content ResultContent `json:"content,omitempty"`
	// used as ctoken and continuation query params for getting paginated data
	ContinuationToken string `json:"continuation,omitempty"`
}

type ResultContent struct {
	Songs              Songs              `json:"songs"`
	Videos             Videos             `json:"videos"`
	Artists            Artists            `json:"artists"`
	Albums             Albums             `json:"albums"`
	CommunityPlaylists CommunityPlaylists `json:"communityPlaylists"`
	FeaturedPlaylists  FeaturedPlaylists  `json:"featuredPlaylists"`
}

type (
	SongOrVideo struct {
		SongOrVideoId string `json:"songOrVideoId"`
		Name          string `json:"name"`
		AlbumName     string `json:"albumName"`
		AlbumId       string `json:"albumId"`
		Duration      string `json:"duration"`
		Interactions  string `json:"interactions"`

		ChannelName string `json:"channelName"`
		ChannelId   string `json:"channelId"`
	}
	Songs []SongOrVideo
)

func (songs *Songs) ScrapeAndSet(
	shelfContents []APIRespSectionContent,
) {
	preSongs := make(Songs, 0, len(shelfContents))
	*songs = preSongs

	for _, content := range shelfContents {
		songOrVideo := SongOrVideo{
			SongOrVideoId: content.MusicResponsiveListItemRenderer.PlaylistItemData.VideoId,
		}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.Text.Runs

			if i == 0 {
				songOrVideo.Name = textRuns.GetText(0)
				continue
			}

			if i == 1 {
				songOrVideo.ChannelName = textRuns.GetText(0)
				songOrVideo.AlbumName = textRuns.GetText(2)

				pageType, browseId, _ := flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs.GetNavData(0)

				if (pageType == constants.MUSIC_PAGE_TYPE_ARTIST) ||
					(pageType == constants.MUSIC_PAGE_TYPE_USER_CHANNEL) {
					songOrVideo.ChannelId = browseId
				}

				{
					pageType, browseId, _ := flexColumn.
						MusicResponsiveListItemFlexColumnRenderer.
						Text.Runs.GetNavData(2)

					if pageType == constants.MUSIC_PAGE_TYPE_ALBUM {
						songOrVideo.AlbumId = browseId
					}
				}

				songOrVideo.Duration = textRuns.GetText(uint8(len(textRuns) - 1))
				songOrVideo.Interactions = textRuns.GetText(2)

				continue
			}
		}

		*songs = append(*songs, songOrVideo)
	}
}

type Videos []SongOrVideo

func (videos *Videos) ScrapeAndSet(
	shelfContents []APIRespSectionContent,
) {
	preVideos := make(Videos, 0, len(shelfContents))
	*videos = preVideos

	for _, content := range shelfContents {
		songOrVideo := SongOrVideo{
			SongOrVideoId: content.MusicResponsiveListItemRenderer.PlaylistItemData.VideoId,
		}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.Text.Runs

			if i == 0 {
				songOrVideo.Name = textRuns.GetText(0)
				continue
			}

			if i == 1 {
				songOrVideo.ChannelName = textRuns.GetText(0)

				pageType, browseId, _ := flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs.GetNavData(0)

				if (pageType == constants.MUSIC_PAGE_TYPE_ARTIST) ||
					(pageType == constants.MUSIC_PAGE_TYPE_USER_CHANNEL) {
					songOrVideo.ChannelId = browseId
				}

				songOrVideo.Duration = textRuns.GetText(uint8(len(textRuns) - 1))
				songOrVideo.Interactions = textRuns.GetText(2)

				continue
			}

			songOrVideo.Interactions = textRuns.GetText(0)
		}

		*videos = append(*videos, songOrVideo)
	}
}

type (
	Artist struct {
		Name        string `json:"name"`
		Subscribers string `json:"subscribers"`
		ChannelId   string `json:"channelId"`
	}
	Artists []Artist
)

func (artists *Artists) ScrapeAndSet(
	shelfContents []APIRespSectionContent,
) {
	preArtists := make(Artists, 0, len(shelfContents))
	*artists = preArtists

	for _, content := range shelfContents {
		artist := Artist{}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.
				Text.Runs

			if i == 0 {
				artist.Name = flexColumn.
					MusicResponsiveListItemFlexColumnRenderer.
					Text.Runs.GetText()
				continue
			}

			artist.Subscribers = strings.Split(
				textRuns.GetText(2),
				" ",
			)[0]
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
	Album struct {
		Name            string `json:"name"`
		OtherInfo       string `json:"otherInfo"`
		ArtistName      string `json:"artistName"`
		ArtistChannelId string `json:"artistChannelId"`
		YearOfRelease   string `json:"yearOfRelease"`
	}
	Albums []Album
)

func (albums *Albums) ScrapeAndSet(shelfContents []APIRespSectionContent) {
	preArtists := make(Albums, 0, len(shelfContents))
	*albums = preArtists

	otherInfoBuilder := strings.Builder{}

	for _, content := range shelfContents {
		album := Album{}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.MusicResponsiveListItemFlexColumnRenderer.Text.Runs

			if i == 0 {
				album.Name = textRuns.GetText()
				continue
			}

			otherInfoBuilder.WriteString(textRuns.GetText())

			pageType, browseId, _ := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.
				Text.Runs.GetNavData(2)

			if pageType == constants.MUSIC_PAGE_TYPE_ARTIST {
				album.ArtistChannelId = browseId
			}

			album.YearOfRelease = textRuns.GetText(uint8(len(textRuns) - 1))
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
	CommunityPlaylist struct {
		Name            string `json:"name"`
		OtherInfo       string `json:"otherInfo"`
		PlaylistId      string `json:"playlistId"`
		ArtistChannelId string `json:"artistChannelId"`
		Interactions    string `json:"interactions"`
	}
	CommunityPlaylists []CommunityPlaylist
)

func (communityPlaylists *CommunityPlaylists) ScrapeAndSet(shelfContents []APIRespSectionContent) {
	preCommunityPlaylists := make(CommunityPlaylists, 0, len(shelfContents))
	*communityPlaylists = preCommunityPlaylists

	otherInfoBuilder := strings.Builder{}

	for _, content := range shelfContents {
		communityPlaylist := CommunityPlaylist{}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.
				Text.Runs

			if i == 0 {
				communityPlaylist.Name = textRuns.GetText()
				continue
			}

			otherInfoBuilder.WriteString(textRuns.GetText())

			_, browseId, _ := textRuns.GetNavData(0)
			communityPlaylist.ArtistChannelId = browseId
			communityPlaylist.Interactions = textRuns.GetText(uint8(len(textRuns) - 1))
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
	FeaturedPlaylist struct {
		Name      string `json:"name"`
		OtherInfo string `json:"otherInfo"`
	}
	FeaturedPlaylists []FeaturedPlaylist
)

func (featuredPlaylists *FeaturedPlaylists) ScrapeAndSet(
	shelfContents []APIRespSectionContent,
) {
	preCommunityPlaylists := make(FeaturedPlaylists, 0, len(shelfContents))
	*featuredPlaylists = preCommunityPlaylists

	otherInfoBuilder := strings.Builder{}

	for _, content := range shelfContents {
		featuredPlaylist := FeaturedPlaylist{}

		for i, flexColumn := range content.MusicResponsiveListItemRenderer.FlexColumns {
			textRuns := flexColumn.
				MusicResponsiveListItemFlexColumnRenderer.
				Text.Runs

			if i == 0 {
				featuredPlaylist.Name = textRuns.GetText()
				continue
			}

			otherInfoBuilder.WriteString(textRuns.GetText())
		}

		featuredPlaylist.OtherInfo = otherInfoBuilder.String()
		*featuredPlaylists = append(*featuredPlaylists, featuredPlaylist)
		otherInfoBuilder.Reset()
	}
}
