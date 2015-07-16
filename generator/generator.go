package generator

import (
	"fmt"

	"github.com/melias122/psl/filter"
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
		numbers     = make([]int, g.m)
		combination = make([]int, 0, g.n)
	)
	for i := range numbers {
		numbers[i] = i + 1
	}
	g.generate(numbers, combination)
}

func (g *Generator) check(combination []int) bool {
	for _, filter := range g.filters {
		if !filter.Check(combination) {
			return false
		}
	}
	return true
}

func (g *Generator) generate(numbers, combination []int) {
	lenNumbers := len(numbers)
	lenCombination := len(combination)
	for i, number := range numbers {
		if g.n-lenCombination > lenNumbers-i {
			break
		}
		combination = append(combination, number)
		if g.check(combination) {
			if len(combination) == g.n {
				// Found
				fmt.Println(combination)
			} else {
				g.generate(numbers[i+1:], combination)
			}
		}
		// _ = combination[lenCombination]
		combination = combination[:lenCombination]
	}
}
