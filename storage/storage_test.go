package storage

import (
	"llamaslayers.net/go.cms/document"
	"llamaslayers.net/go.cms/formatter"
	"testing"
	"time"
)

func TestConvertUint32(t *testing.T) {
	t.Parallel()

	i := uint32(1)
	buf, err := FromUint32(make([]byte, 0), i)
	if err != nil {
		t.Error(err)
	}
	j, buf, err := ToUint32(buf)
	if err != nil {
		t.Error(err)
	}
	if len(buf) != 0 {
		t.Error("Leftover buffer size is nonzero: ", len(buf))
	}
	if i != j {
		t.Error("expected(", i, ") != result(", j, ")")
	}

	i = uint32(0x12345678)
	buf, err = FromUint32(make([]byte, 0), i)
	if err != nil {
		t.Error(err)
	}
	j, buf, err = ToUint32(buf)
	if err != nil {
		t.Error(err)
	}
	if len(buf) != 0 {
		t.Error("Leftover buffer size is nonzero: ", len(buf))
	}
}

func TestConvertString(t *testing.T) {
	t.Parallel()

	s1 := "Test string OF SCIENCE"
	buf, err := FromString(make([]byte, 0), s1)
	if err != nil {
		t.Error(err)
	}
	s2, buf, err := ToString(buf)
	if err != nil {
		t.Error(err)
	}
	if len(buf) != 0 {
		t.Error("Leftover buffer size is nonzero: ", len(buf))
	}
	if s1 != s2 {
		t.Error("expected(", s1, ") != result(", s2, ")")
	}
}

func TestConvertTime(t *testing.T) {
	t.Parallel()

	t1 := time.Now()
	buf, err := FromTime(make([]byte, 0), t1)
	if err != nil {
		t.Error(err)
	}
	t2, buf, err := ToTime(buf)
	if err != nil {
		t.Error(err)
	}
	if len(buf) != 0 {
		t.Error("Leftover buffer size is nonzero: ", len(buf))
	}
	if t1 != t2 {
		t.Error("expected(", t1, ") != result(", t2, ")")
	}
}

func TestConvertDocument(t *testing.T) {
	t.Parallel()

	doc1 := &document.Document{
		"Test document of SCIENCE",
		document.Content{
			&document.Paragraph{
				document.Content{
					&document.LeafElement{"SCIENCE IS"},
					&document.Bold{document.Content{&document.LeafElement{"COOL!"}}},
				},
			},
			&document.Article{
				"SCIENCE",
				time.Time{},
				document.Content{
					&document.LeafElement{"I'm"},
					&document.Italic{document.Content{&document.LeafElement{"OLD"}}},
				},
			},
		},
	}
	buf, err := FromElement(make([]byte, 0), doc1)
	if err != nil {
		t.Error(err)
	}
	doc2, buf, err := ToElement(buf)
	if err != nil {
		t.Error(err)
	}
	if len(buf) != 0 {
		t.Error("Leftover buffer size is nonzero: ", len(buf))
	}
	// We can't compare the documents directly because they contain pointers.
	d1, d2 := formatter.HTML.Format(doc1), formatter.HTML.Format(doc2)
	if d1 != d2 {
		t.Error("expected(", d1, ") != result(", d2, ")")
	}
}
