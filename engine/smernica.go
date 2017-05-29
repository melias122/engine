package engine

func Smernica(k Kombinacia, n, m int) (sm float64) {
	if len(k) < 2 {
		return .0
	}
	nSm := 0
	n--
	m--
	for i, n0 := range k[:len(k)-1] {
		for j, n1 := range k[i+1:] {
			sm += (float64(n1-n0) / float64(m)) / (float64(j+1) / float64(n))
			nSm++
		}
	}
	return sm / float64(nSm)
}
