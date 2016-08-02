package engine

import (
	"bytes"
	"strconv"
	"strings"
)

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

type filterNtica struct {
	n     int
	ntica Tica
	filterPriority
}

func NewFilterNtica(n int, tica Tica) Filter {
	return filterNtica{
		n:              n,
		ntica:          tica,
		filterPriority: P2,
	}
}

func (n filterNtica) Check(k Kombinacia) bool {
	nticaK := Ntica(k)
	if len(k) == n.n {
		return bytes.Equal(nticaK, n.ntica)
	}
	return true
}

func (n filterNtica) CheckSkupina(skupina Skupina) bool {
	return true
}

func (n filterNtica) String() string {
	return "Ntica: " + n.ntica.String()
}

type filterSTLNtica struct {
	n       int
	pozicie []byte
	ntica   Filter
	filterPriority
}

func NewFilterSTLNtica(n int, tica Tica, pozicie []byte) Filter {
	return filterSTLNtica{
		n:              n,
		ntica:          NewFilterNtica(n, tica),
		pozicie:        pozicie,
		filterPriority: P2,
	}
}

func (s filterSTLNtica) Check(k Kombinacia) bool {
	if !s.ntica.Check(k) {
		return false
	}
	if s.n == len(k) {
		return bytes.Equal(NticaPozicie(k), s.pozicie)
	}
	return true
}

func (s filterSTLNtica) CheckSkupina(h Skupina) bool {
	return true
}

func (s filterSTLNtica) String() string {
	var pozicie []string
	for i, p := range s.pozicie {
		if p == 1 {
			pozicie = append(pozicie, strconv.Itoa(i+1))
		}
	}
	return s.ntica.String() + "\n" + "STL Ntica: " + strings.Join(pozicie, ", ")
}
