package marky

import (
	"fmt"
	"strings"
)

type Heading struct {
	Caption  string
	Level    int
	Altenate bool
}

func (h Heading) String() string {
	return fmt.Sprintf("%s %s\n\n", strings.Repeat("#", h.Level), h.Caption)
}
