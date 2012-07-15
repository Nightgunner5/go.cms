package main

import "flag"

var flagReadme = flag.Bool("readme", false, "Prints a README and exits.")
var flagReadmeFormat = flag.String("readme-format", "Markdown", "Other formats available are LaTeX and HTML.")

func main() {
	flag.Parse()

	if *flagReadme {
		readme()
		return
	}
}
