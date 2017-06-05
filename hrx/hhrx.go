package hrx

import "github.com/melias122/engine/engine"

type HHrx struct {
	*Cislo

	n, m int
}

func NewHHrx(n, m int) *HHrx {
	return &HHrx{
		Cislo: newCislo(nil, n, m),
		n:     n,
		m:     m,
	}
}

func (h *HHrx) Add(kombs []engine.Kombinacia) {
	for _, k := range kombs {
		h.Cislo.add(k)
	}
}

func (h *HHrx) X(k engine.Kombinacia) float64 {
	max := h.Cislo.max
	for _, c := range k {
		i := h.Cislo.Rp(c)
		if i > max {
			max = i
		}
	}
	return h.Cislo.x(max, k)
}
