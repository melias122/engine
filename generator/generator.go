package generator

import (
	"fmt"

	"github.com/melias122/psl/filter"
	"github.com/melias122/psl/komb"
)

type Generator struct {
	n, m    int
	filters []filter.Filter
}

func NewGenerator(n, m int) *Generator {
	generator := &Generator{
		n: n,
		m: m,
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
