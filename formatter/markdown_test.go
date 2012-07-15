package formatter

import (
	"fmt"
	. "llamaslayers.net/go.cms/document"
	"testing"
)

func ExampleMarkdownFormatter_Format() {
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
	fmt.Println(Markdown.Format(doc))

	// Output:
	// Test page
	// ==========
	// 
	// This _is_ **a _test._**
}

func BenchmarkMarkdown(b *testing.B) {
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
		Markdown.Format(doc)
	}
}
