package komb

import "math"

func Kk(n, m int, k0, k1 []int) float64 {
	var kk float64
	for i := 0; i < n; i++ {
		// kk += math.Pow((a-p)/float64(m), 4) / float64(n)
		x := float64(k1[i]) - float64(k0[i])
		x *= x //^2
		x *= x //^4
		kk += x / float64(n)
	}
	kk = float64(1) - math.Sqrt(kk)
	kk *= kk //^2
	kk *= kk //^4
	kk *= kk //^8
	return kk
	// return math.Pow(float64(1)-math.Sqrt(kk), 8)
}
