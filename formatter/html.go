package formatter

import (
	"fmt"
	. "github.com/Nightgunner5/go.cms/document"
	"html"
)

type HTMLFormatter int

var HTML = HTMLFormatter(0)

func (HTMLFormatter) Format(element Element) string {
	switch el := element.(type) {
	case *Document:
		return `<!DOCTYPE html>
<html>
<head>
` + HTML.documentHeader(el) + `
</head>
<body>
` + HTML.formatContent(el.Contents()) + `
</body>
</html>`
	case *Italic:
		return "<em>" + HTML.formatContent(el.Contents()) + "</em>"
	case *Bold:
		return "<strong>" + HTML.formatContent(el.Contents()) + "</strong>"
	case *Underline:
		return "<u>" + HTML.formatContent(el.Contents()) + "</u>"
	case *LineBreak:
		return "<br>"
	case *Paragraph:
		return "<p>" + HTML.formatContent(el.Contents()) + "</p>"
	case *LeafElement:
		return html.EscapeString(el.Text)
	}
	return fmt.Sprintf("%#v", element)
}

func (HTMLFormatter) documentHeader(doc *Document) string {
	return "<title>" + html.EscapeString(doc.Title) + "</title>"
}

func (HTMLFormatter) formatContent(content Content) string {
	s := ""
	for _, element := range content {
		s += HTML.Format(element)
		if _, isInline := element.(InlineElement); isInline {
			s += " "
		} else {
			s += "\n"
		}
	}
	return s[:len(s)-1]
}
