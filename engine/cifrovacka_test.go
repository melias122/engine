package engine

import (
	"testing"
)

func TestCifrovackaTeorMax(t *testing.T) {
	tests := []struct {
		n, m int
		w    Cifrovacka
	}{
		{5, 35, Cifrovacka{c: [10]byte{3, 4, 4, 4, 4, 4, 3, 3, 3, 3}}},
		{5, 90, Cifrovacka{c: [10]byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}}},
	}
	for _, test := range tests {
		got := NewCifrovackaMax(test.n, test.m)
		if test.w != *got {
			t.Errorf("Expected %v, got: %v", test.w, got)
		}
	}
}

func TestNewCifrovacka(t *testing.T) {
	tests := []struct {
		k Kombinacia
		w Cifrovacka
	}{
		{k: Kombinacia{}, w: Cifrovacka{c: [10]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}},
		{k: Kombinacia{10}, w: Cifrovacka{c: [10]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}}},
		{k: Kombinacia{1}, w: Cifrovacka{c: [10]byte{0, 1, 0, 0, 0, 0, 0, 0, 0, 0}}},
		{k: Kombinacia{9}, w: Cifrovacka{c: [10]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}}},
		{k: Kombinacia{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 30}, w: Cifrovacka{c: [10]byte{2, 1, 1, 1, 1, 1, 1, 1, 1, 1}}},
	}
	for _, test := range tests {
		c := NewCifrovacka(test.k)
		if *c != test.w {
			t.Errorf("Expected %v, got: %v", test.w, c)
		}
	}
}
