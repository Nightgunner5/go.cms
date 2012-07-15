package formatter

import . "llamaslayers.net/go.cms/document"

// A Formatter, quite simply, takes an Element and formats it so that a third-party application
// (browser, TeX compiler, etc.) can display it in a human-readable format.
type Formatter interface {
	Format(Element) string
}
