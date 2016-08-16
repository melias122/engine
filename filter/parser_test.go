package filter

import (
	"testing"

	"github.com/melias122/engine"
)

func TestParseNtica(t *testing.T) {
	tests := []struct {
		s string
		w engine.Tica
		n int
	}{
		{s: "", w: engine.Tica{}, n: 5},
		{s: "             \t\n   \t\t\n", w: engine.Tica{}, n: 5},
		{s: "5", w: engine.Tica{5, 0, 0, 0, 0}, n: 5},
		{s: "5 0 0 0 0", w: engine.Tica{5, 0, 0, 0, 0}, n: 5},
		{s: "5 0 0 0 0 0", w: engine.Tica{}, n: 5},
	}
	for _, test := range tests {
		n, e := ParseNtica(test.n, test.s)
		if e != nil {
			if n.String() != test.w.String() {
				t.Errorf("Expected: %s, Got: %s", test.w, n)
			}
		}
	}
}

func TestParseXtica(t *testing.T) {
	tests := []struct {
		s    string
		n, m int
		w    engine.Tica
		e    bool
	}{
		{n: 5, m: 35, s: "", w: engine.Tica{}, e: true},
		{n: 5, m: 35, s: "    ", w: engine.Tica{}, e: true},
		{n: 5, m: 35, s: "    \t\t\t\t\t\t \n\n  \t      ", w: engine.Tica{}, e: true},
		{n: 5, m: 35, s: "1 2", e: true}, // 1+2 != 5
		{n: 5, m: 35, s: "1 2 0 0 1", e: true},
		{n: 5, m: 35, s: "1 2 2 2", e: true},
		{n: 5, m: 35, s: "5,0,0", e: true},
		{n: 5, m: 35, s: "5;", e: true},

		{n: 5, m: 35, s: "5", w: engine.Tica{5, 0, 0, 0}},
		{n: 5, m: 35, s: "5 ", w: engine.Tica{5, 0, 0, 0}},
		{n: 5, m: 35, s: "3 2", w: engine.Tica{3, 2, 0, 0}},
		{n: 5, m: 35, s: "1 2 0 2", w: engine.Tica{1, 2, 0, 2}},
	}
	for _, test := range tests {
		x, e := ParseXtica(test.n, test.m, test.s)
		if e != nil {
			if x.String() != test.w.String() {
				t.Errorf("Expected: %s, Got: %s", test.w, x)
			}
		} else {
			if test.e {
				t.Errorf("Expected: error (%s)", test.s)
			}
		}
	}
}
