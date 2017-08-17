package hrx

import (
	"sort"

	"github.com/melias122/engine/engine"
)

// Generator generates all possible combinations of current hrx.
type Generator struct {
	n      int
	xcisla engine.Xcisla
	next   engine.Xcisla
}

// NewGenerator creates Generator for hrx.
func NewGenerator(hrx engine.Rc, n, m int) *Generator {

	group := make(map[int]int)
	for c := 1; c <= m; c++ {
		group[hrx.Rp(c)]++
	}

	idx := make([]int, 0, len(group))
	for sk := range group {
		idx = append(idx, sk)
	}
	sort.Ints(idx)

	xcisla := make(engine.Xcisla, 0, len(idx))
	for _, sk := range idx {
		xcisla = append(xcisla, engine.Tab{
			Sk: sk,

			// Maximalny pocet cisiel, ktore mozu prejst zo skupiny
			Max: min(n, group[sk]),
		})
	}

	first := make(engine.Xcisla, 0, len(xcisla))
	for left := n; left > 0; {
		t := xcisla[len(first)]
		t.Max = min(left, t.Max)
		first = append(first, t)
		left -= t.Max
	}

	return &Generator{
		n:      n,
		xcisla: xcisla,
		next:   first,
	}
}

// Next generate next combination and returns
// returns true or false whenether there is more
// combinations.
func (g *Generator) Next() bool {

	var (
		i, left int
		tab     *engine.Tab
	)

again:
	for {
		if len(g.next) == 0 {
			return false
		}

		tab = &g.next[len(g.next)-1]
		tab.Max--
		left++

		if tab.Max == 0 {
			g.next = g.next[:len(g.next)-1]
		}

		if tab.Sk != g.xcisla[len(g.xcisla)-1].Sk {
			break
		}
	}

	// TODO(m): This loop should be removed if Tab is wrapped
	// with position index
	i = len(g.xcisla) - 1
	for ; i >= 0; i-- {
		if g.xcisla[i].Sk == tab.Sk {
			break
		}
	}

	i++
	for ; left > 0 && i < len(g.xcisla); i++ {
		t := g.xcisla[i]
		t.Max = min(left, t.Max)
		g.next = append(g.next, t)
		left -= t.Max
	}

	if left > 0 {
		goto again
	}

	return true
}

// Xcisla returns current combination of Xcisla.
// It must not be modified.
func (g *Generator) Xcisla() engine.Xcisla {
	return g.next
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
