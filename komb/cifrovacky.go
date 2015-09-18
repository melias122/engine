package komb

import "strconv"

type Cifrovacky [10]byte

func CifrovackyTeorMax(n, m int) Cifrovacky {
	var c Cifrovacky
	for i := 1; i <= m; i++ {
		j := i % 10
		if j == 0 && c[9] <= byte(n) {
			c[9]++
		} else if c[j-1] <= byte(n) {
			c[j-1]++
		}
	}
	return c
}

func MakeCifrovacky(k Kombinacia) Cifrovacky {
	var c Cifrovacky
	for i := range k {
		j := k[i] % 10 // cifra
		switch j {
		case 0:
			c[9]++
		default:
			c[j-1]++
		}
	}
	return c
}

func (c Cifrovacky) Strings() []string {
	s := make([]string, len(c))
	for i, c := range c {
		s[i] = strconv.Itoa(int(c))
	}
	return s
}

func (k Kombinacia) Cifrovacky() Cifrovacky {
	return MakeCifrovacky(k)
}
