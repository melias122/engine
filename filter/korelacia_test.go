package filter

import (
	"testing"

	"github.com/melias122/psl/komb"
)

func TestKorelacia(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1, 14, 15, 17, 19}, NewKorelacia(5, 35, 0.0, 1.0, komb.Kombinacia{2, 7, 13, 32, 35}), true},  // "0.34137300"
		{komb.Kombinacia{1, 14, 15, 17, 19}, NewKorelacia(5, 35, 0.0, 0.3, komb.Kombinacia{2, 7, 13, 32, 35}), false}, // "0.34137300"
		{komb.Kombinacia{1, 14, 15, 17, 19}, NewKorelacia(5, 35, 0.4, 1.0, komb.Kombinacia{2, 7, 13, 32, 35}), false}, // "0.34137300"
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewKorelacia(5, 35, 1.0, 1.0, komb.Kombinacia{1, 2, 3, 4, 5}), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewKorelacia(5, 35, 0.998, 0.999, komb.Kombinacia{1, 2, 3, 4, 5}), false},
		{komb.Kombinacia{1, 2, 3, 4}, NewKorelacia(5, 35, 1.0, 1.0, komb.Kombinacia{1, 2, 3, 4}), true},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
