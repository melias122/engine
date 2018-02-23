package engine

func Zhoda(k0, k1 Kombinacia) (zh int) {
	if k0 == nil || k1 == nil {
		return
	}
	for i, j := 0, 0; i < len(k0) && j < len(k1); {
		if k0[i] == k1[j] {
			zh++
			i++
			j++
		} else if k0[i] < k1[j] {
			i++
		} else {
			j++
		}
	}
	return
}
