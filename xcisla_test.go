package engine

import (
	"testing"
)

// BenchmarkXcislaIndex-4	20000000	        83.0 ns/op
func BenchmarkXcislaIndex(b *testing.B) {
	xcisla := NewXcisla(90)
	for i := 0; i < 90; i++ {
		xcisla.move(1, 0, i+1)
	}
	for i := 0; i < b.N; i++ {
		xcisla.index(i % 90)
	}
}

func BenchmarkXcislaMove(b *testing.B) {
	xcisla := NewXcisla(90)
	for i := 0; i < b.N; i++ {
		xcisla.move(1, i, i+1)
	}
}

var xcislaTest = []Skupina{
	{Xcisla: Xcisla{{1, 5}}},
	{Xcisla: Xcisla{{1, 1}, {2, 1}, {3, 1}, {5, 2}}},
	{Xcisla: Xcisla{{1, 1}, {2, 1}, {3, 1}, {6, 2}}},
	{Xcisla: Xcisla{{1, 1}, {2, 1}, {3, 1}, {7, 2}}},
	{Xcisla: Xcisla{{1, 1}, {2, 1}, {3, 1}, {8, 2}}},
	{Xcisla: Xcisla{{1, 1}, {2, 2}, {3, 1}}},
	{Xcisla: Xcisla{{1, 2}, {2, 2}, {3, 1}}},
	{Xcisla: Xcisla{{1, 2}, {2, 1}, {3, 1}}},
	{Xcisla: Xcisla{{1, 2}, {2, 1}, {3, 1}, {10, 1}}},
}

func TestFilterXcisla(t *testing.T) {
	tabs := Xcisla{
		{1, 1},
		{2, 2},
		{2, 1},
		{1, 2},
	}
	filter := NewFilterXcisla(tabs)
	for _, s := range xcislaTest {
		if ok := filter.CheckSkupina(s); ok {
			t.Log(s.Xcisla)
		}
	}
}
