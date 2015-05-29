package math

import (
	"fmt"
	"testing"
)

func BenchmarkValue(b *testing.B) {
	n, m := 30, 90
	r := 0
	for i := 0; i < b.N; i++ {
		p := r%100 + 1
		x := r%n + 1
		y := r%m + 1
		r++
		Value(p, x, y, n, m)
	}
}

func TestValue(t *testing.T) {
	tests := []struct {
		p, x, y, n, m int
		w             string
	}{
		{0, 0, 0, 0, 0, "0.0000000000"},
		{122, 1, 1, 5, 35, "0.0026306710"},
		{122, 1, 2, 5, 35, "0.0000000000"},
		{22, 13, 1, 5, 35, "0.0030075188"},
		{60, 13, 2, 5, 35, "0.0032467532"},
		{39, 13, 3, 5, 35, "0.0025580480"},
		{12, 13, 4, 5, 35, "0.0024793388"},
		{2, 13, 5, 5, 35, "0.0040404040"},
		{43, 35, 1, 5, 35, "0.0000000000"},
	}
	for _, x := range tests {
		r := fmt.Sprintf("%.10f", Value(x.p, x.x, x.y, x.n, x.m))
		if r != x.w {
			t.Fatalf("Expected: (%s), Have: (%s)", x.w, r)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		x, y, n, m int
		w          string
	}{
		{1, 1, 5, 35, "46376"},
		{1, 2, 5, 35, "0"},
		{5, 1, 5, 35, "27405"},
		{5, 2, 5, 35, "16240"},
		{5, 3, 5, 35, "2610"},
		{5, 4, 5, 35, "120"},
		{5, 5, 5, 35, "1"},
	}
	for _, x := range tests {
		r := Max(x.x, x.y, x.n, x.m).String()
		if r != x.w {
			t.Fatalf("Expected: (%s), Have: (%s)", x.w, r)
		}
	}
}
