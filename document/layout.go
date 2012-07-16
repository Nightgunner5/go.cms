// Credit for the idea for this package goes to https://github.com/daaku/go.h.
package document

import "time"

type LineBreak struct {
}

func (lb *LineBreak) Element() Element {
	return lb
}

type Paragraph struct {
	Content
}

func (para *Paragraph) Element() Element {
	return para
}

type Article struct {
	Title     string
	Timestamp time.Time
	Content
}

func (article *Article) Element() Element {
	return article
}
