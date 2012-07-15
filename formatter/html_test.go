package formatter

import (
	"fmt"
	. "llamaslayers.net/go.cms/document"
	"testing"
)

func ExampleHTMLFormatter_Format() {
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
	fmt.Println(HTML.Format(doc))

	// Output:
	// <!DOCTYPE html>
	// <html>
	// <head>
	// <meta charset="utf-8">
	// <title>Test page</title>
	// <link href="/css/bootstrap.css" rel="stylesheet">
	// </head>
	// <body>
	// <div class="container">
	// <h1 class="page-header">Test page</h1>
	// <p>This <em>is</em> <strong>a <em>test.</em></strong></p>
	// </div>
	// <script src="/js/bootstrap.js" async></script>
	// </body>
	// </html>
}

func BenchmarkHTML(b *testing.B) {
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
		HTML.Format(doc)
	}
}
