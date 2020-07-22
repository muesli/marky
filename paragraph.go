package marky

import "strings"

type Paragraph struct {
	BlockElement
}

func NewParagraph() *Paragraph {
	return &Paragraph{}
}

func (p *Paragraph) Add(e Element) *Paragraph {
	p.elements = append(p.elements, e)
	return p
}

func (p Paragraph) String() string {
	var s strings.Builder

	s.WriteString(p.BlockElement.String())
	s.WriteRune('\n')

	return s.String()
}
