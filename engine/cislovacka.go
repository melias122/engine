package engine

// Cislovacky su P, N, Pr, Mc, Vc, C19, C0, cC, Cc, CC
type Cislovacka struct {
	P   byte
	N   byte
	Pr  byte
	Mc  byte
	Vc  byte
	C19 byte
	C0  byte
	McC byte `csv:"cC"`
	VCc byte `csv:"Cc"`
	CC  byte `csv:"CC"`
}

// NewCislovacky vytvori Cislovacky pre cislo n. Cislovacky maju zmysel pre n z intervalu <1, 99>
func NewCislovacka(k Kombinacia) Cislovacka {
	var c Cislovacka
	for _, n := range k {
		if IsP(n) {
			c.P++
		} else {
			c.N++
		}
		if IsPr(n) {
			c.Pr++
		}
		if IsMc(n) {
			c.Mc++
		} else {
			c.Vc++
		}
		if IsC19(n) {
			c.C19++
		}
		if IsC0(n) {
			c.C0++
		}
		if IscC(n) {
			c.McC++
		}
		if IsCc(n) {
			c.VCc++
		}
		if IsCC(n) {
			c.CC++
		}
	}
	return c
}

func NewCislovackaMax(n, m int) Cislovacka {
	k := make(Kombinacia, m)
	for i := range k {
		k[i] = i + 1
	}
	c := NewCislovacka(k)
	nb := byte(n)
	if c.P > nb {
		c.P = nb
	}
	if c.N > nb {
		c.N = nb
	}
	if c.Pr > nb {
		c.Pr = nb
	}
	if c.Mc > nb {
		c.Mc = nb
	}
	if c.Vc > nb {
		c.Vc = nb
	}
	if c.C19 > nb {
		c.C19 = nb
	}
	if c.C0 > nb {
		c.C0 = nb
	}
	if c.McC > nb {
		c.McC = nb
	}
	if c.VCc > nb {
		c.VCc = nb
	}
	if c.CC > nb {
		c.CC = nb
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
// Velke cisla su 6..0, 16..20, ..., 86..90
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
