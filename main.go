package main

import (
	"flag"
	"llamaslayers.net/go.cms/http"
	"llamaslayers.net/go.cms/storage"
)

var flagReadme = flag.Bool("readme", false, "Prints a README and exits.")
var flagReadmeFormat = flag.String("readme-format", "Markdown", "Other formats available are LaTeX and HTML.")

var flagHttpAddr = flag.String("http", ":8080", "The host (optional) and port to listen on for HTTP connections.")
var flagNoSpdy = flag.Bool("nospdy", false, "Disables SPDY and TLS.")
var flagFakeLag = flag.Int64("fakelag", 0, "If fakelag is positive, each request will wait fakelag milliseconds before responding.")

var flagDBLocation = flag.String("db", "./gocms.db", "Location for data storage (sqlite3 database)")

func main() {
	flag.Parse()

	if *flagReadme {
		readme()
		return
	}

	storage.Startup(*flagDBLocation)
	http.Startup(*flagHttpAddr, *flagFakeLag, *flagNoSpdy)
}
