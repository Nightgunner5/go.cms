package http

import (
	. "llamaslayers.net/go.cms/document"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		notFoundHandler(w, req)
		return
	}
	DoPage(w, req, HomeDocument(), http.StatusOK)
}

func HomeDocument() *Document {
	return &Document{
		"Home",
		Content{
			&Paragraph{
				Content{
					&LeafElement{"Test"},
				},
			},
		},
	}
}
