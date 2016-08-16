package engine

import "strings"

func Zhoda(k0, k1 Kombinacia) int {
	if k0 == nil || k1 == nil {
		return 0
	}
	var zhoda int
	for i, j := 0, 0; i < len(k0) && j < len(k1); {
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

// ZhodaPresun urcuje poziciu presunu cisla
// z Kombinacie k0 do Kombinacie k1
type ZhodaPresun struct {
	p [][2]int
}

func NewZhodaPresun(k0, k1 Kombinacia) ZhodaPresun {
	if k0 == nil || k1 == nil {
		return ZhodaPresun{}
	}
	var zp ZhodaPresun
	for i, j := 0, 0; i < len(k0) && j < len(k1); {
		if k0[i] == k1[j] {
			zp.p = append(zp.p, [2]int{i + 1, j + 1})
			i++
			j++
		} else if k0[i] < k1[j] {
			i++
		} else {
			j++
		}
	}
	return zp
}

func (zp ZhodaPresun) String() string {
	s := make([]string, len(zp.p))
	for i, p := range zp.p {
		s[i] = itoa(int(p[0])) + "|" + itoa(int(p[1]))
	}
	return strings.Join(s, ", ")
}
