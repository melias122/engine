package generator

import (
	"fmt"
	"strconv"

	"github.com/melias122/psl/filter"
	"github.com/melias122/psl/komb"
)

type Generator struct {
	n, m    int
	filters []filter.Filter
}

func NewGenerator(n, m int, filters []filter.Filter) *Generator {
	generator := &Generator{
		n:       n,
		m:       m,
		filters: filters,
	}
	return generator
}

func (g *Generator) Generate() {
	var (
		cisla      = make(komb.Kombinacia, g.m)
		kombinacia = make(komb.Kombinacia, 0, g.n)
	)
	for i := range cisla {
		cisla[i] = byte(i + 1)
	}
	g.generate(cisla, kombinacia)
}

func (g *Generator) check(k komb.Kombinacia) bool {
	for _, filter := range g.filters {
		if !filter.Check(k) {
			return false
		}
	}
	return true
}

func (g *Generator) generate(cisla, kombinacia komb.Kombinacia) {
	lenCisla := len(cisla)
	lenKombinacia := len(kombinacia)
	for i, cislo := range cisla {
		if g.n-lenKombinacia > lenCisla-i {
			break
		}
		kombinacia = append(kombinacia, cislo)
		if g.check(kombinacia) {
			if len(kombinacia) == g.n {
				// Found
				fmt.Println(kombinacia)
			} else {
				g.generate(cisla[i+1:], kombinacia)
			}
		}
		// _ = kombinacia[lenCombination]
		kombinacia = kombinacia[:lenKombinacia]
	}
}

func header(n int) []string {
	var h []string
	for i := 1; i <= n; i++ {
		h = append(h, strconv.Itoa(i))
	}
	h = append(h, []string{
		"P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc", "CC", "ZH",
		"Sm", "Kk", "N-tice", "X-tice",
		"ƩR1-DO", "ƩR1-DO", "ƩSTL1-DO", "\u0394ƩSTL1-DO", "\u0394(ƩR1-DO-ƩSTL1-DO)",
		"HHRX", "\u0394HHRX",
		"ƩR OD-DO", "\u0394ƩR OD-DO", "ƩSTL OD-DO", "\u0394ƩSTL OD-DO", "\u0394(ƩROD-DO-ƩSTLOD-DO)",
		"HRX", "\u0394HRX",
		"ƩKombinacie",
	}...,
	)
	return h
}
