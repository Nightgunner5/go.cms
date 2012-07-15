package formatter

import (
	"fmt"
	. "llamaslayers.net/go.cms/document"
	"testing"
)

func ExampleLaTeXFormatter_Format() {
	doc := &Document{
		"Test page",
		Content{
			&Paragraph{
				Content{
					&LeafElement{"This"},
					&Italic{Content{&LeafElement{"is"}}},
					&Bold{Content{&LeafElement{"a"},
						&Italic{Content{&LeafElement{"test."}}}}},
				},
			},
		},
	}
	fmt.Println(LaTeX.Format(doc))

	// Output:
	// \documentclass{article}
	// \usepackage{graphicx}
	// \usepackage{hyperref}
	// \title{Test page}
	// \begin{document}
	// This \textit{is} \textbf{a \textit{test.}} \\ \\
	// \end{document}
}

func BenchmarkLaTeX(b *testing.B) {
	b.StopTimer()
	doc := &Document{
		"Test page",
		Content{
			&Paragraph{
				Content{
					&LeafElement{"This"},
					&Italic{Content{&LeafElement{"is"}}},
					&Bold{Content{&LeafElement{"a"},
						&Italic{Content{&LeafElement{"test."}}}}},
				},
			},
		},
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		LaTeX.Format(doc)
	}
}
