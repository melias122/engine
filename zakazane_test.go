package psl

import (
	"testing"
)

func TestFilterZakazane(t *testing.T) {
	tests := []struct {
		k Kombinacia
		f Filter
		w bool
	}{
		{Kombinacia{1}, NewFilterZakazane([]int{2, 3}, 5, 35), true},
		{Kombinacia{1, 2}, NewFilterZakazane([]int{2, 3}, 5, 35), false},
		{Kombinacia{1, 3}, NewFilterZakazane([]int{2, 3}, 5, 35), false},
		{Kombinacia{1, 4}, NewFilterZakazane([]int{2, 3}, 5, 35), true},
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
		k Kombinacia
		f Filter
		w bool
	}{
		{Kombinacia{1}, NewFilterZakazaneSTL(map[int][]int{1: {1}}, 5, 35), false},
		{Kombinacia{1}, NewFilterZakazaneSTL(map[int][]int{1: {2}}, 5, 35), true},
		{Kombinacia{2, 3, 5}, NewFilterZakazaneSTL(map[int][]int{2: {3}}, 5, 35), false},
		{Kombinacia{2, 3, 6}, NewFilterZakazaneSTL(map[int][]int{2: {4}}, 5, 35), true},
		{Kombinacia{2, 3, 6}, NewFilterZakazaneSTL(map[int][]int{3: {7}}, 5, 35), true},

		{Kombinacia{2}, NewFilterZakazaneSTL(map[int][]int{3: {3}}, 5, 35), true},
		{Kombinacia{2, 3}, NewFilterZakazaneSTL(map[int][]int{1: {2}, 2: {3}}, 5, 35), false},
		{Kombinacia{2, 4, 6}, NewFilterZakazaneSTL(map[int][]int{1: {2}, 2: {4}, 3: {6}}, 5, 35), false},
	}
	for i, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v), test: %d", test.w, ok, i+1)
		}
	}
}
