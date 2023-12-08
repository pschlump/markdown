package parser

import "testing"

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
