package filter

import (
	"testing"

	"github.com/melias122/engine"
)

func TestFilterZhodaRange(t *testing.T) {
	tests := []struct {
		k engine.Kombinacia
		f Filter
		w bool
	}{
		{engine.Kombinacia{1}, NewFilterZhodaRange(2, 3, engine.Kombinacia{1}, 3), true},
		{engine.Kombinacia{1}, NewFilterZhodaRange(0, 0, engine.Kombinacia{1}, 3), false},
		{engine.Kombinacia{1}, NewFilterZhodaRange(0, 3, engine.Kombinacia{1}, 3), true},
		{engine.Kombinacia{1}, NewFilterZhodaRange(0, 3, engine.Kombinacia{2}, 3), true},
		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(2, 2, engine.Kombinacia{1, 2}, 3), true},
		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 2, engine.Kombinacia{1, 2, 3}, 3), false},
		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 3, engine.Kombinacia{4, 5, 6}, 3), true},
		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(1, 3, engine.Kombinacia{4, 5, 6}, 3), false},

		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 3, engine.Kombinacia{1, 2, 3}, 3), true},
		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(1, 3, engine.Kombinacia{1, 2, 3}, 3), true},
		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(2, 3, engine.Kombinacia{1, 2, 3}, 3), true},
		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(3, 3, engine.Kombinacia{1, 2, 3}, 3), true},

		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 2, engine.Kombinacia{1, 2, 3}, 3), false},
		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 1, engine.Kombinacia{1, 2, 3}, 3), false},
		{engine.Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 0, engine.Kombinacia{1, 2, 3}, 3), false},
	}
	for i, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) (test %d)", test.w, ok, i+1)
		}
	}
}
