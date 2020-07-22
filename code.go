package marky

import "strings"

type Code struct {
	Source   string
	Language string
	Fenced   bool
}

func (c Code) String() string {
	var s strings.Builder
	s.WriteRune('\n')

	if c.Fenced || c.Language != "" {
		s.WriteString("```")
		s.WriteString(c.Language)
		s.WriteRune('\n')
		s.WriteString(c.Source)
		s.WriteString("```")
		s.WriteRune('\n')
	} else {
		for _, v := range strings.Split(c.Source, "\n") {
			if len(v) > 0 {
				s.WriteString("    ")
			}
			s.WriteString(v)
			s.WriteRune('\n')
		}
	}

	return s.String()
}
