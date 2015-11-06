package psl

import (
	"testing"

	// "github.com/melias122/psl/komb"
)

func TestXtica(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewXtica(5, 35, komb.Tica{5, 0, 0, 0}), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewXtica(5, 35, komb.Tica{4, 1, 0, 0}), false},
		{komb.Kombinacia{1, 2, 3, 4, 11}, NewXtica(5, 35, komb.Tica{4, 1, 0, 0}), true},

		{komb.Kombinacia{1}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), true},
		{komb.Kombinacia{1, 2}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), false},
		{komb.Kombinacia{1, 11}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), true},
		{komb.Kombinacia{1, 11, 12}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), false},
		{komb.Kombinacia{1, 11, 20}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), false},
		{komb.Kombinacia{1, 11, 21}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), true},
		{komb.Kombinacia{1, 11, 21, 22}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), false},
		{komb.Kombinacia{1, 11, 21, 30}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), false},
		{komb.Kombinacia{1, 11, 21, 31}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), true},
		{komb.Kombinacia{1, 11, 21, 31}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), true},
		{komb.Kombinacia{1, 11, 21, 31, 32}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), true},
		{komb.Kombinacia{1, 11, 21, 31, 35}, NewXtica(5, 35, komb.Tica{1, 1, 1, 2}), true},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
