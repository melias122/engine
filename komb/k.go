package komb

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/melias122/psl/num"
)

type Kombinacia []byte

func (k Kombinacia) String() string {
	s := make([]string, len(k))
	for i, n := range k {
		s[i] = strconv.Itoa(int(n))
	}
	return strings.Join(s, " ")
}

func (k Kombinacia) Contains(cislo byte) bool {
	return bytes.IndexByte(k, cislo) > -1
}

func (k Kombinacia) Cislovacky() num.Cislovacky {
	var c num.Cislovacky
	for _, cislo := range k {
		c.Plus(num.NewC(int(cislo)))
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

func (k Kombinacia) SucetRS(n num.Nums) (float64, float64) {
	var r, s float64
	for i, cislo := range k {
		if n[cislo-1] != nil {
			r += n[cislo-1].R()
			s += n[cislo-1].S(i + 1)
		}
	}
	return r, s
}

func (k Kombinacia) SucetRSNext(n num.Nums) (float64, float64) {
	var r, s float64
	for i, cislo := range k {
		if n[cislo-1] != nil {
			r += n[cislo-1].RNext()
			s += n[cislo-1].SNext(i + 1)
		}
	}
	return r, s
}
