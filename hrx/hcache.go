package hrx

import "github.com/melias122/engine/engine"

type hkey struct {
	cislo   int16
	pozicia int16
	pocet   int32
}

type hCache struct {
	n, m  int
	cache map[hkey]float64
}

func newHCache(n, m int) *hCache {
	return &hCache{
		n:     n,
		m:     m,
		cache: make(map[hkey]float64, 1<<10),
	}
}

func (h *hCache) H(cislo, pozicia, pocet int) float64 {
	//	if cislo > (h.m/2)+h.m%2 {
	//		cislo = (cislo - (h.m + 1)) * (-1)
	//		pozicia = h.n - pozicia + 1
	//	}
	k := hkey{int16(cislo), int16(pozicia), int32(pocet)}
	v, ok := h.cache[k]
	if ok {
		return v
	}
	v = engine.H(cislo, pozicia, pocet, h.n, h.m)
	h.cache[k] = v
	return v
}
