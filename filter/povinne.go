package filter

import "github.com/melias122/psl/komb"

type povinne struct {
	cisla []bool
}

func NewPovinne(cisla []bool) Filter {
	return povinne{
		cisla: cisla,
	}
}

func (p povinne) Check(k komb.Kombinacia) bool {
	return true
}

type povinneStl struct {
	cisla [][]bool
}

func NewPovinneStl(cisla [][]bool) Filter {
	return povinneStl{
		cisla: cisla,
	}
}

func (p povinneStl) Check(k komb.Kombinacia) bool {
	return true
}
