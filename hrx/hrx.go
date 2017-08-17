package hrx

import (
	"github.com/melias122/engine/engine"
)

// Maximalny pocet riadkov nez nastane udalost 101.
const max101 = 19

type Hrx struct {
	*Cislo

	n, m  int
	cache *hCache
	Uc    engine.Uc
}

func NewHrx(n, m int) *Hrx {
	cache := newHCache(n, m)
	return &Hrx{
		Cislo: newCislo(cache, n, m),
		n:     n,
		m:     m,
		cache: cache,
	}
}

func (h *Hrx) Add(kombs []engine.Kombinacia) {
	var (
		is101 = make(map[int]bool)
	)
	for _, k := range kombs {
		for _, c := range k {
			is101[c] = true
		}
		if len(is101) == h.m {
			break
		}
	}

	// Ak sa v hhrx vyskytli vsetky cisla 1..m
	// nastala udalost 101
	if len(is101) == h.m {

		var reverse bool
		if h.Uc.Cislo == 0 {
			reverse = true
		} else {
			last := kombs[len(kombs)-1]
			for _, num := range last {
				// Ak na riadku narazime na Uc Cislo
				// porebujeme ho spatne dohladat
				if num == h.Uc.Cislo {
					reverse = true
				}
			}
			if !reverse {
				h.Cislo.add(last)
			}
		}

		// Uc cislo je 0 len raz po udalosti 101.
		// Spatne dohladanie Uc cisla a riadku a inrementovanie cisiel Roddo
		if reverse {
			// Nova hrx zostava a resetovanie cisiel Roddo
			h.Cislo = newCislo(h.cache, h.n, h.m)
			is101 := make(map[int]bool)
			newUc := engine.Uc{Riadok: len(kombs)}
			// Spatne nacitava kombinacie a incremtuje Roddo
			// a Hrx az pokial nenastane udalost 101
			// udalost 101 nastava ked sa kazde cislo vyskytne aspon 1

			for len(is101) < h.m {
				newUc.Riadok--
				k := kombs[newUc.Riadok]
				h.Cislo.add(k)
				for _, num := range k {
					is101[num] = true
					if len(is101) == h.m {
						newUc.Cislo = num
						break
					}
				}
			}
			// Nastavenie noveho Uc cisla a riadku pre archiv
			h.Uc = newUc
		}
	}
}

func (h *Hrx) X(k engine.Kombinacia) float64 {
	return h.Cislo.x(max101, k)
}
