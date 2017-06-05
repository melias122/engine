package engine

type Cifrovacka struct {
	C1 byte
	C2 byte
	C3 byte
	C4 byte
	C5 byte
	C6 byte
	C7 byte
	C8 byte
	C9 byte
	C0 byte
}

func NewCifrovacka(k Kombinacia) Cifrovacka {
	var c Cifrovacka
	for _, n := range k {
		c.set(n)
	}
	return c
}

func NewCifrovackaMax(n, m int) Cifrovacka {
	var c Cifrovacka
	for i := 1; i <= m; i++ {
		c.set(i)
	}
	return c
}

func (c *Cifrovacka) set(n int) {
	switch n % 10 {
	case 1:
		c.C1++
	case 2:
		c.C2++
	case 3:
		c.C3++
	case 4:
		c.C4++
	case 5:
		c.C5++
	case 6:
		c.C6++
	case 7:
		c.C7++
	case 8:
		c.C8++
	case 9:
		c.C9++
	case 0:
		c.C0++
	}
}
