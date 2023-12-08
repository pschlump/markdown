package parser

import "testing"

// func (p *Parser) isPrefixHeading(data []byte) bool {
func TestIsPrefixHeading(t *testing.T) {
	tests := []struct {
		data   []byte
		expect bool
	}{
		{data: []byte("abc\ndef"), expect: false},
		{data: []byte("#abc\ndef"), expect: true},
		{data: []byte("##abc\ndef"), expect: true},
		{data: []byte("# abc\ndef"), expect: true},
		{data: []byte("## abc\ndef"), expect: true},
		//		{data: []byte(`\\\a`), expect: true},
		//		{data: []byte(`b\\\a`), expect: true},
		//		{data: []byte(`b\\\ac`), expect: true},
		//
		//		{data: []byte(`\\a`), expect: false},
		//		{data: []byte(`\\\\a`), expect: false},
		//		{data: []byte(`\\\\ab`), expect: false},
	}

	p := New()

	for ii, test := range tests {
		if b := p.isPrefixHeading([]byte(test.data)); b != test.expect {
			t.Errorf("Test %d For data ->%s<- expected %v got %v\n", ii, test.data, test.expect, b)
		}
	}

}
