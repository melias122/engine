package engine

import (
	"strconv"
	"strings"
)

type Kombinacia []int

func (k Kombinacia) Len() int {
	return len(k)
}

func (k Kombinacia) Copy() Kombinacia {
	cp := make(Kombinacia, len(k))
	copy(cp, k)
	return cp
}

func (k Kombinacia) String() string {
	buf := make([]byte, 0, len(k)*3)
	space := ""
	for _, n := range k {
		buf = append(buf, space...)
		space = " "
		buf = strconv.AppendInt(buf, int64(n), 10)
	}
	return string(buf)
}

func (k Kombinacia) Contains(num int) bool {
	for i := range k {
		if k[i] == num {
			return true
		}
	}
	return false
}

func (k Kombinacia) Cislovacky() Cislovacky {
	var c Cislovacky
	for _, cislo := range k {
		c.Plus(NewCislovacky(int(cislo)))
	}
	return c
}

func (k Kombinacia) Sucet() int {
	var sucet int
	for _, cislo := range k {
		sucet += int(cislo)
	}
	return sucet
}

func (k Kombinacia) SucetRS(n Nums) (float64, float64) {
	var r, s float64
	for i, cislo := range k {
		if n[cislo-1] != nil {
			r += n[cislo-1].R()
			s += n[cislo-1].S(i + 1)
		}
	}
	return r, s
}

func (k Kombinacia) SucetRSNext(n Nums) (float64, float64) {
	var r, s float64
	for i, cislo := range k {
		if n[cislo-1] != nil {
			r += n[cislo-1].RNext()
			s += n[cislo-1].SNext(i + 1)
		}
	}
	return r, s
}

func (k Kombinacia) SledPN() string {
	s := make([]string, len(k))
	for i, j := range k {
		if IsP(int(j)) {
			s[i] = "P"
		} else {
			s[i] = "N"
		}
	}
	return strings.Join(s, " ")
}

func (k Kombinacia) SledPNPr() string {
	s := make([]string, len(k))
	for i, j := range k {
		j := int(j)
		if IsPr(j) {
			s[i] = "Pr"
		} else if IsN(j) {
			s[i] = "N"
		} else {
			s[i] = "P"
		}
	}
	return strings.Join(s, " ")
}

func (k Kombinacia) SledMcVc() string {
	s := make([]string, len(k))
	for i, j := range k {
		if IsMc(int(j)) {
			s[i] = "Mc"
		} else {
			s[i] = "Vc"
		}
	}
	return strings.Join(s, " ")
}

func (k Kombinacia) SledPrirodzene() string {
	s := make([]string, len(k))
	for i, j := range k {
		j := int(j)
		if IsC19(j) {
			s[i] = "C19"
		} else if IsC0(j) {
			s[i] = "C0"
		} else if IscC(j) {
			s[i] = "cC"
		} else if IsCc(j) {
			s[i] = "Cc"
		} else { // CC
			s[i] = "CC"
		}
	}
	return strings.Join(s, " ")
}
