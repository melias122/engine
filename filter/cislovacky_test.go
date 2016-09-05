package filter

import (
	"testing"

	"gitlab.com/melias122/engine"
)

func TestCislovackyRange(t *testing.T) {
	tests := []struct {
		k engine.Kombinacia
		f Filter
		w bool
	}{
		{engine.Kombinacia{1}, NewFilterCislovackyRange(0, 0, engine.P, 1), true},
		{engine.Kombinacia{1}, NewFilterCislovackyRange(0, 1, engine.P, 1), true},
		{engine.Kombinacia{1}, NewFilterCislovackyRange(1, 1, engine.P, 1), false},
		{engine.Kombinacia{1}, NewFilterCislovackyRange(1, 1, engine.P, 3), true},
		{engine.Kombinacia{1, 2}, NewFilterCislovackyRange(1, 1, engine.P, 3), true},
		{engine.Kombinacia{1, 2, 3}, NewFilterCislovackyRange(1, 1, engine.P, 3), true},
		{engine.Kombinacia{1, 2, 4}, NewFilterCislovackyRange(1, 1, engine.P, 3), false},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterCislovackyRange(0, 1, engine.P, 5), false},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterCislovackyRange(0, 2, engine.P, 5), true},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterCislovackyRange(2, 2, engine.P, 5), true},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterCislovackyRange(2, 3, engine.P, 5), true},
		{engine.Kombinacia{1, 2, 3, 4, 5}, NewFilterCislovackyRange(3, 3, engine.P, 5), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}

func TestCislovackyExact(t *testing.T) {
	tests := []struct {
		k engine.Kombinacia
		// f func() (Filter, error)

		ints []int
		c    engine.Cislovacka
		n    int

		w bool
	}{
		{engine.Kombinacia{2}, []int{0, 2}, engine.P, 4, true},
		{engine.Kombinacia{2, 4}, []int{0, 2}, engine.P, 4, true},
		{engine.Kombinacia{2, 4, 6, 7}, []int{0, 2}, engine.P, 4, false},
		{engine.Kombinacia{2, 4, 6, 7}, []int{1, 3}, engine.P, 4, true},
		{engine.Kombinacia{2, 4, 7, 9}, []int{1, 3}, engine.P, 4, false},

		{engine.Kombinacia{1, 3, 7, 9}, []int{1, 3}, engine.P, 4, false},
	}
	for _, test := range tests {
		f, err := NewFilterCislovackyExact(test.ints, test.c, test.n)
		if err != nil {
			t.Fatal(err)
		}
		ok := f.Check(test.k)
		if ok != test.w {
			t.Error(f)
			t.Error(test)
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
