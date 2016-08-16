package engine

func Xtica(m int, k Kombinacia) Tica {
	xtica := make(Tica, (m+9)/10)
	for _, n := range k {
		xtica[(n-1)/10]++
	}
	return xtica
}
