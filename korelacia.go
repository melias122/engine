package engine

import "math"

func Korelacia(k0, k1 Kombinacia, n, m int) float64 {
	if len(k0) != n || len(k1) != n {
		return 0.0
	}

	var kk float64
	for i := 0; i < n; i++ {
		x := (float64(k1[i]) - float64(k0[i])) / float64(m)
		x *= x //^2
		x *= x //^4
		kk += x / float64(n)
	}
	kk = float64(1) - math.Sqrt(kk)
	kk *= kk //^2
	kk *= kk //^4
	kk *= kk //^8
	return kk
}
