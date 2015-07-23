package num

import (
	"fmt"
	"testing"
)

func BenchmarkValue(b *testing.B) {
	k := key{1, 1, 30, 90, 1500}
	for i := 0; i < b.N; i++ {
		value(k)
	}
	b.ReportAllocs()
}

func makekey(x, y, n, m, p int) key {
	k := newph(x, y, n, m).key
	k.pocet = uint32(p)
	return k
}

func TestValue(t *testing.T) {
	tests := []struct {
		k key
		w string
	}{
		{makekey(0, 0, 0, 0, 0), "0.0000000000"},
		{makekey(1, 1, 5, 35, 122), "0.0026306710"},
		{makekey(1, 2, 5, 35, 122), "0.0000000000"},
		{makekey(13, 1, 5, 35, 22), "0.0030075188"},
		{makekey(13, 2, 5, 35, 60), "0.0032467532"},
		{makekey(13, 3, 5, 35, 39), "0.0025580480"},
		{makekey(13, 4, 5, 35, 12), "0.0024793388"},
		{makekey(13, 5, 5, 35, 2), "0.0040404040"},
		{makekey(35, 1, 5, 35, 43), "0.0000000000"},
	}
	for _, x := range tests {
		r := fmt.Sprintf("%.10f", value(x.k))
		if r != x.w {
			t.Fatalf("Expected: (%s), Have: (%s)", x.w, r)
		}
	}
}
