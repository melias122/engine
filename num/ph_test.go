package num

import (
	"fmt"
	"testing"
)

func TestInc(t *testing.T) {
	ph := newph(1, 1, 5, 35)
	for i := 1; i <= 122; i++ {
		ph.inc()
	}

	r := fmt.Sprintf("%d", ph.p)
	w := "122"
	if r != w {
		t.Fatalf("Expected: (%s), Have: (%s)", w, r)
	}

	r = fmt.Sprintf("%.10f", ph.h)
	w = "0.0026306710"
	if r != w {
		t.Fatalf("Expected: (%s), Have: (%s)", w, r)
	}
}

func TestReset(t *testing.T) {
	ph := newph(1, 1, 5, 35)
	for i := 1; i <= 10; i++ {
		ph.inc()
	}
	ph.reset()

	r := ph.p
	w := 0
	if r != w {
		t.Fatalf("Expected: (%d), Have: (%d)", w, r)
	}

	r2 := ph.h
	w2 := 0.0
	if r != w {
		t.Fatalf("Expected: (%f), Have: (%f)", w2, r2)
	}
}

func BenchmarkStruct(b *testing.B) {
	m := make(map[key]float64, 256)
	e := key{1, 2, 3}
	for i := 0; i < b.N; i++ {
		m[e] = float64(i)
	}
}

func BenchmarkArr(b *testing.B) {
	m := make(map[[3]int]float64, 256)
	e := [3]int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		m[e] = float64(i)
	}
}

func BenchmarkInt(b *testing.B) {
	m := make(map[int]float64, 256)
	// e := [3]int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		m[0] = float64(i)
	}
}

func BenchmarkValue(b *testing.B) {
	n, m := 30, 90
	k := []key{}
	for i := 0; i < 10; i++ {
		k = append(k, key{i%100 + 1, i%n + 1, i%m + 1})
	}
	b.ResetTimer()
	t := 0
	for i := 0; i < b.N; i++ {
		if t == 10 {
			t = 0
		}
		Value(k[t].p, k[t].x, k[t].y, n, m)
		t++
	}
	b.ReportAllocs()
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
