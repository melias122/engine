package psl

type cislo struct {
	cislo byte
	pocet *int
}

func newCislo(c int, pocet *int) cislo {
	return cislo{byte(c), pocet}
}

type cisla []cislo

func newCisla(nums Nums, xcisla Xcisla) cisla {
	var cisla cisla
	skupinaPocet := make(map[int]*int)
	for _, tab := range xcisla {
		pocet := tab.Max
		skupinaPocet[tab.Sk] = &pocet
	}
	for _, N := range nums {
		if pocet, ok := skupinaPocet[N.PocetR()]; ok {
			cisla = append(cisla, newCislo(N.Cislo(), pocet))
		}
	}
	return cisla
}
