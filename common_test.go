package psl

import "testing"

func BenchmarkFtoa(b *testing.B) {
	f := 0.123123123123123123123123123123
	for i := 0; i < b.N; i++ {
		ftoa(f)
	}
}

func TestFtoa(t *testing.T) {
	tests := []struct {
		f float64
		s string
	}{
		{f: 0, s: "0"},
		{f: 1.1, s: "1,1"},
		{f: 10.123123123123123, s: "10,123123123123123"},
	}
	for _, test := range tests {
		s := ftoa(test.f)
		if s != test.s {
			t.Logf("Expected: '%s', got: '%s'", test.s, s)
		}
	}
}
