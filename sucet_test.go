package psl

import (
	"testing"
)

func TestFilterSucet(t *testing.T) {
	tests := []struct {
		k Kombinacia
		f Filter
		w bool
	}{
		{Kombinacia{1}, NewFilterSucet(3, 14, 14), true},
		{Kombinacia{1, 11}, NewFilterSucet(3, 14, 14), true},
		{Kombinacia{1, 2, 10}, NewFilterSucet(3, 14, 14), false},
		{Kombinacia{1, 2, 11}, NewFilterSucet(3, 14, 14), true},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(5, 0, 14), false},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(5, 0, 15), true},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(5, 15, 15), true},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(5, 15, 55), true},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(5, 16, 16), false},

		{Kombinacia{1, 2, 3, 4, 5}, NewFilterSucet(5, 30, 50), false},
		{Kombinacia{1, 2, 3, 4, 25}, NewFilterSucet(5, 30, 50), true},
		{Kombinacia{1, 2, 3, 4, 45}, NewFilterSucet(5, 30, 50), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
