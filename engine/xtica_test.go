package engine

import (
	"testing"
)

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
		tica := NewXtica(test.t, test.m)
		if tica.String() != test.w {
			t.Fatalf("Excepted: (%s), Have: (%s)", test.w, tica)
		}
	}
}
