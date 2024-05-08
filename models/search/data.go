package search

type Result struct {
	Title string `json:"title"`
	// for community playlists, songs, albums, videos, etc.
	Content ResultContent `json:"content,omitempty"`
	// used as ctoken and continuation query params for getting paginated data
	ContinuationToken string `json:"continuation,omitempty"`
}

type ResultContent struct {
	SongOrVideos       []SongOrVideo       `json:"songOrVideos"`
	Artists            []Artist            `json:"artists"`
	Albums             []Album             `json:"albums"`
	CommunityPlaylists []CommunityPlaylist `json:"communityPlaylists"`
	FeaturedPlaylists  []FeaturedPlaylist  `json:"featuredPlaylists"`
}

type SongOrVideo struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	AlbumName       string `json:"albumName"`
	AlbumId         string `json:"albumId"`
	ArtistName      string `json:"artistName"`
	ArtistChannelId string `json:"artistChannelId"`
	Duration        string `json:"duration"`
	PlaysCount      string `json:"playsCount"`
}

type Artist struct {
	Name            string `json:"name"`
	OtherInfo       string `json:"otherInfo"`
	ArtistChannelId string `json:"artistChannelId"`
}

type Album struct {
	Name            string `json:"name"`
	OtherInfo       string `json:"otherInfo"`
	ArtistChannelId string `json:"artistChannelId"`
}

type CommunityPlaylist struct {
	Name       string `json:"name"`
	OtherInfo  string `json:"otherInfo"`
	PlaylistId string `json:"playlistId"`
}

type FeaturedPlaylist struct {
	Name      string `json:"name"`
	OtherInfo string `json:"otherInfo"`
}
