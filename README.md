# marky

Generate markdown programmatically

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/marky)
[![Build Status](https://github.com/muesli/marky/workflows/build/badge.svg)](https://github.com/muesli/marky/actions)
[![Coverage Status](https://coveralls.io/repos/github/muesli/marky/badge.svg?branch=master)](https://coveralls.io/github/muesli/marky?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/marky)](http://goreportcard.com/report/muesli/marky)

## Usage

```go
d := NewDocument().
    Add(Heading{
        Caption: "This is marky",
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

ioutil.WriteFile("readme.md", []byte(d.String()), 0644)
```

### Links

```go
d.Add(Link{
    Text: "A Link",
    URL:  "https://a.url",
})
```

### Images

```go
d.Add(Image{
    Text:  "An Image",
    Image: "/image.png",
    URL: "https://a.url",
})
```

### Quotes

```go
e := Quote{}
e.Add(Text{
    Text: "Dorothy followed her through many of the beautiful rooms in her castle.",
})
d.Add(e)
```

### Codeblocks

```go
d.Add(Code{
    Source: `echo "Hello"\necho "World"`,
    Language: "bash",
})
```

### Lists

```go
e := List{
    Ordered: true,
}
e.Add(Text{
    Text: "First Item",
}).Add(Text{
    Text: "Second Item",
})
d.Add(e)
```

### Tasks

```go
e := List{}
e.Add(Task{
    Text: "First Task",
}).Add(Task{
    Text: "Second Task",
    Done: true,
})
d.Add(e)
```
