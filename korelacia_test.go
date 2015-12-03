package psl

import (
	"strconv"
	"testing"
)

func TestKorelacia(t *testing.T) {
	tests := []struct {
		n, m   int
		k0, k1 Kombinacia
		w      string
	}{
		{5, 35, Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 2, 3, 4, 5}, "1.00000000"},
		{5, 35, Kombinacia{2, 7, 13, 32, 35}, Kombinacia{}, "0.00000000"},
		{5, 35, Kombinacia{1, 14, 15, 17, 19}, Kombinacia{2, 7, 13, 32, 35}, "0.34137300"},
		{5, 35, Kombinacia{2, 7, 13, 32, 35}, Kombinacia{1, 14, 15, 17, 19}, "0.34137300"},
		{5, 35, Kombinacia{4, 9, 10, 25, 27}, Kombinacia{1, 14, 15, 17, 19}, "0.74810803"},
		{5, 35, Kombinacia{1, 2, 13, 21, 31}, Kombinacia{4, 9, 10, 25, 27}, "0.84906826"},
		{5, 35, Kombinacia{17, 21, 29, 32, 34}, Kombinacia{1, 2, 13, 21, 31}, "0.18197335"},
	}
	for _, test := range tests {
		korelacia := strconv.FormatFloat(Korelacia(test.k0, test.k1, test.n, test.m), 'f', 8, 64)
		if korelacia != test.w {
			t.Fatalf("Excepted: (%s), Have: (%s)", test.w, korelacia)
		}
	}
}

func TestFilterKorelacia(t *testing.T) {
	tests := []struct {
		k Kombinacia
		f Filter
		w bool
	}{
		{Kombinacia{1, 14, 15, 17, 19}, NewFilterKorelacia(0.0, 1.0, Kombinacia{2, 7, 13, 32, 35}, 5, 35), true},   // "0.34137300"
		{Kombinacia{1, 14, 15, 17, 19}, NewFilterKorelacia(0.0, 0.34, Kombinacia{2, 7, 13, 32, 35}, 5, 35), true},  // "0.34137300"
		{Kombinacia{1, 14, 15, 17, 19}, NewFilterKorelacia(0.41, 1.0, Kombinacia{2, 7, 13, 32, 35}, 5, 35), false}, // "0.34137300"
		{Kombinacia{1, 2, 3, 4, 5}, NewFilterKorelacia(1.0, 1.0, Kombinacia{1, 2, 3, 4, 5}, 5, 35), true},
		{Kombinacia{1, 2, 3, 4}, NewFilterKorelacia(1.0, 1.0, Kombinacia{1, 2, 3, 4}, 5, 35), true},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) %v", test.w, ok, test)
		}
	}
}
