package komb

import (
	"bytes"
	"strconv"
)

type Tica []byte

func (t Tica) String() string {
	var buf bytes.Buffer
	for i, n := range t {
		if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.Itoa(int(n)))
	}
	return buf.String()
}

func Xtica(m int, k Kombinacia) Tica {
	return xtica(m, k)
}

func xtica(m int, k Kombinacia) Tica {
	xtica := make(Tica, (m+9)/10)
	for _, number := range k {
		xtica[(number-1)/10]++
	}
	return xtica
}

func Ntica(kombinacia Kombinacia) Tica {
	tica, _ := ntica(kombinacia)
	return tica
}

func NticaPozicie(kombinacia Kombinacia) []byte {
	_, pozicie := ntica(kombinacia)
	return pozicie
}

func ntica(k Kombinacia) (Tica, []byte) {
	if len(k) == 0 {
		return Tica{}, []byte{}
	}
	if len(k) == 1 {
		return Tica{1}, []byte{0}
	}
	var (
		n       int
		tica    = make(Tica, len(k))
		pozicie = make([]byte, len(k))
	)
	for i := range k[:len(k)-1] {
		if k[i]+1 == k[i+1] {
			n++
			pozicie[i] = 1
			pozicie[i+1] = 1
		} else {
			tica[n]++
			n = 0
		}
	}
	tica[n]++
	return tica, pozicie
}
