package formatter

import . "llamaslayers.net/go.cms/document"

type LaTeXFormatter int

var LaTeX Formatter = LaTeXFormatter(0)

func (LaTeXFormatter) Format(doc Element) string {
	// TODO
	return ""
}
