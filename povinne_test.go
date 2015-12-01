package psl

import (
	"testing"
)

// func TestFilterPovinne(t *testing.T) {
// 	tests := []struct {
// 		k Kombinacia
// 		f Filter
// 		r bool
// 	}{
// 		{k: Kombinacia{1}, f: NewFilterPovinne([]int{1}, 3, 10), r: true},
// 		{k: Kombinacia{1, 2}, f: NewFilterPovinne([]int{1}, 3, 10), r: true},
// 		{k: Kombinacia{1, 2, 3}, f: NewFilterPovinne([]int{1}, 3, 10), r: true},
// 		//...
//
// 		{k: Kombinacia{2}, f: NewFilterPovinne([]int{1}, 3, 10), r: false},
// 		{k: Kombinacia{2, 3}, f: NewFilterPovinne([]int{1}, 3, 10), r: false},
// 		{k: Kombinacia{2, 4}, f: NewFilterPovinne([]int{1}, 3, 10), r: false},
// 		{k: Kombinacia{2, 9, 10}, f: NewFilterPovinne([]int{1}, 3, 10), r: false},
// 		//...
//
// 		{k: Kombinacia{1}, f: NewFilterPovinne([]int{2}, 3, 10), r: true},
// 		{k: Kombinacia{1, 2}, f: NewFilterPovinne([]int{2}, 3, 10), r: true},
// 		{k: Kombinacia{1, 2, 3}, f: NewFilterPovinne([]int{2}, 3, 10), r: true},
// 		{k: Kombinacia{1, 3}, f: NewFilterPovinne([]int{2}, 3, 10), r: false},
//
// 		{k: Kombinacia{1}, f: NewFilterPovinne([]int{2, 4}, 3, 10), r: true},
// 		{k: Kombinacia{1, 2}, f: NewFilterPovinne([]int{2, 4}, 3, 10), r: true},
// 		{k: Kombinacia{1, 2, 3}, f: NewFilterPovinne([]int{2, 4}, 3, 10), r: false},
// 		{k: Kombinacia{1, 2, 4}, f: NewFilterPovinne([]int{2, 4}, 3, 10), r: true},
//
// 		{k: Kombinacia{1, 2, 3}, f: NewFilterPovinne([]int{7, 8, 9, 10}, 3, 10), r: false},
//
// 		// {k: Kombinacia{1, 3}, f: NewFilterPovinne([]int{2, 4, 6, 8}, 3, 10), r: false},
// 	}
// 	for _, test := range tests {
// 		r := test.f.Check(test.k)
// 		if r != test.r {
// 			t.Fail()
// 			t.Log("{", test.k, "}", test.f)
// 			t.Logf("Expected (%v), got (%v)", test.r, r)
// 		}
// 	}
// }

func TestFilterPovinneSTL(t *testing.T) {
	tests := []struct {
		k Kombinacia
		f Filter
		r bool
	}{
		{k: Kombinacia{1}, f: NewFilterPovinneSTL(MapInts{1: {1}}, 3, 10), r: true},
		{k: Kombinacia{1, 2}, f: NewFilterPovinneSTL(MapInts{1: {1}}, 3, 10), r: true},
		{k: Kombinacia{1, 2, 3}, f: NewFilterPovinneSTL(MapInts{1: {1}}, 3, 10), r: true},

		{k: Kombinacia{2}, f: NewFilterPovinneSTL(MapInts{1: {1}}, 3, 10), r: false},
		{k: Kombinacia{2, 3}, f: NewFilterPovinneSTL(MapInts{1: {1}}, 3, 10), r: false},
		{k: Kombinacia{2, 3, 4}, f: NewFilterPovinneSTL(MapInts{1: {1}}, 3, 10), r: false},

		{k: Kombinacia{1}, f: NewFilterPovinneSTL(MapInts{2: {2}}, 3, 10), r: true},
		{k: Kombinacia{1, 2}, f: NewFilterPovinneSTL(MapInts{2: {2}}, 3, 10), r: true},
		{k: Kombinacia{1, 2, 3}, f: NewFilterPovinneSTL(MapInts{2: {2}}, 3, 10), r: true},

		{k: Kombinacia{1, 3}, f: NewFilterPovinneSTL(MapInts{2: {2}}, 3, 10), r: false},
		{k: Kombinacia{2, 3}, f: NewFilterPovinneSTL(MapInts{2: {2}}, 3, 10), r: false},

		{k: Kombinacia{1}, f: NewFilterPovinneSTL(MapInts{1: {2}, 3: {4}}, 3, 10), r: false},
		{k: Kombinacia{2}, f: NewFilterPovinneSTL(MapInts{1: {2}, 3: {4}}, 3, 10), r: true},
		{k: Kombinacia{2, 3}, f: NewFilterPovinneSTL(MapInts{1: {2}, 3: {4}}, 3, 10), r: true},
		{k: Kombinacia{2, 3, 4}, f: NewFilterPovinneSTL(MapInts{1: {2}, 3: {4}}, 3, 10), r: true},
		{k: Kombinacia{2, 3, 5}, f: NewFilterPovinneSTL(MapInts{1: {2}, 3: {4}}, 3, 10), r: false},
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
