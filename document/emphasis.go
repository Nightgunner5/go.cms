// Credit for the idea for this package goes to https://github.com/daaku/go.h.
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

type Underline struct {
	Content
}

func (u *Underline) Element() Element {
	return u
}

func (u *Underline) Inline() {
}
