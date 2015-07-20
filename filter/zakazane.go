package filter

import "github.com/melias122/psl/komb"

type zakazane struct {
	cisla []bool
}

func NewZakazane(cisla []bool) Filter {
	return zakazane{
		cisla: cisla,
	}
}

func (z zakazane) Check(k komb.Kombinacia) bool {
	cislo := k[len(k)-1] - 1
	if z.cisla[cislo] {
		return false
	} else {
		return true
	}
}

type zakazaneStl struct {
	cisla [][]bool
}

func NewZakazaneStl(cisla [][]bool) Filter {
	return zakazaneStl{
		cisla: cisla,
	}
}

func (z zakazaneStl) Check(k komb.Kombinacia) bool {
	stlpec := len(k) - 1
	cislo := k[stlpec] - 1
	if z.cisla[stlpec][cislo] == true {
		return false
	}
	return true
}
