package psl

type Nums []*Num

func (n Nums) Is101() bool {
	for _, N := range n {
		if N == nil {
			return false
		}
	}
	return true
}

func (c Nums) Len() int           { return len(c) }
func (c Nums) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Nums) Less(i, j int) bool { return c[i].Cislo() < c[j].Cislo() }

type ByPocetR struct {
	Nums
}

func (by ByPocetR) Less(i, j int) bool {
	return by.Nums[i].PocetR() < by.Nums[j].PocetR()
}
