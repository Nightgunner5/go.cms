package main

import (
	"flag"
	"llamaslayers.net/go.cms/http"
)

var flagReadme = flag.Bool("readme", false, "Prints a README and exits.")
var flagReadmeFormat = flag.String("readme-format", "Markdown", "Other formats available are LaTeX and HTML.")
var flagHttpAddr = flag.String("http", ":8080", "The host (optional) and port to listen on for HTTP connections.")

func main() {
	flag.Parse()

	if *flagReadme {
		readme()
		return
	}

	http.Startup(*flagHttpAddr)
}
