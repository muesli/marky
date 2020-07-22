package marky

import "strings"

// BlockElement is a special kind of Element, which can contain other Elements.
type BlockElement struct {
	elements []Element
}

func (b *BlockElement) String() string {
	var s strings.Builder

	for _, e := range b.elements {
		s.WriteString(e.String())
	}
	s.WriteRune('\n')

	return s.String()
}
