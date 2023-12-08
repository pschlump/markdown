package parser

import (
	"testing"

	"github.com/pschlump/dbgo"
)

// Test: func (p *Parser) isPrefixHeading(data []byte) bool {
// From block.go

func TestIsPrefixHeading(t *testing.T) {

	tests := []struct {
		data   []byte
		expect bool
	}{
		{data: []byte("abc\ndef"), expect: false},
		{data: []byte("#abc\ndef"), expect: false},  // why?  Is this a defect?
		{data: []byte("##abc\ndef"), expect: false}, // why?  Is this a defect?
		{data: []byte("# abc\ndef"), expect: true},
		{data: []byte("## abc\ndef"), expect: true},
		{data: []byte("##,# abc\ndef"), expect: false},
		{data: []byte("#,# abc\ndef"), expect: false},
	}

	p := New()

	for ii, test := range tests {
		if b := p.isPrefixHeading(test.data); b != test.expect {
			t.Errorf("Test %d For data ->%s<- expected %v got %v\n", ii, test.data, test.expect, b)
		}
	}

}

// func skipSpace(data []byte, i int) int {
func TestSkipSpace(t *testing.T) {

	tests := []struct {
		data   []byte
		start  int
		expect int
	}{
		{data: []byte("dd"), expect: 0},
		{data: []byte(" dd"), expect: 1},
		{data: []byte("  dd"), expect: 2},
		{data: []byte("center"), expect: 0},
	}

	for ii, test := range tests {
		b := skipSpace(test.data, test.start)
		if b != test.expect {
			t.Errorf("Test %d For data ->%s<- expected %v got %v\n", ii, test.data, test.expect, b)
		}
	}

}

// func (p *Parser) htmlFindTag(data []byte) (string, bool) {
func TestHtmlFindTag(t *testing.T) {

	tests := []struct {
		data   []byte
		expect bool
	}{
		{data: []byte("dd"), expect: true},      // <dd>
		{data: []byte("dd "), expect: true},     // <dd >
		{data: []byte("  dd"), expect: false},   // <  dd>
		{data: []byte("center"), expect: false}, // <dd>
	}

	p := New()

	for ii, test := range tests {
		k, b := p.htmlFindTag(test.data)
		if b != test.expect {
			t.Errorf("Test %d For data ->%s<- expected %v got %v\n", ii, test.data, test.expect, b)
		}
		_ = k
	}

}

// func isFenceLine(data []byte, syntax *string, oldmarker string) (end int, marker string) {
func TestIsFenceLine(t *testing.T) {
	tests := []struct {
		data            []byte
		syntaxRequested bool
		wantEnd         int
		wantMarker      string
		wantSyntax      string
	}{
		{ // 0
			data:       []byte("```"),
			wantEnd:    3,
			wantMarker: "```",
		},
		{ // 1
			data:       []byte("```\nstuff here\n"),
			wantEnd:    4,
			wantMarker: "```",
		},
		{ // 2
			data:            []byte("```\nstuff here\n"),
			syntaxRequested: true,
			wantEnd:         4,
			wantMarker:      "```",
		},
		{ // 3
			data:    []byte("stuff here\n```\n"),
			wantEnd: 0,
		},
		{ // 4
			data:            []byte("```"),
			syntaxRequested: true,
			wantEnd:         3,
			wantMarker:      "```",
		},
		{ // 5
			data:            []byte("``` go"),
			syntaxRequested: true,
			wantEnd:         6,
			wantMarker:      "```",
			wantSyntax:      "go",
		},
		{ // 6
			data:            []byte("~~~"),
			syntaxRequested: true,
			wantEnd:         3,
			wantMarker:      "~~~",
		},
		{ // 7
			data:            []byte("~~~~~"),
			syntaxRequested: true,
			wantEnd:         5,
			wantMarker:      "~~~~~",
		},
		{ // 8
			data:            []byte("~~~~~ python"),
			syntaxRequested: true,
			wantEnd:         12,
			wantMarker:      "~~~~~",
			wantSyntax:      "python",
		},
		{ // 9
			data:            []byte("~~"),
			syntaxRequested: true,
			wantEnd:         0,
			wantMarker:      "~~",
		},
	}

	s := ""
	for ii, test := range tests {
		var syntax *string
		syntax = &s
		if test.syntaxRequested {
			syntax = new(string)
		}
		end, marker := isFenceLine(test.data, syntax, "")
		if db1 {
			dbgo.Printf("%(yellow)test %d: end= %v marker=%v syntax=->%(cyan)%s%(yellow)<-\n", ii, end, marker, *syntax)
		}
		if end == 0 && test.wantEnd == 0 {
		} else if end > 0 && test.wantEnd == 0 {
			t.Errorf("got end %v, want %v - this indicates finding a block marke where there should not be a block marker", end, test.wantEnd)
		} else {
			if got, want := end, test.wantEnd; got != want {
				t.Errorf("got end %v, want %v", got, want)
			}
			if got, want := marker, test.wantMarker; got != want {
				t.Errorf("got marker %q, want %q", got, want)
			}
			if test.syntaxRequested {
				if got, want := *syntax, test.wantSyntax; got != want {
					t.Errorf("got syntax %q, want %q", got, want)
				}
			}
		}
	}
}

const db1 = false
