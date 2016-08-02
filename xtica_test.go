package engine

import (
	"testing"
)

func TestXcislaContains(t *testing.T) {
	tests := []struct {
		xcisla Xcisla
		test   map[Tab]bool
	}{
		{xcisla: Xcisla{}, test: map[Tab]bool{Tab{}: false, Tab{1, 1}: false}},
		{xcisla: Xcisla{{0, 0}}, test: map[Tab]bool{Tab{0, 0}: true, Tab{1, 1}: false}},
		{
			xcisla: Xcisla{
				{1, 1}, {2, 2}, {4, 0},
			},
			test: map[Tab]bool{
				Tab{0, 0}: false, Tab{1, 1}: true, Tab{2, 2}: true, Tab{3, 3}: false,
				Tab{3, 0}: false, Tab{3, 1}: false, Tab{4, 0}: true, Tab{4, 1}: false,
				Tab{5, 1}: false,
			},
		},
	}
	for _, test := range tests {
		for tab, exp := range test.test {
			if ok := test.xcisla.Contains(tab); ok != exp {
				t.Fatalf("Expected: (%v:%v), got: (%v:%v)", tab, exp, test.xcisla, ok)
			}
		}
	}
}

func TestXtica(t *testing.T) {
	tests := []struct {
		m int
		t Kombinacia
		w string
	}{
		{9, Kombinacia{1, 2, 3, 4, 5}, "5"},
		{10, Kombinacia{1, 2, 3, 4, 5}, "5"},
		{11, Kombinacia{1, 2, 3, 4, 5}, "5 0"},
		{11, Kombinacia{1, 10, 11, 12, 13}, "2 3"},
		{90, Kombinacia{1, 10, 11, 12, 13}, "2 3 0 0 0 0 0 0 0"},
		{90, Kombinacia{1, 10, 11, 12, 90}, "2 2 0 0 0 0 0 0 1"},
		{90, Kombinacia{10, 20, 30, 40, 50, 60, 70, 80, 90}, "1 1 1 1 1 1 1 1 1"},
		{90, Kombinacia{9, 19, 29, 39, 49, 59, 69, 79, 89}, "1 1 1 1 1 1 1 1 1"},
		{90, Kombinacia{11, 21, 31, 41, 51, 61, 71, 81}, "0 1 1 1 1 1 1 1 1"},
	}
	for _, test := range tests {
		tica := Xtica(test.m, test.t)
		if tica.String() != test.w {
			t.Fatalf("Excepted: (%s), Have: (%s)", test.w, tica)
		}
	}
}

func TestFilterXtica(t *testing.T) {
	tests := []struct {
		k Kombinacia
		f Filter
		w bool
	}{
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterXtica(5, 35, Tica{5, 0, 0, 0}), true},
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterXtica(5, 35, Tica{4, 1, 0, 0}), false},
		{Kombinacia{1, 2, 3, 4, 11}, NewFilterXtica(5, 35, Tica{4, 1, 0, 0}), true},

		{Kombinacia{1}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), true},
		{Kombinacia{1, 2}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), false},
		{Kombinacia{1, 11}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), true},
		{Kombinacia{1, 11, 12}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), false},
		{Kombinacia{1, 11, 20}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), false},
		{Kombinacia{1, 11, 21}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), true},
		{Kombinacia{1, 11, 21, 22}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), false},
		{Kombinacia{1, 11, 21, 30}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), false},
		{Kombinacia{1, 11, 21, 31}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), true},
		{Kombinacia{1, 11, 21, 31}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), true},
		{Kombinacia{1, 11, 21, 31, 32}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), true},
		{Kombinacia{1, 11, 21, 31, 35}, NewFilterXtica(5, 35, Tica{1, 1, 1, 2}), true},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
