package filter

import (
	"testing"

	"gitlab.com/melias122/engine"
)

func TestFilterZakazane(t *testing.T) {
	tests := []struct {
		k engine.Kombinacia
		f Filter
		w bool
	}{
		{engine.Kombinacia{1}, NewFilterZakazane([]int{2, 3}, 5, 35), true},
		{engine.Kombinacia{1, 2}, NewFilterZakazane([]int{2, 3}, 5, 35), false},
		{engine.Kombinacia{1, 3}, NewFilterZakazane([]int{2, 3}, 5, 35), false},
		{engine.Kombinacia{1, 4}, NewFilterZakazane([]int{2, 3}, 5, 35), true},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}

func TestFilterZakazaneSTL(t *testing.T) {
	tests := []struct {
		k engine.Kombinacia
		f Filter
		w bool
	}{
		{engine.Kombinacia{1}, NewFilterZakazaneSTL(MapInts{1: {1}}, 5, 35), false},
		{engine.Kombinacia{1}, NewFilterZakazaneSTL(MapInts{1: {2}}, 5, 35), true},
		{engine.Kombinacia{2, 3, 5}, NewFilterZakazaneSTL(MapInts{2: {3}}, 5, 35), false},
		{engine.Kombinacia{2, 3, 6}, NewFilterZakazaneSTL(MapInts{2: {4}}, 5, 35), true},
		{engine.Kombinacia{2, 3, 6}, NewFilterZakazaneSTL(MapInts{3: {7}}, 5, 35), true},

		{engine.Kombinacia{2}, NewFilterZakazaneSTL(MapInts{3: {3}}, 5, 35), true},
		{engine.Kombinacia{2, 3}, NewFilterZakazaneSTL(MapInts{1: {2}, 2: {3}}, 5, 35), false},
		{engine.Kombinacia{2, 4, 6}, NewFilterZakazaneSTL(MapInts{1: {2}, 2: {4}, 3: {6}}, 5, 35), false},
	}
	for i, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v), test: %d", test.w, ok, i+1)
		}
	}
}
