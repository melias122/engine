package engine

import (
	"strconv"
	"testing"
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
		smernica := strconv.FormatFloat(Smernica(test.k, test.n, test.m), 'f', 8, 64)
		if smernica != test.w {
			t.Fatalf("Excepted: (%s), Have: (%s)", test.w, smernica)
		}
	}
}
