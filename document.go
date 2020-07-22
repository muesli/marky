package marky

import "strings"

type Document struct {
	BlockElement
}

func NewDocument() *Document {
	return &Document{}
}

func (d *Document) Add(e Element) *Document {
	d.elements = append(d.elements, e)
	return d
}

func (d Document) String() string {
	var s strings.Builder

	v := d.BlockElement.String()
	s.WriteString(strings.TrimSpace(v))
	s.WriteRune('\n')

	return s.String()
}
