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

// TODO: add channel id for song or videos
type SongOrVideo struct {
	VideoId   string `json:"videoId"`
	Name      string `json:"name"`
	OtherInfo string `json:"otherInfo"`
}

type Artist struct {
	Name            string `json:"name"`
	OtherInfo       string `json:"otherInfo"`
	ArtistChannelId string `json:"artistChannelId"`
}

type Album struct {
	Name      string `json:"name"`
	OtherInfo string `json:"otherInfo"`
}

type CommunityPlaylist struct {
	Name      string `json:"name"`
	OtherInfo string `json:"otherInfo"`
}

type FeaturedPlaylist struct {
	Name      string `json:"name"`
	OtherInfo string `json:"otherInfo"`
}
