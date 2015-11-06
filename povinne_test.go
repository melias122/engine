package psl

import (
	"testing"

	// "github.com/melias122/psl/komb"
)

func TestPovinne(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		r bool
	}{
		{k: komb.Kombinacia{1}, f: Povinne([]int{1}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 2}, f: Povinne([]int{1}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 2, 3}, f: Povinne([]int{1}, 3, 10), r: true},
		//...

		{k: komb.Kombinacia{2}, f: Povinne([]int{1}, 3, 10), r: false},
		{k: komb.Kombinacia{2, 3}, f: Povinne([]int{1}, 3, 10), r: false},
		{k: komb.Kombinacia{2, 4}, f: Povinne([]int{1}, 3, 10), r: false},
		{k: komb.Kombinacia{2, 9, 10}, f: Povinne([]int{1}, 3, 10), r: false},
		//...

		{k: komb.Kombinacia{1}, f: Povinne([]int{2}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 2}, f: Povinne([]int{2}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 2, 3}, f: Povinne([]int{2}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 3}, f: Povinne([]int{2}, 3, 10), r: false},

		{k: komb.Kombinacia{1}, f: Povinne([]int{2, 4}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 2}, f: Povinne([]int{2, 4}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 2, 3}, f: Povinne([]int{2, 4}, 3, 10), r: false},
		{k: komb.Kombinacia{1, 2, 4}, f: Povinne([]int{2, 4}, 3, 10), r: true},

		{k: komb.Kombinacia{1, 2, 3}, f: Povinne([]int{7, 8, 9, 10}, 3, 10), r: false},

		// {k: komb.Kombinacia{1, 3}, f: Povinne([]int{2, 4, 6, 8}, 3, 10), r: false},
	}
	for _, test := range tests {
		r := test.f.Check(test.k)
		if r != test.r {
			t.Fail()
			t.Log("{", test.k, "}", test.f)
			t.Logf("Expected (%v), got (%v)", test.r, r)
		}
	}
}

func TestPovinneSTL(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		r bool
	}{
		{k: komb.Kombinacia{1}, f: PovinneSTL(map[int][]int{1: {1}}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 2}, f: PovinneSTL(map[int][]int{1: {1}}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 2, 3}, f: PovinneSTL(map[int][]int{1: {1}}, 3, 10), r: true},

		{k: komb.Kombinacia{2}, f: PovinneSTL(map[int][]int{1: {1}}, 3, 10), r: false},
		{k: komb.Kombinacia{2, 3}, f: PovinneSTL(map[int][]int{1: {1}}, 3, 10), r: false},
		{k: komb.Kombinacia{2, 3, 4}, f: PovinneSTL(map[int][]int{1: {1}}, 3, 10), r: false},

		{k: komb.Kombinacia{1}, f: PovinneSTL(map[int][]int{2: {2}}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 2}, f: PovinneSTL(map[int][]int{2: {2}}, 3, 10), r: true},
		{k: komb.Kombinacia{1, 2, 3}, f: PovinneSTL(map[int][]int{2: {2}}, 3, 10), r: true},

		{k: komb.Kombinacia{1, 3}, f: PovinneSTL(map[int][]int{2: {2}}, 3, 10), r: false},
		{k: komb.Kombinacia{2, 3}, f: PovinneSTL(map[int][]int{2: {2}}, 3, 10), r: false},

		{k: komb.Kombinacia{1}, f: PovinneSTL(map[int][]int{1: {2}, 3: {4}}, 3, 10), r: false},
		{k: komb.Kombinacia{2}, f: PovinneSTL(map[int][]int{1: {2}, 3: {4}}, 3, 10), r: true},
		{k: komb.Kombinacia{2, 3}, f: PovinneSTL(map[int][]int{1: {2}, 3: {4}}, 3, 10), r: true},
		{k: komb.Kombinacia{2, 3, 4}, f: PovinneSTL(map[int][]int{1: {2}, 3: {4}}, 3, 10), r: true},
		{k: komb.Kombinacia{2, 3, 5}, f: PovinneSTL(map[int][]int{1: {2}, 3: {4}}, 3, 10), r: false},
	}
	for _, test := range tests {
		r := test.f.Check(test.k)
		if r != test.r {
			t.Fail()
			t.Log("{", test.k, "}", test.f)
			t.Logf("Expected (%v), got (%v)", test.r, r)
		}
	}
}
