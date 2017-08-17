package engine

import (
	"testing"
)

// BenchmarkXcislaIndex-4	20000000	        83.0 ns/op
func BenchmarkXcislaIndex(b *testing.B) {
	xcisla := NewXcisla(90)
	for i := 0; i < 90; i++ {
		xcisla.Move(1, 0, i+1)
	}
	for i := 0; i < b.N; i++ {
		xcisla.index(i % 90)
	}
}

func BenchmarkXcislaMove(b *testing.B) {
	xcisla := NewXcisla(90)
	for i := 0; i < b.N; i++ {
		xcisla.Move(1, i, i+1)
	}
}

func TestXcislaContains(t *testing.T) {
	tests := []struct {
		xcisla Xcisla
		test   map[Tab]bool
	}{
		{xcisla: Xcisla{}, test: map[Tab]bool{Tab{}: false, Tab{1, 1}: false}},
		{xcisla: Xcisla{{0, 0}}, test: map[Tab]bool{Tab{0, 0}: true, Tab{1, 1}: false}},
		{
			xcisla: Xcisla{
				{1, 1}, {2, 2}, {4, 0},
			},
			test: map[Tab]bool{
				Tab{0, 0}: false, Tab{1, 1}: true, Tab{2, 2}: true, Tab{3, 3}: false,
				Tab{3, 0}: false, Tab{3, 1}: false, Tab{4, 0}: true, Tab{4, 1}: false,
				Tab{5, 1}: false,
			},
		},
	}
	for _, test := range tests {
		for tab, exp := range test.test {
			if ok := test.xcisla.Contains(tab); ok != exp {
				t.Fatalf("Expected: (%v:%v), got: (%v:%v)", tab, exp, test.xcisla, ok)
			}
		}
	}
}
