package marky

import "strings"

type Quote struct {
	BlockElement
}

func (b *Quote) Add(e Element) *Quote {
	b.elements = append(b.elements, e)
	return b
}

func (q Quote) String() string {
	var s strings.Builder
	s.WriteRune('\n')

	for _, e := range q.elements {
		for _, v := range strings.Split(e.String(), "\n") {
			s.WriteRune('>')
			if !strings.HasPrefix(v, ">") {
				s.WriteRune(' ')
			}

			s.WriteString(v)
			s.WriteRune('\n')
		}
	}

	s.WriteRune('\n')
	return s.String()
}
