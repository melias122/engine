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
