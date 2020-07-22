package marky

import "strings"

type Image struct {
	Text  string
	Title string
	Image string
	URL   string
}

func (i Image) String() string {
	var s strings.Builder

	if i.URL != "" {
		s.WriteRune('[')
	}

	s.WriteString("![")
	s.WriteString(i.Text)
	s.WriteRune(']')

	s.WriteRune('(')
	s.WriteString(i.Image)
	if i.Title != "" {
		s.WriteString(" \"")
		s.WriteString(i.Title)
		s.WriteRune('"')
	}
	s.WriteRune(')')

	if i.URL != "" {
		s.WriteString("](")
		s.WriteString(i.URL)
		s.WriteRune(')')
	}

	return s.String()
}
