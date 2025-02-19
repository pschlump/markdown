package parser

import (
	"bytes"
	"testing"

	"github.com/pschlump/dbgo"
	"github.com/pschlump/markdown/ast"
)

func TestSanitizedAnchorName(t *testing.T) {
	// icky! xyzzy - TODO - PJS - change this.
	tests := []string{
		"This is a header",
		"this-is-a-header",

		"This is also          a header",
		"this-is-also-a-header",

		"main.go",
		"main-go",

		"Article 123",
		"article-123",

		"<- Let's try this, shall we?",
		"let-s-try-this-shall-we",

		"        ",
		"empty",

		"Hello, 世界",
		"hello-世界",

		"世界",
		"世界",

		"⌥",
		"empty",
	}
	n := len(tests)
	for i := 0; i < n; i += 2 {
		text := tests[i]
		want := tests[i+1]
		if got := sanitizeHeadingID(text); got != want {
			t.Errorf("SanitizedAnchorName(%q):\ngot %q\nwant %q", text, got, want)
		}
	}
}

// func (p *Parser) Parse(input []byte) ast.Node {

func TestParser_01(t *testing.T) {
	input := "| a | b |\n| - | - |\n|	foo | bar |\n"
	p := NewWithExtensions(CommonExtensions)
	doc := p.Parse([]byte(input))
	var buf bytes.Buffer
	ast.Print(&buf, doc)
	got := buf.String()
	exp := "Table\n  TableHeader\n    TableRow\n      TableCell\n        Text 'a'\n      TableCell\n        Text 'b'\n  TableBody\n    TableRow\n      TableCell\n        Text '\\tfoo'\n      TableCell\n        Text 'bar'\n"
	if got != exp {
		t.Errorf("\nInput   [%#v]\nExpected[%#v]\nGot     [%#v]\n", input, exp, got)
	}
}

func TestParser_02(t *testing.T) {
	/*
		input := "| a | b |\n| - | - |\n|	foo | bar |\n"
		p := NewWithExtensions(CommonExtensions)
		doc := p.Parse([]byte(input))
		var buf bytes.Buffer
		ast.Print(&buf, doc)
		got := buf.String()
		exp := "Table\n  TableHeader\n    TableRow\n      TableCell\n        Text 'a'\n      TableCell\n        Text 'b'\n  TableBody\n    TableRow\n      TableCell\n        Text '\\tfoo'\n      TableCell\n        Text 'bar'\n"
		if got != exp {
			t.Errorf("\nInput   [%#v]\nExpected[%#v]\nGot     [%#v]\n", input, exp, got)
		}
	*/

	tests := []struct {
		input       []byte
		expect      string
		syntaxError bool
	}{
		// ------------------------------------------------------------------------
		{ // 000
			input: []byte("dd"),
			expect: `Paragraph
  Text 'dd'
`,
			syntaxError: false,
		},
		// ------------------------------------------------------------------------
		{ // 001
			input: []byte(`A

`),
			expect: `Paragraph
  Text 'A'
`,
			syntaxError: false,
		},
		// ------------------------------------------------------------------------
		{ // 002
			input: []byte(`
# Main Heading

This is a paragraph.
On 2 lines.

## Sub Heading

With some text after this

## A thrid 

Some other text
- A List
- More List
- Still More List

Some other text 2

- A List
- More List
- Still More List

`),
			expect: `Heading
  Text 'Main Heading'
Paragraph
  Text 'This is a paragraph.\nOn 2 lines.'
Heading
  Text 'Sub Heading'
Paragraph
  Text 'With some text after this'
Heading
  Text 'A thrid'
Paragraph
  Text 'Some other text\n- A List\n- More Lis…'
Paragraph
  Text 'Some other text 2'
List 'tight flags=start'
  ListItem 'flags=start'
    Paragraph
      Text 'A List'
  ListItem
    Paragraph
      Text 'More List'
  ListItem
    Paragraph
      Text 'Still More List'
`,
			syntaxError: false,
		},
		// ------------------------------------------------------------------------
		{ // 003
			input: []byte(`# highlighed sections

A **bold** in a paragraph.

An _underscore_ in a word.

An ` + "`" + `not-a-word` + "`" + ` in a block.

`),
			expect: `Heading
  Text 'highlighed sections'
Paragraph
  Text 'A'
  Strong
    Text 'bold'
  Text 'in a paragraph.'
Paragraph
  Text 'An'
  Emph
    Text 'underscore'
  Text 'in a word.'
Paragraph
  Text 'An'
  Code 'not-a-word'
  Text 'in a block.'
`,
			syntaxError: false,
		},
		// ------------------------------------------------------------------------
		{ // 004
			input: []byte(`# links sections

Code examples:
* https://onlinetool.io/goplayground/#txO7hJ-ibeU : basic markdown => HTML

[![pkg.go.dev](https://pkg.go.dev/badge/github.com/pschlump/markdown)](https://pkg.go.dev/github.com/pschlump/markdown)

For more documentation read [this guide](https://blog.kowalczyk.info/article/cxn3/advanced-markdown-processing-in-go.html)

` + "```" + `go
import (
    "github.com/pschlump/markdown"
)
` + "```" + `

  ` + "```" + `
  Name    | Age
  --------|------
  Bob     ||
  Alice   | 23
  ========|======
  Total   | 23
  ` + "```" + `

{#id3 .myclass fontsize="tiny"}
# Header 1

[Simplified BSD License](LICENSE.txt)
`),
			expect: `Heading
  Text 'links sections'
Paragraph
  Text 'Code examples:\n*'
  Link 'url=https://onlinetool.io/goplayground/#txO7hJ-ibeU'
    Text 'https://onlinetool.io/goplayground/#t…'
  Text ': basic markdown => HTML'
Paragraph
  Text
  Link 'url=https://pkg.go.dev/github.com/pschlump/markdown'
    Text
    Image 'url=https://pkg.go.dev/badge/github.com/pschlump/markdown'
      Text 'pkg.go.dev'
Paragraph
  Text 'For more documentation read'
  Link 'url=https://blog.kowalczyk.info/article/cxn3/advanced-markdown-processing-in-go.html'
    Text 'this guide'
CodeBlock:go 'import (\n    "github.com/pschlump/ma…'
CodeBlock: 'Name    | Age\n  --------|------\n …'
Paragraph
  Text '{#id3 .myclass fontsize="tiny"}'
Heading
  Text 'Header 1'
Paragraph
  Text
  Link 'url=LICENSE.txt'
    Text 'Simplified BSD License'
`,
			syntaxError: false,
		},
	}

	for ii, test := range tests {
		dbgo.Printf("%(green)Test %03d ================================ \n", ii)
		p := NewWithExtensions(CommonExtensions)

		doc := p.Parse([]byte(test.input))
		var buf bytes.Buffer
		ast.Print(&buf, doc)
		got := buf.String()

		if db2 {
			dbgo.Printf("%(cyan)got ->%s<-\n", got)
		}

		if test.expect != got {
			t.Errorf("Test %03d For input ->%s<- expected %s got %s\n", ii, test.input, test.expect, got)
		}
	}

}

const db2 = true
