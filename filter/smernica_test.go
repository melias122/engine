package filter

import (
	"testing"

	"github.com/melias122/psl/komb"
)

func TestSmernica(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{14, 16, 26, 27}, NewSmernica(5, 35, 0.0, 1.0), true},         // 0.499...
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.0, 1.0), true},     // 0.499...
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.0, 0.49), true},    // 0.499...
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.49, 0.5), true},    // 0.499...
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.499, 0.499), true}, // 0.499...
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.5, 0.5), true},
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.500000001, 0.6), false}, // 0.499...
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) (test %v)", test.w, ok, test)
		}
	}
}
