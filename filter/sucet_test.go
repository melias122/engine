package filter

import (
	"testing"

	"gitlab.com/melias122/engine"
)

func TestFilterSucet(t *testing.T) {
	tests := []struct {
		k engine.Kombinacia
		f Filter
		w bool
	}{
		{engine.Kombinacia{1}, NewFilterSucet(14, 14, 3), true},
		{engine.Kombinacia{1, 11}, NewFilterSucet(14, 14, 3), true},
		{engine.Kombinacia{1, 2, 10}, NewFilterSucet(14, 14, 3), false},
		{engine.Kombinacia{1, 2, 11}, NewFilterSucet(14, 14, 3), true},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(0, 14, 5), false},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(0, 15, 5), true},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(15, 15, 5), true},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(15, 55, 5), true},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(16, 16, 5), false},

		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(30, 50, 5), false},
		{engine.Kombinacia{1, 2, 3, 4, 25}, NewFilterSucet(30, 50, 5), true},
		{engine.Kombinacia{1, 2, 3, 4, 45}, NewFilterSucet(30, 50, 5), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
