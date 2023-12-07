package rtf

// func (r *Renderer) writeDocumentHeader(w io.Writer) {

import (
	"bytes"
	"testing"
)

func Test_writeDocumentHeader(t *testing.T) {
	buf := &bytes.Buffer{}
	//	code := []byte(`println("hello")
	//more code //<<4>>
	//bliep bliep
	//`)
	out := `<!DOCTYPE html>
<html>
<head>
  <title></title>
  <meta name="GENERATOR" content="github.com/pschlump/markdown markdown processor for Go">
  <meta charset="utf-8">
</head>
<body>

`
	opts := RendererOptions{}
	opts.Comments = [][]byte{[]byte("//")}
	opts.Flags = CompletePage

	r := NewRenderer(opts)
	r.writeDocumentHeader(buf)

	expect := buf.String()
	if expect != out {
		t.Errorf("Expected -->>%s<<-- got -->>%s<<--\n", out, expect)
	}
}
