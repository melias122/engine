package psl

import (
	"testing"

	// "github.com/melias122/psl/komb"
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

func TestZhoda(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1}, ZhodaRange(3, 2, 3, komb.Kombinacia{1}), true},
		{komb.Kombinacia{1}, ZhodaRange(3, 0, 0, komb.Kombinacia{1}), false},
		{komb.Kombinacia{1}, ZhodaRange(3, 0, 3, komb.Kombinacia{1}), true},
		{komb.Kombinacia{1}, ZhodaRange(3, 0, 3, komb.Kombinacia{2}), true},
		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 2, 2, komb.Kombinacia{1, 2}), true},
		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 0, 2, komb.Kombinacia{1, 2, 3}), false},
		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 0, 3, komb.Kombinacia{4, 5, 6}), true},
		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 1, 3, komb.Kombinacia{4, 5, 6}), false},

		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 0, 3, komb.Kombinacia{1, 2, 3}), true},
		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 1, 3, komb.Kombinacia{1, 2, 3}), true},
		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 2, 3, komb.Kombinacia{1, 2, 3}), true},
		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 3, 3, komb.Kombinacia{1, 2, 3}), true},

		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 0, 2, komb.Kombinacia{1, 2, 3}), false},
		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 0, 1, komb.Kombinacia{1, 2, 3}), false},
		{komb.Kombinacia{1, 2, 3}, ZhodaRange(3, 0, 0, komb.Kombinacia{1, 2, 3}), false},
	}
	for i, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) (test %d)", test.w, ok, i+1)
		}
	}
}
