package artist

type SectionName string

const (
	SECTION_ALBUMS        SectionName = "Albums"
	SECTION_SINGLES       SectionName = "Singles"
	SECTION_VIDEOS        SectionName = "Videos"
	SECTION_FEATURED_ON   SectionName = "Featured on"
	SECTION_ALIKE_ARTISTS SectionName = "Fans might also like"
)

var VALID_ARTIST_SECTIONS = map[SectionName]bool{
	SECTION_ALBUMS:        true,
	SECTION_SINGLES:       true,
	SECTION_VIDEOS:        true,
	SECTION_FEATURED_ON:   true,
	SECTION_ALIKE_ARTISTS: true,
}
