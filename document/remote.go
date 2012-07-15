package document

type Image struct {
	URL         string
	Description string
}

func (img *Image) Element() Element {
	return img
}

func (img *Image) Inline() {
}

type Link struct {
	URL string
	Content
}

func (link *Link) Element() Element {
	return link
}

func (link *Link) Inline() {
}
