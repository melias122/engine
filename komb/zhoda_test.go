package komb

import "testing"

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
