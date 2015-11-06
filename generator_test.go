package psl

import (
	"testing"

	// "github.com/melias122/psl/filter"
	// "github.com/melias122/psl/hrx"
	// "github.com/melias122/psl/komb"
)

func Get412() (int, int, *hrx.H, *hrx.H, []hrx.Presun) {
	const n = 4
	const m = 12
	k_hrx := [][]int{
		[]int{1, 3, 5, 7},
		[]int{2, 4, 6, 8},
		[]int{3, 6, 9, 12},
		[]int{4, 8, 10, 12},
		[]int{5, 9, 10, 11},
		[]int{6, 7, 8, 12},
	}
	k_hhrx := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{1, 3, 5, 7},
		[]int{2, 4, 6, 8},
		[]int{3, 6, 9, 12},
		[]int{4, 8, 10, 12},
		[]int{5, 9, 10, 11},
		[]int{6, 7, 8, 12},
	}

	Hrx := hrx.NewHrx(n, m)
	for _, i := range k_hrx {
		for y, x := range i {
			Hrx.Add(x, y)
		}
	}

	HHrx := hrx.NewHHrx(n, m)
	for _, i := range k_hhrx {
		for y, x := range i {
			HHrx.Add(x, y)
		}
	}
	presuny := []hrx.Presun{
		hrx.Presun{hrx.Tab{1, 3}, hrx.Tab{2, 1}},                // 0 3 1
		hrx.Presun{hrx.Tab{1, 3}, hrx.Tab{3, 1}},                // 0 3 0 1
		hrx.Presun{hrx.Tab{1, 2}, hrx.Tab{2, 2}},                // 0 2 2
		hrx.Presun{hrx.Tab{1, 2}, hrx.Tab{2, 1}, hrx.Tab{3, 1}}, // 0 2 1 1
		hrx.Presun{hrx.Tab{1, 2}, hrx.Tab{3, 2}},                // 0 2 0 2
		hrx.Presun{hrx.Tab{1, 1}, hrx.Tab{2, 3}},                // 0 1 3
		hrx.Presun{hrx.Tab{1, 1}, hrx.Tab{2, 2}, hrx.Tab{3, 1}}, // 0 1 2 1
		hrx.Presun{hrx.Tab{1, 1}, hrx.Tab{2, 1}, hrx.Tab{3, 2}}, // 0 1 1 2
		hrx.Presun{hrx.Tab{1, 1}, hrx.Tab{3, 3}},                // 0 1 0 3
		hrx.Presun{hrx.Tab{2, 4}},                               // 0 0 4
		hrx.Presun{hrx.Tab{2, 3}, hrx.Tab{3, 1}},                // 0 0 3 1
		hrx.Presun{hrx.Tab{2, 2}, hrx.Tab{3, 2}},                // 0 0 2 2
		hrx.Presun{hrx.Tab{2, 1}, hrx.Tab{3, 3}},                // 0 0 1 3
	}
	return n, m, Hrx, HHrx, presuny
}

// BenchmarkGenerator-4	  200000	      8942 ns/op	     640 B/op	      14 allocs/op
func BenchmarkGenerator(b *testing.B) {
	n, _, Hrx, _, presuny := Get412()
	filters := filter.Filters{}
	ch := make(chan komb.Kombinacia)
	Generator := NewGenerator(n, Hrx.Cisla, ch, filters)
	go func() {
		for range ch {
		}
	}()
	for i := 0; i < b.N; i++ {
		Generator.Generate(presuny[0])
	}
}
