package filter

import (
	"bytes"

	"github.com/melias122/psl/komb"
)

type zakazane struct {
	cisla []byte
}

func NewZakazane(cisla []byte) Filter {
	return zakazane{
		cisla: cisla,
	}
}

func (z zakazane) Check(k komb.Kombinacia) bool {
	if len(k) > 0 && len(z.cisla) > 0 {
		cislo := k[len(k)-1]
		if bytes.Contains(z.cisla, []byte{cislo}) {
			return false
		} else {
			return true
		}
	}
	return true
}

type zakazaneStl struct {
	zakazane []Filter
}

func NewZakazaneStl(cisla [][]byte) Filter {
	zakazane := make([]Filter, len(cisla))
	for i, z := range cisla {
		zakazane[i] = NewZakazane(z)
	}
	return zakazaneStl{
		zakazane: zakazane,
	}
}

func (z zakazaneStl) Check(k komb.Kombinacia) bool {
	if len(k) > 0 && len(k) <= len(z.zakazane) {
		return z.zakazane[len(k)-1].Check(k)
	}
	return true
}
