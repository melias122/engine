package komb

import (
	"bytes"
	"testing"
)

func TestCifrovackyTeorMax(t *testing.T) {
	tmax := CifrovackyTeorMax(5, 35)
	exp := Cifrovacky{4, 4, 4, 4, 4, 3, 3, 3, 3, 3}
	if bytes.Compare(tmax[:], exp[:]) != 0 {
		t.Errorf("Expected %s, got: %s", exp, tmax)
	}
}

func TestMakeCifrovacky(t *testing.T) {
	tests := []struct {
		k Kombinacia
		w Cifrovacky
	}{
		{k: Kombinacia{}, w: Cifrovacky{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{k: Kombinacia{1, 2, 3, 4, 5}, w: Cifrovacky{1, 1, 1, 1, 1, 0, 0, 0, 0, 0}},
		{k: Kombinacia{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, w: Cifrovacky{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{k: Kombinacia{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, w: Cifrovacky{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
	}
	for _, test := range tests {
		c := MakeCifrovacky(test.k)
		if bytes.Compare(c[:], test.w[:]) != 0 {
			t.Errorf("Expected %s, got: %s", test.w, c)
		}
	}
}
