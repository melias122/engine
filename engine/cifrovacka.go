//go:generate stringer -type Cif -trimprefix Cif
package engine // import "github.com/melias122/engine/engine"

type Cif byte

const (
	CifC0 Cif = iota
	CifC1
	CifC2
	CifC3
	CifC4
	CifC5
	CifC6
	CifC7
	CifC8
	CifC9
)

type Cifrovacka struct {
	c [10]byte
}

func (c *Cifrovacka) Get(i Cif) int {
	return int(c.c[i])
}

func NewCifrovacka(k Kombinacia) *Cifrovacka {
	var c Cifrovacka
	for _, n := range k {
		c.c[n%10]++
	}
	return &c
}

func NewCifrovackaMax(n, m int) *Cifrovacka {
	var c Cifrovacka
	for i := 1; i <= m; i++ {
		c.c[i%10]++
	}
	return &c
}
