package engine

import (
	"testing"
)

func TestCifrovackaTeorMax(t *testing.T) {
	tests := []struct {
		n, m int
		w    Cifrovacka
	}{
		{5, 35, Cifrovacka{4, 4, 4, 4, 4, 3, 3, 3, 3, 3}},
		{5, 90, Cifrovacka{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}},
	}
	for _, test := range tests {
		got := NewCifrovackaMax(test.n, test.m)
		if test.w != got {
			t.Errorf("Expected %v, got: %v", test.w, got)
		}
	}
}

func TestNewCifrovacka(t *testing.T) {
	tests := []struct {
		k Kombinacia
		w Cifrovacka
	}{
		{k: Kombinacia{}, w: Cifrovacka{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{k: Kombinacia{1, 2, 3, 4, 5}, w: Cifrovacka{1, 1, 1, 1, 1, 0, 0, 0, 0, 0}},
		{k: Kombinacia{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, w: Cifrovacka{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{k: Kombinacia{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, w: Cifrovacka{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
	}
	for _, test := range tests {
		c := NewCifrovacka(test.k)
		if c != test.w {
			t.Errorf("Expected %v, got: %v", test.w, c)
		}
	}
}
