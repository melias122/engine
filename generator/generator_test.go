package generator

import (
	"testing"

	"github.com/melias122/psl/filter"
	"github.com/melias122/psl/hrx"
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

	Hrx := hrx.New(n, m)
	for _, i := range k_hrx {
		for y, x := range i {
			Hrx.Add(x, y)
		}
	}

	HHrx := hrx.New(n, m)
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

// func TestFilter(t *testing.T) {
// 	n, _, Hrx, HHrx, presuny := Get412()
// 	filters := filter.Filters{
// 		// filter.NewZakazane([]byte{1, 3, 5, 7, 11}),
// 		// filter.NewSucet(n, 22, 38),
// 	}
// 	GenerateFilter(n, , filters filter.Filters)
// 	// for i := 0; i < b.N; i++ {
// 		// xGenerate(n, Hrx.Cisla, presuny, filters)
// 	// }
// }

// func TestGenerator(t *testing.T) {
// 	n, _, Hrx, _, presuny := Get412()
// 	// riadok := archiv.Riadok{
// 	// 	K: komb.Kombinacia{1, 2, 3, 4},
// 	// }
// 	filters := filter.Filters{
// 	// filter.NewSucet(n, 22, 38),
// 	// filter.NewZakazane([]byte{3}),
// 	}
//
// 	// vystup := NewV1(n, m, Hrx, HHrx, riadok)
// }

// func xGenerate(n int, HrxCisla num.Nums, presuny []hrx.Presun, filters filter.Filters) {
// 	var (
// 		wg             sync.WaitGroup
// 		chanKombinacie = make(chan komb.Kombinacia, 16)
// 		chanPresuny    = make(chan hrx.Presun, 2)
// 	)
//
// 	go func(ch chan hrx.Presun) {
// 		defer close(ch)
// 		for _, p := range presuny {
// 			chanPresuny <- p
// 		}
// 	}(chanPresuny)
//
// 	for i := 0; i < runtime.NumCPU(); i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			g := NewGenerator(n, HrxCisla, chanKombinacie, filters)
// 			for p := range chanPresuny {
// 				g.Generate(p)
// 			}
// 		}()
// 	}
// 	go func() {
// 		defer close(chanKombinacie)
// 		wg.Wait()
// 	}()
//
// 	// var cnt int
// 	for range chanKombinacie {
// 		// cnt++
// 	}
// 	// if cnt != 495 {
// 	// 	panic("")
// 	// }
// }

func BenchmarkGenerator(b *testing.B) {
	n, _, Hrx, _, presuny := Get412()
	filters := filter.Filters{
		filter.NewZakazane([]byte{1, 3, 5, 7, 11}),
		filter.NewSucet(n, 22, 38),
	}
	for i := 0; i < b.N; i++ {
		xGenerate(n, Hrx.Cisla, presuny, filters)
	}
}
