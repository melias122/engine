package psl

import (
	"fmt"
	"testing"
	"time"
)

// 1. BenchmarkGenerator-4 	   30000	     49036 ns/op	    1264 B/op	      53 allocs/op  // old
// 2. BenchmarkKombinator-4	   30000	     53144 ns/op	    1536 B/op	      48 allocs/op 	// new
func BenchmarkKombinator(b *testing.B) {
	archiv, err := NewArchiv("profile/412.csv", "-", 4, 12)
	if err != nil {
		b.Fatal(err)
	}
	var (
		k     = kombinator{}
		cisla = cisla(archiv.Hrx.Cisla, archiv.Skupiny[2].Xcisla)
		n     = 4
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := k.run(cisla, nil, n)
		for range ch {
		}
	}
	b.ReportAllocs()
}

func TestKombinator(t *testing.T) {
	archiv, err := NewArchiv("profile/412.csv", "-", 4, 12)
	if err != nil {
		t.Fatal(err)
	}
	var (
		k = kombinator{}

		cisla = cisla(archiv.Hrx.Cisla, archiv.Skupiny[0].Xcisla)
		n     = archiv.n
	)
	ch := k.run(cisla, nil, n)
	for range ch {
	}
}

func TestGenerator2(t *testing.T) {
	t.SkipNow()
	archiv, err := NewArchiv("profile/590.csv", "-", 5, 90)
	if err != nil {
		t.Fatal(err)
	}
	filters := Filters{
		NewFilterSucet(archiv.n, 100, 130),
	}
	g := NewGenerator2(archiv, filters)
	g.Start()
	// go func() {
	for {
		str, ok := g.Progress()
		fmt.Println(str)
		if !ok {
			return
		}
		time.Sleep(250 * time.Millisecond)
	}
	// }()
	// g.Stop()
	// g.Wait()
	// time.Sleep(250 * time.Millisecond)
}
