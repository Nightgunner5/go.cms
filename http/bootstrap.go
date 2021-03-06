package http

import "net/http"

func init() {
	http.Handle("/css/bootstrap.css", MakeStaticPage("text/css; charset=UTF-8", BootstrapCSS))
	http.Handle("/js/bootstrap.js", MakeStaticPage("application/javascript; charset=UTF-8", BootstrapJS))
	http.Handle("/img/glyphicons-halflings.png", MakeStaticPageNoCompress("image/png", GlyphIcons))
	http.Handle("/img/glyphicons-halflings-white.png", MakeStaticPageNoCompress("image/png", GlyphIconsWhite))
}
