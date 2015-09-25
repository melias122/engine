package komb

import (
	"strconv"
	"strings"
)

func Zhoda(k0, k1 Kombinacia) int {
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

// presun urcuje poziciu presunu cisla
// z Kombinacie k0 do Kombinacie k1
type presun [][2]byte

func ZhodaPresun(k0, k1 Kombinacia) presun {
	var p presun
	for i, j := 0, 0; i < len(k0) && j < len(k1); {
		if k0[i] == k1[j] {
			p = append(p, [2]byte{k0[i], k1[j]})
			i++
			j++
		} else if k0[i] < k1[j] {
			i++
		} else {
			j++
		}
	}
	return p
}

func (p presun) String() string {
	s := make([]string, len(p))
	for i, p := range p {
		s[i] = strconv.Itoa(int(p[0])) + "/" + strconv.Itoa(int(p[1]))
	}
	return strings.Join(s, ", ")
}
