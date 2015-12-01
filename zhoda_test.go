package psl

import (
	"testing"
)

func TestZhodaPresun(t *testing.T) {
	tests := []struct {
		k0, k1 Kombinacia
		presun
	}{
		{Kombinacia{1, 2, 3}, Kombinacia{4, 5, 6}, presun{}},
		{Kombinacia{1, 2, 3}, Kombinacia{1, 4, 5}, presun{{1, 1}}},
		{Kombinacia{1, 2, 3}, Kombinacia{1, 2, 3}, presun{{1, 1}, {2, 2}, {3, 3}}},
		{
			Kombinacia{1, 2, 10, 20, 30, 40, 50},
			Kombinacia{10, 20, 30, 50, 60, 61, 62},
			presun{{3, 1}, {4, 2}, {5, 3}, {7, 4}},
		},
	}
	for _, test := range tests {
		presun := ZhodaPresun(test.k0, test.k1)
		if test.presun.String() != presun.String() {
			t.Errorf("Expected: %s, Got: %s", test.presun, presun)
		}
	}
}

func TestZhoda(t *testing.T) {
	tests := []struct {
		k1, k2 Kombinacia
		zhoda  int
	}{
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 2, 3, 4, 5}, 5},
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 2, 3, 4, 6}, 4},
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 2, 3, 9, 10}, 3},
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 2, 8, 9, 10}, 2},
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 7, 8, 9, 10}, 1},
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{6, 7, 8, 9, 10}, 0},
	}
	for _, test := range tests {
		zhoda := Zhoda(test.k1, test.k2)
		if zhoda != test.zhoda {
			t.Fatalf("Excepted: (%d), Have: (%d)", test.zhoda, zhoda)
		}
	}
}

func BenchmarkZhoda(b *testing.B) {
	k0 := Kombinacia{1, 3, 4, 5, 6}
	k1 := Kombinacia{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Zhoda(k0, k1)
	}
}

func TestFilterZhodaRange(t *testing.T) {
	tests := []struct {
		k Kombinacia
		f Filter
		w bool
	}{
		{Kombinacia{1}, NewFilterZhodaRange(2, 3, Kombinacia{1}, 3), true},
		{Kombinacia{1}, NewFilterZhodaRange(0, 0, Kombinacia{1}, 3), false},
		{Kombinacia{1}, NewFilterZhodaRange(0, 3, Kombinacia{1}, 3), true},
		{Kombinacia{1}, NewFilterZhodaRange(0, 3, Kombinacia{2}, 3), true},
		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(2, 2, Kombinacia{1, 2}, 3), true},
		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 2, Kombinacia{1, 2, 3}, 3), false},
		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 3, Kombinacia{4, 5, 6}, 3), true},
		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(1, 3, Kombinacia{4, 5, 6}, 3), false},

		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 3, Kombinacia{1, 2, 3}, 3), true},
		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(1, 3, Kombinacia{1, 2, 3}, 3), true},
		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(2, 3, Kombinacia{1, 2, 3}, 3), true},
		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(3, 3, Kombinacia{1, 2, 3}, 3), true},

		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 2, Kombinacia{1, 2, 3}, 3), false},
		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 1, Kombinacia{1, 2, 3}, 3), false},
		{Kombinacia{1, 2, 3}, NewFilterZhodaRange(0, 0, Kombinacia{1, 2, 3}, 3), false},
	}
	for i, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) (test %d)", test.w, ok, i+1)
		}
	}
}
