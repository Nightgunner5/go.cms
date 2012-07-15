package formatter

import (
	"fmt"
	. "github.com/Nightgunner5/go.cms/document"
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
	// <title>Test page</title>
	// </head>
	// <body>
	// <p>This <em>is</em> <strong>a <em>test.</em></strong></p>
	// </body>
	// </html>
}
