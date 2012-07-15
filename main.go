package main

import (
	"flag"
	"fmt"
	. "llamaslayers.net/go.cms/document"
	. "llamaslayers.net/go.cms/formatter"
)

var flagReadme = flag.Bool("readme", false, "Prints a README and exits.")
var flagReadmeFormat = flag.String("readme-format", "Markdown", "Other formats available are LaTeX and HTML.")

func main() {
	flag.Parse()

	if *flagReadme {
		readme := &Document{
			"go.cms Readme",
			Content{
				&Link{
					"http://travis-ci.org/Nightgunner5/go.cms",
					Content{&Image{"https://secure.travis-ci.org/Nightgunner5/go.cms.png?branch=master", "Build Status"}},
				},
			},
		}

		switch *flagReadmeFormat {
		case "Markdown":
			fmt.Println(Markdown.Format(readme))
		case "LaTeX":
			fmt.Println(LaTeX.Format(readme))
		case "HTML":
			fmt.Println(HTML.Format(readme))
		default:
			fmt.Println("Unknown readme-format value. Capitalization matters!")
		}
		return
	}
}
