// func From...([]byte, sometype) ([]byte, error)
//     first parameter is the current byte slice that should be appended to.
//     second parameter is the type to be converted
//     first return value is the byte slice with the data appended to it
//     second return value is error
// 
// func To...([]byte) (sometype, []byte, error)
//     parameter is a byte slice to parse the front of
//     first return value is the type as parsed from the byte slice
//     second return value is the remaining byte slice
//     third return value is error
package storage

import (
	"fmt"
	"llamaslayers.net/go.cms/document"
	"time"
)

type ElementType uint8

const (
	EndOfContent ElementType = iota
	Document
	Article
	Paragraph
	LineBreak
	Italic
	Bold
	Image
	Link
	LeafElement
)

func FromElement(buf []byte, element document.Element) ([]byte, error) {
	var err error
	switch el := element.(type) {
	case *document.Document:
		buf = append(buf, byte(Document))
		buf, err = FromString(buf, el.Title)
		if err != nil {
			return nil, err
		}
	case *document.Article:
		buf = append(buf, byte(Article))
		buf, err = FromString(buf, el.Title)
		if err != nil {
			return nil, err
		}
		buf, err = FromTime(buf, el.Timestamp)
		if err != nil {
			return nil, err
		}
	case *document.Paragraph:
		buf = append(buf, byte(Paragraph))
	case *document.LineBreak:
		buf = append(buf, byte(LineBreak))
	case *document.Italic:
		buf = append(buf, byte(Italic))
	case *document.Bold:
		buf = append(buf, byte(Bold))
	case *document.Image:
		buf = append(buf, byte(Image))
		buf, err = FromString(buf, el.URL)
		if err != nil {
			return nil, err
		}
		buf, err = FromString(buf, el.Description)
		if err != nil {
			return nil, err
		}
	case *document.Link:
		buf = append(buf, byte(Link))
		buf, err = FromString(buf, el.URL)
		if err != nil {
			return nil, err
		}
	case *document.LeafElement:
		buf = append(buf, byte(LeafElement))
		buf, err = FromString(buf, el.Text)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unknown element type for %#v", element)
	}
	if con, ok := element.(document.Container); ok {
		for _, el := range con.Contents() {
			buf, err = FromElement(buf, el)
			if err != nil {
				return nil, err
			}
		}
		buf = append(buf, byte(EndOfContent))
	}
	return buf, nil
}

func ToElement(buf []byte) (document.Element, []byte, error) {
	var element document.Element

	elementType := ElementType(buf[0])
	buf = buf[1:]
	var err error

	switch elementType {
	case Document:
		var title string
		title, buf, err = ToString(buf)
		if err != nil {
			return nil, nil, err
		}
		element = &document.Document{title, document.Content{}}
	case Article:
		var title string
		title, buf, err = ToString(buf)
		if err != nil {
			return nil, nil, err
		}
		var timestamp time.Time
		timestamp, buf, err = ToTime(buf)
		if err != nil {
			return nil, nil, err
		}
		element = &document.Article{title, timestamp, document.Content{}}
	case Paragraph:
		element = &document.Paragraph{document.Content{}}
	case LineBreak:
		element = &document.LineBreak{}
	case Italic:
		element = &document.Italic{document.Content{}}
	case Bold:
		element = &document.Bold{document.Content{}}
	case Image:
		var url string
		url, buf, err = ToString(buf)
		if err != nil {
			return nil, nil, err
		}
		var description string
		description, buf, err = ToString(buf)
		if err != nil {
			return nil, nil, err
		}
		element = &document.Image{url, description}
	case Link:
		var url string
		url, buf, err = ToString(buf)
		if err != nil {
			return nil, nil, err
		}
		element = &document.Link{url, document.Content{}}
	case LeafElement:
		var text string
		text, buf, err = ToString(buf)
		if err != nil {
			return nil, nil, err
		}
		element = &document.LeafElement{text}
	default:
		return nil, nil, fmt.Errorf("Unknown element type: %d", elementType)
	}

	if con, ok := element.(document.Container); ok {
		for ElementType(buf[0]) != EndOfContent {
			var el document.Element
			el, buf, err = ToElement(buf)
			if err != nil {
				return nil, nil, err
			}
			con.AddChild(el)
		}
		buf = buf[1:]
	}
	return element, buf, nil
}

// Little endian
func FromUint32(buf []byte, in uint32) ([]byte, error) {
	for i := 0; i < 4; i++ {
		buf = append(buf, byte(in&0xff))
		in >>= 8
	}
	return buf, nil
}

// Little endian
func ToUint32(in []byte) (uint32, []byte, error) {
	if len(in) /* MY FRIEND! I COME! */ < 4 {
		return 0, nil, fmt.Errorf("A 32-bit integer requires 4 bytes of data space, but the input only has %d bytes.", len(in))
	}
	out := uint32(in[0])
	out |= uint32(in[1]) << 8
	out |= uint32(in[2]) << 16
	out |= uint32(in[3]) << 24
	return out, in[4:], nil
}

func FromString(buf []byte, in string) ([]byte, error) {
	buf, err := FromUint32(buf, uint32(len(in)))
	if err != nil {
		return nil, err
	}
	buf = append(buf, []byte(in)...)
	return buf, nil
}

func ToString(in []byte) (string, []byte, error) {
	l, buf, err := ToUint32(in)
	if err != nil {
		return "", nil, err
	}
	if uint32(len(buf)) < l {
		return "", nil, fmt.Errorf("String length is %d, but buffer is only %d bytes long!", l, len(buf))
	}
	return string(buf[0:l]), buf[l:], nil
}

func FromTime(buf []byte, in time.Time) ([]byte, error) {
	encoded, err := in.GobEncode()
	if err != nil {
		return nil, err
	}
	buf = append(buf, byte(len(encoded)))
	buf = append(buf, encoded...)
	return buf, nil
}

func ToTime(in []byte) (time.Time, []byte, error) {
	var t time.Time
	l := in[0]
	err := t.GobDecode(in[1 : l+1])
	if err != nil {
		return t, nil, err
	}
	return t, in[l+1:], nil
}
