package generator

import (
	"sync/atomic"

	"github.com/melias122/engine/engine"
	"github.com/melias122/engine/filter"
)

type task struct {
	n       int
	hrxNums engine.Nums
	xcisla  engine.Xcisla
	filters filter.Filters

	_done uint32
}

func (t *task) Run() error {
	var (
		n       = t.n
		nums    = makeNums(t.hrxNums, t.xcisla)
		filters = t.filters

		indices = make([]int, 1, n)
		k       = make(engine.Kombinacia, 0, n)
	)
	for len(indices) > 0 && !t.done() {
		j := len(indices)

		// i je index daneho cisla
		i := indices[j-1]

		// na tomto leveli uz nie su dalsie cisla
		// ideme o level nizsie
		if i == len(nums) {
			indices = indices[:j-1]
			continue
		}

		// skusime cislo
		num := nums[i]

		// v predchadzajucom kroku sme nasli kombinaciu
		// skusime dalsie cislo na tomto leveli
		if len(k) == j && num.num == k[len(k)-1] {
			k = k[:len(k)-1]
			num.inc()
			indices[j-1]++
			continue
		}

		// ak pocet cisiel z danej hrx skupiny
		// je vacsi ako 0, berieme cislo do kombinacie
		// a znizime pocet cisiel v skupine
		if num.zero() {
			indices[j-1]++
			continue
		}

		k = append(k, num.num)
		num.dec()

		// ak kombinacia nevyhovuje filtru
		// skusime dalsie cislo
		if !filters.Check(k) {
			continue
		}

		// cisel v kombinacii este nie je n
		// skusime dalsie cislo
		if len(k) < n {
			indices = append(indices, i+1)
			continue
		}
	}
	return nil
}

func (t *task) Cancel() {
	atomic.AddUint32(&t._done, 1)
}

func (t *task) done() bool {
	return atomic.LoadUint32(&t._done) > 0
}

func makeNums(hrxNums engine.Nums, xcisla engine.Xcisla) []num {
	var (
		nums  []num
		skMax = make(map[int]*int, xcisla.Len())
	)
	for _, tab := range xcisla {
		total := tab.Max
		skMax[tab.Sk] = &total
	}
	for _, Num := range hrxNums {
		sk := Num.PocetR()
		if total, ok := skMax[sk]; ok {
			nums = append(nums, num{
				num:   Num.Cislo(),
				total: total,
			})
		}
	}
	return nums
}

type num struct {
	num   int
	total *int
}

func (n *num) zero() bool {
	return *n.total == 0
}

func (n *num) inc() {
	*n.total++
}

func (n *num) dec() {
	*n.total--
}
