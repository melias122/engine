package komb

func Zhoda(k0, k1 Kombinacia) int {
	return zhoda(k0, k1)
}

func zhoda(k0, k1 Kombinacia) int {
	var zhoda, i, j int
	for i < len(k0) && j < len(k1) {
		if k0[i] == k1[j] {
			zhoda++
			i++
			j++
		} else if k0[i] < k1[j] {
			i++
		} else {
			j++
		}
	}
	return zhoda
}
