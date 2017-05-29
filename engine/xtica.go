package engine

func NewXtica(k Kombinacia, m int) Tica {
	xtica := make(Tica, (m+9)/10)
	for _, n := range k {
		xtica[(n-1)/10]++
	}
	return xtica
}
