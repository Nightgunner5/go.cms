package formatter

import (
	"fmt"
	"html"
	. "llamaslayers.net/go.cms/document"
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
<div class="container">
<h1 class="page-header">` + html.EscapeString(el.Title) + `</h1>
` + HTML.formatContent(el.Contents()) + `
</div>
<script src="/js/bootstrap.js" async></script>
</body>
</html>`
	case *Italic:
		return "<em>" + HTML.formatContent(el.Contents()) + "</em>"
	case *Bold:
		return "<strong>" + HTML.formatContent(el.Contents()) + "</strong>"
	case *LineBreak:
		return "<br>"
	case *Paragraph:
		return "<p>" + HTML.formatContent(el.Contents()) + "</p>"
	case *Image:
		return "<img src=\"" + html.EscapeString(el.URL) + "\" alt=\"" + html.EscapeString(el.Description) + "\" title=\"" + html.EscapeString(el.Description) + "\">"
	case *Link:
		return "<a href=\"" + html.EscapeString(el.URL) + "\">" + HTML.formatContent(el.Contents()) + "</a>"
	case *LeafElement:
		return html.EscapeString(el.Text)
	}
	return fmt.Sprintf("%#v", element)
}

func (HTMLFormatter) documentHeader(doc *Document) string {
	return `<meta charset="utf-8">
<title>` + html.EscapeString(doc.Title) + `</title>
<link href="/css/bootstrap.css" rel="stylesheet">`
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
