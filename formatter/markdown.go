package formatter

import (
	"fmt"
	"html"
	. "llamaslayers.net/go.cms/document"
)

type MarkdownFormatter int

var Markdown = MarkdownFormatter(0)

// BUG: This code does not currently escape textual content.
func (MarkdownFormatter) Format(element Element) string {
	switch el := element.(type) {
	case *Document:
		return html.EscapeString(el.Title) + "\n==========\n\n" + Markdown.formatContent(el.Contents())
	case *Italic:
		return "_" + Markdown.formatContent(el.Contents()) + "_"
	case *Bold:
		return "**" + Markdown.formatContent(el.Contents()) + "**"
	case *LineBreak:
		return "  \n"
	case *Paragraph:
		return Markdown.formatContent(el.Contents()) + "\n\n"
	case *Image:
		return "![" + html.EscapeString(el.Description) + "](" + html.EscapeString(el.URL) + ")"
	case *Link:
		return "[" + Markdown.formatContent(el.Contents()) + "](" + html.EscapeString(el.URL) + ")"
	case *LeafElement:
		return html.EscapeString(el.Text)
	}
	return fmt.Sprintf("%#v", element)
}

func (MarkdownFormatter) formatContent(content Content) string {
	s := ""
	for _, element := range content {
		s += Markdown.Format(element)
		if _, isInline := element.(InlineElement); isInline {
			s += " "
		} else {
			s += "\n"
		}
	}
	return s[:len(s)-1]
}
