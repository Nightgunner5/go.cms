// Credit for the idea for this package goes to https://github.com/daaku/go.h.
package document

// Dummy interface to allow somewhat stronger typing.
type Element interface {
	Element() Element
}

type InlineElement interface {
	Inline()
}

// A leaf element which contains only text.
type LeafElement struct {
	Text string
}

func (el *LeafElement) Element() Element {
	return el
}

func (el *LeafElement) Inline() {
}

type Content []Element

// A container element which may contain any number of children
// which may or may not be Containers.
type Container interface {
	Element
	Contents() Content
	AddChild(Element)
	RemoveChild(Element) // The child is only to be removed if it is a direct descendant of this container.
}

func (con *Content) Contents() Content {
	return *con
}

func (con *Content) AddChild(el Element) {
	*con = append(*con, el)
}

func (con *Content) RemoveChild(el Element) {
	for i, child := range *con {
		if child == el {
			*con = append((*con)[:i], (*con)[i+1:]...)
			return
		}
	}
}
