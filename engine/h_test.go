package engine

import (
	"fmt"
	"testing"
)

func BenchmarkH(b *testing.B) {
	x := 1
	y := 1
	n := 30
	m := 90
	pocet := 1500
	for i := 0; i < b.N; i++ {
		H(x, y, pocet, n, m)
	}
	b.ReportAllocs()
}

func TestH(t *testing.T) {
	tests := []struct {
		x, y, n, m, pocet int
		w                 string
	}{
		{0, 0, 0, 0, 0, "0.0000000000"},
		{1, 1, 5, 35, 122, "0.0026306710"},
		{1, 2, 5, 35, 122, "0.0000000000"},
		{13, 1, 5, 35, 22, "0.0030075188"},
		{13, 2, 5, 35, 60, "0.0032467532"},
		{13, 3, 5, 35, 39, "0.0025580480"},
		{13, 4, 5, 35, 12, "0.0024793388"},
		{13, 5, 5, 35, 2, "0.0040404040"},
		{35, 1, 5, 35, 43, "0.0000000000"},
	}
	for i, x := range tests {
		r := fmt.Sprintf("%.10f", H(x.x, x.y, x.pocet, x.n, x.m))
		if r != x.w {
			t.Errorf("Expected: (%s), Have: (%s) (test %d)", x.w, r, i+1)
		}
	}
}
