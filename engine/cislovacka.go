//go:generate stringer -type Cxx -trimprefix Cxx
package engine

type Cxx byte

const (
	CxxP Cxx = iota
	CxxN
	CxxPr
	CxxMc
	CxxVc
	CxxC19
	CxxC0
	CxxcC
	CxxCc
	CxxCC
)

// Cislovacky su P, N, Pr, Mc, Vc, C19, C0, cC, Cc, CC
type Cislovacka struct {
	c [10]byte
}

func (c *Cislovacka) Get(i Cxx) int {
	return int(c.c[i])
}

// NewCislovacky vytvori Cislovacky pre cislo n. Cislovacky maju zmysel pre n z intervalu <1, 99>
func NewCislovacka(k Kombinacia) *Cislovacka {
	var c Cislovacka
	for _, n := range k {
		if IsP(n) {
			c.c[CxxP]++
		} else {
			c.c[CxxN]++
		}
		if IsPr(n) {
			c.c[CxxPr]++
		}
		if IsMc(n) {
			c.c[CxxMc]++
		} else {
			c.c[CxxVc]++
		}
		if IsC19(n) {
			c.c[CxxC19]++
		}
		if IsC0(n) {
			c.c[CxxC0]++
		}
		if IscC(n) {
			c.c[CxxcC]++
		}
		if IsCc(n) {
			c.c[CxxCc]++
		}
		if IsCC(n) {
			c.c[CxxCC]++
		}
	}
	return &c
}

func NewCislovackaMax(n, m int) *Cislovacka {
	k := make(Kombinacia, m)
	for i := range k {
		k[i] = i + 1
	}
	c := NewCislovacka(k)
	nb := byte(n)
	if c.c[CxxP] > nb {
		c.c[CxxP] = nb
	}
	if c.c[CxxN] > nb {
		c.c[CxxN] = nb
	}
	if c.c[CxxPr] > nb {
		c.c[CxxPr] = nb
	}
	if c.c[CxxMc] > nb {
		c.c[CxxMc] = nb
	}
	if c.c[CxxVc] > nb {
		c.c[CxxVc] = nb
	}
	if c.c[CxxC19] > nb {
		c.c[CxxC19] = nb
	}
	if c.c[CxxC0] > nb {
		c.c[CxxC0] = nb
	}
	if c.c[CxxcC] > nb {
		c.c[CxxcC] = nb
	}
	if c.c[CxxCc] > nb {
		c.c[CxxCc] = nb
	}
	if c.c[CxxCC] > nb {
		c.c[CxxCC] = nb
	}
	return c
}

// IsP kontroluje ci je cislo parne
func IsP(n int) bool {
	return n%2 == 0
}

// IsN kontroluje ci je cislo neparne
func IsN(n int) bool {
	return n%2 == 1
}

// IsPr kontroluje ci je n prvocislo do 100.
func IsPr(n int) bool {
	switch n {
	case 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97:
		return true
	default:
		return false
	}
}

// IsMc kontroluje ci je cislo male cislo.
// Male cisla su 1..5, 11.15, ..., 91..95
func IsMc(n int) bool {
	n %= 10
	return n >= 1 && n <= 5
}

// IsVc kontroluje ci je cislo velke cislo.
// Velke cisla su 6..10, 16..20, ..., 86..90
func IsVc(n int) bool {
	n %= 10
	return n >= 6 || n == 0
}

// IsC19 kontroluje ci je cislo v intervale <1, 9>
func IsC19(n int) bool {
	return n >= 1 && n <= 9
}

// IsC0 kontroluje ci dvojciferne cislo konci na "0"
// C0 cisla su: 10, 20, ..., 90
func IsC0(n int) bool {
	return n%10 == 0 && n >= 10
}

// IscC kontroluje ci v dvojcifernom cisle je prva cifra mensia ako druha.
// cC cisla su: 12, 13, 14, 15, 23, 24 ...
func IscC(n int) bool {
	if n > 11 {
		return n/10 < n%10
	}
	return false
}

// IsCc kontroluje ci v dvojcifernom cisle je prva cifra vacsia ako druha.
// Cc cisla su: 21, 31, 32, 41, 42, 43 ...
func IsCc(n int) bool {
	if n > 20 && !IsC0(n) {
		return n/10 > n%10
	}
	return false
}

// IsCC kontroluje ci v dvojcifernom cisle je prva cifra rovna druhej.
// CC cisla su: 11, 22, 33, 44 ...
func IsCC(n int) bool {
	return n/10 == n%10
}
