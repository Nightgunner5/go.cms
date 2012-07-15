// Credit for the idea for this package goes to https://github.com/daaku/go.h.
package document

type Document struct {
	Title string
	Content
}

func (doc *Document) Element() Element {
	return doc
}
