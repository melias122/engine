package filter

import "github.com/melias122/psl/komb"

type povinne struct {
	n     int
	cisla []byte
}

func NewPovinne(n int, cisla []byte) Filter {
	return povinne{
		n:     n,
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
