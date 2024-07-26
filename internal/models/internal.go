package models

type (
	AppThumbnail struct {
		URL    string `json:"url,omitempty"`
		Width  int    `json:"width,omitempty"`
		Height int    `json:"height,omitempty"`
	}
	AppThumbnails []AppThumbnail
)

func (thumbnails *AppThumbnails) GetThumbnailWithIndex(index int) string {
	url := ""
	defaultIndex := 0

	if len(*thumbnails) == 0 {
		return url
	}

	if len(*thumbnails) > 0 && index < len(*thumbnails) {
		url = (*thumbnails)[index].URL
	} else if len(*thumbnails) > 0 && index > len(*thumbnails) {
		url = (*thumbnails)[defaultIndex].URL
	}
	return url
}

func (thumbnails *AppThumbnails) GetThumbnailWithDimension(width, height int) string {
	url := ""
	if len(*thumbnails) == 0 {
		return url
	}

	for _, thumbnail := range *thumbnails {
		if thumbnail.Width == width && thumbnail.Height == height {
			url = thumbnail.URL
			break
		}
	}

	return url
}
