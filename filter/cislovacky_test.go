package filter

import (
	"testing"

	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

func TestCislovacky(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1}, NewCislovacky(1, 0, 0, num.IsP, ""), true},
		{komb.Kombinacia{1}, NewCislovacky(1, 0, 1, num.IsP, ""), true},
		{komb.Kombinacia{1}, NewCislovacky(1, 1, 1, num.IsP, ""), false},
		{komb.Kombinacia{1}, NewCislovacky(3, 1, 1, num.IsP, ""), true},
		{komb.Kombinacia{1, 2}, NewCislovacky(3, 1, 1, num.IsP, ""), true},
		{komb.Kombinacia{1, 2, 3}, NewCislovacky(3, 1, 1, num.IsP, ""), true},
		{komb.Kombinacia{1, 2, 4}, NewCislovacky(3, 1, 1, num.IsP, ""), false},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewCislovacky(5, 0, 1, num.IsP, ""), false},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewCislovacky(5, 0, 2, num.IsP, ""), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewCislovacky(5, 2, 2, num.IsP, ""), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewCislovacky(5, 2, 3, num.IsP, ""), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewCislovacky(5, 3, 3, num.IsP, ""), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
