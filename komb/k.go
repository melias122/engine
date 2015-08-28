package komb

import (
	"bytes"
	"strconv"

	"github.com/melias122/psl/num"
)

type Kombinacia []byte

func (k Kombinacia) String() string {
	var buf bytes.Buffer
	for i, n := range k {
		if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.Itoa(int(n)))
	}
	return buf.String()
}

func (k Kombinacia) Contains(cislo byte) bool {
	switch len(k) {
	case 0:
		return false
	case 1:
		return k[0] == cislo
	default:
		return bytes.Contains(k, []byte{cislo})
	}
}

func (k Kombinacia) Cislovacky() num.C {
	var c num.C
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
