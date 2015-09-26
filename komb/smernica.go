package komb

func Smernica(n, m int, k Kombinacia) float64 {
	if len(k) < 2 {
		return .0
	}
	var (
		sm  float64
		nSm float64
		M   = float64(m - 1)
		N   = float64(n - 1)
	)
	for i, n0 := range k[:len(k)-1] {
		for j, n1 := range k[i+1:] {
			sm += (float64(n1-n0) / M) / (float64(j+1) / N)
			nSm++
		}
	}
	return sm / nSm
}
