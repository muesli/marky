package marky

import "strings"

type Link struct {
	Text  string
	Title string
	URL   string
}

func (l Link) String() string {
	var s strings.Builder

	s.WriteRune('[')
	s.WriteString(l.Text)
	s.WriteRune(']')

	s.WriteRune('(')
	s.WriteString(l.URL)
	if l.Title != "" {
		s.WriteString(" \"")
		s.WriteString(l.Title)
		s.WriteRune('"')
	}
	s.WriteRune(')')

	return s.String()
}
