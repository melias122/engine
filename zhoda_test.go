package engine

import (
	"testing"
)

func TestZhoda(t *testing.T) {
	tests := []struct {
		k1, k2 Kombinacia
		zhoda  int
	}{
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 2, 3, 4, 5}, 5},
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 2, 3, 4, 6}, 4},
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 2, 3, 9, 10}, 3},
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 2, 8, 9, 10}, 2},
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{1, 7, 8, 9, 10}, 1},
		{Kombinacia{1, 2, 3, 4, 5}, Kombinacia{6, 7, 8, 9, 10}, 0},
	}
	for _, test := range tests {
		zhoda := Zhoda(test.k1, test.k2)
		if zhoda != test.zhoda {
			t.Fatalf("Excepted: (%d), Have: (%d)", test.zhoda, zhoda)
		}
	}
}

// BenchmarkZhoda-4	20000000 59.6 ns/op
func BenchmarkZhoda(b *testing.B) {
	k := make(Kombinacia, 30)
	for i := range k {
		k[i] = byte(i + 1)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Zhoda(k, k)
	}
}
