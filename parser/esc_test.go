package parser

import "testing"

/*
func TestIsEscape(t *testing.T) {
	if x := `\a`; !isEscape([]byte(x), 1) {
		t.Errorf("expected escape for %q, got false", x)
	}
	if x := `\\\a`; !isEscape([]byte(x), 3) {
		t.Errorf("expected escape for %q, got false", x)
	}
	if x := `b\\\a`; !isEscape([]byte(x), 4) {
		t.Errorf("expected escape for %q, got false", x)
	}

	if x := `\\a`; isEscape([]byte(x), 2) {
		t.Errorf("expected no escape for %q, got true", x)
	}
	if x := `\\\\a`; isEscape([]byte(x), 4) {
		t.Errorf("expected no escape for %q, got true", x)
	}
}
*/

func TestIsEscape2(t *testing.T) {
	tests := []struct {
		data    []byte
		lenData int
		expect  bool
	}{
		{data: []byte(`\a`), lenData: 1, expect: true},
		{data: []byte(`\\\a`), lenData: 3, expect: true},
		{data: []byte(`b\\\a`), lenData: 4, expect: true},
		{data: []byte(`b\\\ac`), lenData: 4, expect: true},
		{data: []byte(`\\a`), lenData: 2, expect: false},
		{data: []byte(`\\\\a`), lenData: 4, expect: false},
		{data: []byte(`\\\\ab`), lenData: 4, expect: false},
	}

	for ii, test := range tests {
		if b := isEscape([]byte(test.data), test.lenData); b != test.expect {
			t.Errorf("Test %d For data ->%s<- expected %v got %v\n", ii, test.data, test.expect, b)
		}
	}

}
