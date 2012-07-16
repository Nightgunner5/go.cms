package http

import (
	"bytes"
	"compress/gzip"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	spdy "github.com/jmckaskill/gospdy"
	. "llamaslayers.net/go.cms/document"
	"llamaslayers.net/go.cms/formatter"
	"log"
	"net/http"
	"strings"
	"time"
)

var pauseBeforeResponding time.Duration

func requiredHeaders(w http.ResponseWriter) {
	if pauseBeforeResponding > 0 {
		time.Sleep(pauseBeforeResponding)
	}
	w.Header().Set("Server", "go get llamaslayers.net/go.cms")
}

func Startup(addr string, fakelag int64, nospdy bool) {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/favicon.ico", noContentHandler)

	pauseBeforeResponding = time.Duration(fakelag) * time.Millisecond

	log.Print("Now listening on ", addr)

	if nospdy {
		log.Fatal(http.ListenAndServe(addr, nil))
	} else {
		log.Fatal(spdy.ListenAndServeTLS(addr, "cert.pem", "key.pem", nil))
	}
}

type static_page struct {
	contentType, etag        string
	uncompressed, compressed []byte
}

func _etag(in []byte) string {
	hash := sha1.New()
	// Error can be safely ignored here as it is never set in sha1's code.
	hash.Write(in)
	var output [0]byte
	return "\"" + hex.EncodeToString(hash.Sum(output[:])) + "\""
}

func _compress(in []byte, level int) []byte {
	var b bytes.Buffer
	// Error can be safely ignored here as it is only set for invalid compression levels.
	w, _ := gzip.NewWriterLevel(&b, level)
	// Error can be safely ignored here as it would only be set if the compressor
	// had a serious bug in it and passed the wrong argument.
	w.Write(in)
	w.Close()
	return b.Bytes()
}

// Static pages are compressed and cached as aggressively as possible.
// Expiration dates are 2 weeks from the time of access.
func MakeStaticPage(contentType string, content []byte) http.Handler {
	return &static_page{contentType, _etag(content), content, _compress(content, gzip.BestCompression)}
}

// Use this for content like images, which are already compressed.
func MakeStaticPageNoCompress(contentType string, content []byte) http.Handler {
	return &static_page{contentType, _etag(content), content, nil}
}

func (page *static_page) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	requiredHeaders(w)
	if req.Header.Get("If-None-Match") == page.etag {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("Content-Type", page.contentType)
	w.Header().Set("ETag", page.etag)
	w.Header().Set("Expires", time.Now().AddDate(0, 0, 14).Format(time.RFC1123))
	if page.compressed != nil {
		w.Header().Set("Vary", "Accept-Encoding")
	}
	if page.compressed != nil && strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Length", fmt.Sprint(len(page.compressed)))
		w.Header().Set("Content-Encoding", "gzip")
		w.Write(page.compressed)
	} else {
		w.Header().Set("Content-Length", fmt.Sprint(len(page.uncompressed)))
		w.Write(page.uncompressed)
	}
}

func DoPage(w http.ResponseWriter, req *http.Request, doc *Document, status int) {
	requiredHeaders(w)
	buf := []byte(formatter.HTML.Format(doc))
	etag := _etag(buf)
	if req.Header.Get("If-None-Match") == etag {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Set("Vary", "Accept-Encoding")
	w.Header().Set("ETag", etag)
	if len(doc.Contents()) > 0 {
		if article, ok := doc.Contents()[0].(*Article); ok {
			w.Header().Set("Last-Modified", article.Timestamp.Format(time.RFC1123))
		}
	}
	w.Header().Set("Expires", time.Now().Add(time.Hour).Format(time.RFC1123))

	if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
		buf := _compress(buf, gzip.DefaultCompression)
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Length", fmt.Sprint(len(buf)))
		w.WriteHeader(status)
		gz := gzip.NewWriter(w)
		gz.Write([]byte(formatter.HTML.Format(doc)))
		gz.Close()
	} else {
		w.Header().Set("Content-Length", fmt.Sprint(len(buf)))
		w.WriteHeader(status)
		w.Write(buf)
	}
}

func noContentHandler(w http.ResponseWriter, req *http.Request) {
	requiredHeaders(w)
	w.WriteHeader(http.StatusNoContent)
}

var notFoundDocument = &Document{
	"Four, oh Four!",
	Content{
		&Paragraph{Content{&LeafElement{"I'm very sorry, but I couldn't find the page you wanted."}}},
		&Paragraph{Content{&LeafElement{"Either it ran away or you're delusional. I'm leaning toward the latter."}}},
	},
}

func notFoundHandler(w http.ResponseWriter, req *http.Request) {
	DoPage(w, req, notFoundDocument, http.StatusNotFound)
}
