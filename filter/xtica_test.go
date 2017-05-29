package filter

import (
	"testing"

	"github.com/melias122/engine/engine"
)

func TestFilterXtica(t *testing.T) {
	tests := []struct {
		k engine.Kombinacia
		f Filter
		w bool
	}{
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterXtica(5, 35, engine.Tica{5, 0, 0, 0}), true},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterXtica(5, 35, engine.Tica{4, 1, 0, 0}), false},
		{engine.Kombinacia{1, 2, 3, 4, 11}, NewFilterXtica(5, 35, engine.Tica{4, 1, 0, 0}), true},

		{engine.Kombinacia{1}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), true},
		{engine.Kombinacia{1, 2}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), false},
		{engine.Kombinacia{1, 11}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), true},
		{engine.Kombinacia{1, 11, 12}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), false},
		{engine.Kombinacia{1, 11, 20}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), false},
		{engine.Kombinacia{1, 11, 21}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), true},
		{engine.Kombinacia{1, 11, 21, 22}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), false},
		{engine.Kombinacia{1, 11, 21, 30}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), false},
		{engine.Kombinacia{1, 11, 21, 31}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), true},
		{engine.Kombinacia{1, 11, 21, 31}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), true},
		{engine.Kombinacia{1, 11, 21, 31, 32}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), true},
		{engine.Kombinacia{1, 11, 21, 31, 35}, NewFilterXtica(5, 35, engine.Tica{1, 1, 1, 2}), true},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
