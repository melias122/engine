package engine

import "testing"

func BenchmarkTicaString(b *testing.B) {
	t := make(Tica, 30)
	for i := 0; i < b.N; i++ {
		t.String()
	}
}

func TestTicaString(t *testing.T) {
	tests := []struct {
		t Tica
		s string
	}{
		{t: Tica{}, s: ""},
		{t: Tica{1}, s: "1"},
		{t: Tica{1, 2}, s: "1 2"},
		{t: Tica{0, 1, 0}, s: "0 1 0"},
	}
	for _, test := range tests {
		if test.t.String() != test.s {
			t.Logf("Expected: '%s' got : '%s'", test.s, test.t.String())
		}
	}
}
