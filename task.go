package marky

import "strings"

type Task struct {
	Text string
	Done bool
}

func (t Task) String() string {
	var s strings.Builder

	s.WriteRune('[')
	if t.Done {
		s.WriteRune('x')
	} else {
		s.WriteRune(' ')
	}
	s.WriteString("] ")

	s.WriteString(t.Text)

	return s.String()
}
