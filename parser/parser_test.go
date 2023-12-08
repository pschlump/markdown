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

// todo PJS - xyzzy
// todo PJS - xyzzy
// todo PJS - xyzzy
// todo PJS - xyzzy
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
		{
			input: []byte("dd"),
			expect: `Paragraph
  Text 'dd'
`,
			syntaxError: false,
		},
	}

	// p := New()
	p := NewWithExtensions(CommonExtensions)

	for ii, test := range tests {
		doc := p.Parse([]byte(test.input))
		var buf bytes.Buffer
		ast.Print(&buf, doc)
		got := buf.String()

		if db2 {
			dbgo.Printf("%(red)got ->%s<-\n", got)
		}

		if test.expect != got {
			t.Errorf("Test %d For input ->%s<- expected %s got %s\n", ii, test.input, test.expect, got)
		}
	}

}

const db2 = true
