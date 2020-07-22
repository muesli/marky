package marky

import "strings"

type Text struct {
	Text string

	Code   bool
	Bold   bool
	Italic bool
}

func (t Text) String() string {
	var s strings.Builder

	if t.Code {
		s.WriteString("`")
	}
	if t.Bold {
		s.WriteString("**")
	}
	if t.Italic {
		s.WriteString("*")
	}

	s.WriteString(t.Text)

	if t.Italic {
		s.WriteString("*")
	}
	if t.Bold {
		s.WriteString("**")
	}
	if t.Code {
		s.WriteString("`")
	}

	return s.String()
}
