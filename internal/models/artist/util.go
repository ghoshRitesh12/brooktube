package artist

type SectionName string

const (
	SECTION_SONGS         SectionName = "Songs"
	SECTION_ALBUMS        SectionName = "Albums"
	SECTION_SINGLES       SectionName = "Singles"
	SECTION_VIDEOS        SectionName = "Videos"
	SECTION_FEATURED_ON   SectionName = "Featured on"
	SECTION_ALIKE_ARTISTS SectionName = "Fans might also like"
	SECTION_PLAYLISTS     SectionName = "Playlists" // outlier for artists but common for other channels
)

var VALID_ARTIST_SECTIONS = map[SectionName]bool{
	SECTION_SONGS:         true,
	SECTION_ALBUMS:        true,
	SECTION_SINGLES:       true,
	SECTION_VIDEOS:        true,
	SECTION_FEATURED_ON:   true,
	SECTION_ALIKE_ARTISTS: true,
	SECTION_PLAYLISTS:     true,
}
