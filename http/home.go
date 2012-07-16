package http

import (
	. "llamaslayers.net/go.cms/document"
	"llamaslayers.net/go.cms/storage"
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
		storage.GetLatestArticles(),
	}
}
