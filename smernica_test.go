package psl

import (
	"strconv"
	"testing"

	// "github.com/melias122/psl/komb"
)

func TestSmernica(t *testing.T) {
	tests := []struct {
		n, m int
		k    Kombinacia
		w    string
	}{
		{5, 35, Kombinacia{8, 15, 20, 24, 35}, "0.75000000"},
		{5, 35, Kombinacia{14, 16, 26, 27, 30}, "0.50000000"},
		{5, 35, Kombinacia{1, 4, 6, 26, 33}, "1.00000000"},
		{5, 35, Kombinacia{2, 7, 13, 32, 35}, "1.05392157"},
		{5, 35, Kombinacia{1, 14, 15, 17, 19}, "0.47058824"},
		{5, 35, Kombinacia{4, 9, 10, 25, 27}, "0.72058824"},
		{5, 35, Kombinacia{1, 2, 13, 21, 31}, "0.92156863"},
		{5, 35, Kombinacia{17, 21, 29, 32, 34}, "0.52450980"},
	}
	for _, test := range tests {
		smernica := strconv.FormatFloat(Smernica(test.n, test.m, test.k), 'f', 8, 64)
		if smernica != test.w {
			t.Fatalf("Excepted: (%s), Have: (%s)", test.w, smernica)
		}
	}
}

func TestSmernica(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{14, 16, 26, 27}, NewSmernica(5, 35, 0.0, 1.0), true},         // 0.499...
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.0, 1.0), true},     // 0.499...
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.0, 0.49), true},    // 0.499...
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.49, 0.5), true},    // 0.499...
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.499, 0.499), true}, // 0.499...
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.5, 0.5), true},
		{komb.Kombinacia{14, 16, 26, 27, 30}, NewSmernica(5, 35, 0.500000001, 0.6), false}, // 0.499...
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) (test %v)", test.w, ok, test)
		}
	}
}
