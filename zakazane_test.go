package psl

import (
	"testing"

	// "github.com/melias122/psl/komb"
)

func TestZakazane(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1}, Zakazane([]int{2, 3}, 5, 35), true},
		{komb.Kombinacia{1, 2}, Zakazane([]int{2, 3}, 5, 35), false},
		{komb.Kombinacia{1, 3}, Zakazane([]int{2, 3}, 5, 35), false},
		{komb.Kombinacia{1, 4}, Zakazane([]int{2, 3}, 5, 35), true},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}

func TestZakazaneSTL(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1}, ZakazaneSTL(map[int][]int{1: {1}}, 5, 35), false},
		{komb.Kombinacia{1}, ZakazaneSTL(map[int][]int{1: {2}}, 5, 35), true},
		{komb.Kombinacia{2, 3, 5}, ZakazaneSTL(map[int][]int{2: {3}}, 5, 35), false},
		{komb.Kombinacia{2, 3, 6}, ZakazaneSTL(map[int][]int{2: {4}}, 5, 35), true},
		{komb.Kombinacia{2, 3, 6}, ZakazaneSTL(map[int][]int{3: {7}}, 5, 35), true},

		{komb.Kombinacia{2}, ZakazaneSTL(map[int][]int{3: {3}}, 5, 35), true},
		{komb.Kombinacia{2, 3}, ZakazaneSTL(map[int][]int{1: {2}, 2: {3}}, 5, 35), false},
		{komb.Kombinacia{2, 4, 6}, ZakazaneSTL(map[int][]int{1: {2}, 2: {4}, 3: {6}}, 5, 35), false},
	}
	for i, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v), test: %d", test.w, ok, i+1)
		}
	}
}
