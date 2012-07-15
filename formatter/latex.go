package formatter

import (
	"fmt"
	. "llamaslayers.net/go.cms/document"
)

type LaTeXFormatter int

var LaTeX = LaTeXFormatter(0)

// BUG: This code does not currently escape textual content.
func (LaTeXFormatter) Format(element Element) string {
	switch el := element.(type) {
	case *Document:
		return `\documentclass{article}
\usepackage{graphicx}
\usepackage{hyperref}
\title{` + el.Title + `}
\begin{document}
` + LaTeX.formatContent(el.Contents()) + `
\end{document}`
	case *Italic:
		return "\\textit{" + LaTeX.formatContent(el.Contents()) + "}"
	case *Bold:
		return "\\textbf{" + LaTeX.formatContent(el.Contents()) + "}"
	case *LineBreak:
		return " \\\\ "
	case *Paragraph:
		return LaTeX.formatContent(el.Contents()) + " \\\\ \\\\"
	case *Image:
		// This isn't strictly valid LaTeX formatting, but there's no way in LaTeX to specify an external image.
		return `\begin{figure}[h]
\centering
\includegraphics{` + el.URL + `}
\caption{` + el.Description + `}
\end{figure}`
	case *Link:
		return "\\href{" + el.URL + "}{" + LaTeX.formatContent(el.Contents()) + "}"
	case *LeafElement:
		// TODO: escaping
		return el.Text
	}
	return fmt.Sprintf("%#v", element)
}

func (LaTeXFormatter) formatContent(content Content) string {
	s := ""
	for _, element := range content {
		s += LaTeX.Format(element)
		if _, isInline := element.(InlineElement); isInline {
			s += " "
		} else {
			s += "\n"
		}
	}
	return s[:len(s)-1]
}
