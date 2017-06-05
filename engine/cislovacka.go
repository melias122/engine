package engine

import (
	"fmt"
	"strconv"
)

const (
	P Cislovacka = iota
	N
	Pr
	Mc
	Vc
	C19
	C0
	XcC
	Cc
	CC
)

const _Cislovacka_name = "PNPrMcVcC19C0cCCcCC"

var _Cislovacka_index = [...]uint8{0, 1, 2, 4, 6, 8, 11, 13, 15, 17, 19}

func (i Cislovacka) String() string {
	if i >= Cislovacka(len(_Cislovacka_index)-1) {
		return fmt.Sprintf("Cislovacka(%d)", i)
	}
	return _Cislovacka_name[_Cislovacka_index[i]:_Cislovacka_index[i+1]]
}

var cislovackyFuncs = [...]cislovackaFunc{IsP, IsN, IsPr, IsMc, IsVc, IsC19, IsC0, IscC, IsCc, IsCC}

func (i Cislovacka) Func() cislovackaFunc {
	if i > Cislovacka(len(cislovackyFuncs)-1) {
		return nil
	}
	return cislovackyFuncs[i]
}

// Cislovacky su P, N, Pr, Mc, Vc, C19, C0, cC, Cc, CC
type Cislovacky [10]byte

// FunCislovacky su funkcie, ktore vyhodnocuju ci je cislo danou cislovackou
type cislovackaFunc func(int) bool

func CislovackyMax(n, m int) Cislovacky {
	var c Cislovacky
	for i := 1; i <= m; i++ {
		c2 := NewCislovacky(i)
		c.Plus(c2)
	}
	for i := range c {
		if c[i] > byte(n) {
			c[i] = byte(n)
		}
	}
	return c
}

// NewCislovacky vytvori Cislovacky pre cislo n. Cislovacky maju zmysel pre n z intervalu <1, 99>
func NewCislovacky(n int) Cislovacky {
	if n < 1 || n > 99 {
		panic("could not create Cislovacky")
	}
	var c Cislovacky
	for i, f := range cislovackyFuncs {
		if f(n) {
			c[i]++
		}
	}
	return c
}

func NewKCislovacky(k Kombinacia) Cislovacky {
	var c Cislovacky
	for _, cislo := range k {
		c.Plus(NewCislovacky(int(cislo)))
	}
	return c
}

// Plus pricita k c c2. Teda c = c + c2.
// W: Moze prist k preteceniu!
func (c *Cislovacky) Plus(c2 Cislovacky) {
	for i, j := range c2 {
		c[i] += j
	}
}

// Minus odcita z c c2. Teda c = c - c2.
// W: Moze prist k preteceniu!
func (c *Cislovacky) Minus(c2 Cislovacky) {
	for i, j := range c2 {
		c[i] -= j
	}
}

// String implementuje interface Stringer
func (c *Cislovacky) String() string {
	return bytesToString(c[0:len(c)])
}

// Strings je pomocna funkcia pre vypis jednotlivych cislovacie
func (c *Cislovacky) Strings() []string {
	s := make([]string, len(c))
	for i, c := range c {
		s[i] = strconv.Itoa(int(c))
	}
	return s
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
