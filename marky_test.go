package marky

import (
	"fmt"
	"testing"
)

func TestMarky(t *testing.T) {
	d := NewDocument().
		Add(Heading{
			Caption: "Here's a Heading",
			Level:   1,
		}).
		Add(NewParagraph().
			Add(Text{
				Text: "I really like using Markdown.",
			}),
		).
		Add(NewParagraph().
			Add(Image{
				Text:  "An Image",
				Image: "/image.png",
			}),
		)

	exp := `# Here's a Heading

I really like using Markdown.

![An Image](/image.png)
`

	if d.String() != exp {
		t.Error(diff(exp, d.String()))
	}
}

func TestHeading(t *testing.T) {
	e := Heading{
		Caption: "Here's a Heading",
		Level:   1,
	}
	exp := "# Here's a Heading\n\n"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e.Level = 3
	exp = "### Here's a Heading\n\n"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}
}

func TestText(t *testing.T) {
	e := Text{
		Text: "I just love text",
	}
	exp := "I just love text"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e.Bold = true
	exp = "**I just love text**"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e.Bold = false
	e.Italic = true
	exp = "*I just love text*"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e.Bold = true
	e.Italic = true
	exp = "***I just love text***"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}
}

func TestLink(t *testing.T) {
	e := Link{
		Text: "A Link",
		URL:  "https://a.url",
	}
	exp := "[A Link](https://a.url)"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e.Title = "A Title"
	exp = "[A Link](https://a.url \"A Title\")"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}
}

func TestImage(t *testing.T) {
	e := Image{
		Text:  "An Image",
		Image: "/image.png",
	}
	exp := "![An Image](/image.png)"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e.Title = "A Title"
	exp = "![An Image](/image.png \"A Title\")"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e.URL = "https://a.url"
	exp = "[![An Image](/image.png \"A Title\")](https://a.url)"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}
}

func TestQuote(t *testing.T) {
	e := Quote{}
	e.Add(Text{
		Text: "Dorothy followed her through many of the beautiful rooms in her castle.",
	})
	exp := "\n> Dorothy followed her through many of the beautiful rooms in her castle.\n\n"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e = Quote{}
	e.Add(Text{
		Text: "> The Witch bade her clean the pots and kettles and sweep the floor and keep the fire fed with wood.",
	})
	exp = "\n>> The Witch bade her clean the pots and kettles and sweep the floor and keep the fire fed with wood.\n\n"
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e = Quote{}
	e.Add(Text{
		Text: "First line of quote",
	}).Add(Text{
		Text: "Second line of quote",
	})

	exp = `
> First line of quote
> Second line of quote

`

	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}
}

func TestCode(t *testing.T) {
	e := Code{
		Source: `echo "Hello"
echo "World"
`,
		Language: "bash",
	}

	exp := `
` + "```bash" + `
echo "Hello"
echo "World"
` + "```" + `
`
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e.Language = ""
	e.Fenced = true
	exp = `
` + "```" + `
echo "Hello"
echo "World"
` + "```" + `
`
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e.Fenced = false
	exp = `
    echo "Hello"
    echo "World"

`
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}
}

func TestList(t *testing.T) {
	e := List{
		Ordered: true,
	}
	e.Add(Text{
		Text: "First Item",
	}).Add(Text{
		Text: "Second Item",
	})

	exp := `
1. First Item
2. Second Item

`

	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}

	e.Ordered = false
	exp = `
- First Item
- Second Item

`

	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}
}

func TestTasks(t *testing.T) {
	e := List{}
	e.Add(Task{
		Text: "First Task",
	}).Add(Task{
		Text: "Second Task",
		Done: true,
	})

	exp := `
- [ ] First Task
- [x] Second Task

`
	if e.String() != exp {
		t.Error(diff(exp, e.String()))
	}
}

func diff(s1, s2 string) string {
	return fmt.Sprintf("Expected output:\n%s\n\nGot:\n%s", s1, s2)
}
