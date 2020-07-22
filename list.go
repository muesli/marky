package marky

import (
	"fmt"
	"strings"
)

type List struct {
	BlockElement

	Ordered bool
}

func NewList() *List {
	return &List{}
}

func (l *List) Add(e Element) *List {
	l.elements = append(l.elements, e)
	return l
}

func (l List) String() string {
	var s strings.Builder
	i := 0

	s.WriteRune('\n')
	for _, e := range l.elements {
		i++

		if l.Ordered {
			s.WriteString(fmt.Sprintf("%d. ", i))
		} else {
			s.WriteString("- ")
		}
		s.WriteString(e.String())
		s.WriteRune('\n')
	}
	s.WriteRune('\n')

	return s.String()
}
