package lyrics

type ScrapedData struct {
	Lyrics       string `json:"lyrics,omitempty"`
	LyricsFooter string `json:"lyricsFooter,omitempty"`
}

func (sd *ScrapedData) ScrapeAndSet(content APIResp) {
	contents := content.Contents.SectionListRenderer.Contents
	if len(contents) == 0 {
		return
	}

	sd.Lyrics = contents[0].MusicDescriptionShelfRenderer.
		Description.Runs.GetText()

	sd.LyricsFooter = contents[0].MusicDescriptionShelfRenderer.
		Footer.Runs.GetText()
}
