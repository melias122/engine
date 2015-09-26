package komb

import (
	"strconv"
	"strings"
)

type Tica []byte

func (t Tica) String() string {
	s := make([]string, len(t))
	for i, n := range t {
		s[i] = strconv.Itoa(int(n))
	}
	return strings.Join(s, " ")
}

func Xtica(m int, k Kombinacia) Tica {
	return xtica(m, k)
}

func xtica(m int, k Kombinacia) Tica {
	xtica := make(Tica, (m+9)/10)
	for _, n := range k {
		xtica[(n-1)/10]++
	}
	return xtica
}

func Ntica(k Kombinacia) Tica {
	tica, _ := ntica(k)
	return tica
}

func NticaPozicie(k Kombinacia) []byte {
	_, pozicie := ntica(k)
	return pozicie
}

func ntica(k Kombinacia) (Tica, []byte) {
	if len(k) == 0 {
		return Tica{}, []byte{}
	}
	if len(k) == 1 {
		return Tica{1}, []byte{0}
	}
	var (
		n       int
		tica    = make(Tica, len(k))
		pozicie = make([]byte, len(k))
	)
	for i := range k[:len(k)-1] {
		if k[i]+1 == k[i+1] {
			n++
			pozicie[i] = 1
			pozicie[i+1] = 1
		} else {
			tica[n]++
			n = 0
		}
	}
	tica[n]++
	return tica, pozicie
}

// nticaSS sluzi na prevod suctu alebo
// sucinu na retazec
// Pozn.: neskor mozno do filtra
type nticaSS []int

func (n nticaSS) String() string {
	s := make([]string, len(n))
	for i, n := range n {
		s[i] = strconv.Itoa(n)
	}
	return strings.Join(s, ", ")
}

// NticaSucet urcuje sucet cisiel Kombinacie
// na pozicii ntice
// Kombinacia{1,2,3,5,6}
// NticaSucet{1+2+3, 5+6} == {6, 11}
func NticaSucet(k Kombinacia) nticaSS {
	return ss(k, sucet)
}

// NticaSucet urcuje sucin pozicie a stlpca
// Kombinacie na pozicii ntice
// Kombinacia{10,11,12,15,16}
// NticaSucet{1*2*3, 4*5} == {6, 20}
func NticaSucin(k Kombinacia) nticaSS {
	return ss(k, sucin)
}

type operacia int

const (
	sucet operacia = iota
	sucin
)

func ss(k Kombinacia, o operacia) nticaSS {
	if len(k) < 2 {
		return nticaSS{}
	}
	var (
		n     nticaSS
		spolu int
	)
	for i := range k[:len(k)-1] {
		if k[i]+1 == k[i+1] {
			switch o {
			case sucet:
				if spolu == 0 {
					spolu = int(k[i])
				}
				spolu += int(k[i+1])
			case sucin:
				if spolu == 0 {
					spolu = int(i + 1)
				}
				spolu *= int(i + 2)
			}
		} else if spolu != 0 {
			n = append(n, spolu)
			spolu = 0
		}
	}
	if spolu != 0 {
		n = append(n, spolu)
	}
	return n
}
