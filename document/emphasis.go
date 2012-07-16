package document

type Italic struct {
	Content
}

func (i *Italic) Element() Element {
	return i
}

func (i *Italic) Inline() {
}

type Bold struct {
	Content
}

func (b *Bold) Element() Element {
	return b
}

func (b *Bold) Inline() {
}
