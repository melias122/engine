package komb

func Smernica(n, m int, k []int) float64 {
	var (
		sm  float64
		nSm int
	)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			p1 := (float64(k[j]) - float64(k[i])) / float64(m-1)
			p2 := (float64(j) - float64(i)) / float64(n-1)
			p1 /= p2
			sm += p1
			nSm++
		}
	}
	if nSm > 0 {
		sm /= float64(nSm)
	}
	return sm
}
