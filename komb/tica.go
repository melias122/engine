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
	return ntica(kombinacia)
}

func ntica(kombinacia Kombinacia) Tica {
	if len(kombinacia) == 0 {
		return Tica{}
	}
	var (
		tice   = make(Tica, len(kombinacia))
		result = kombinacia
		tica   int
	)
	result = append(result, 0)
	for i := 0; i < len(result); i++ {
		if i == len(result)-2 {
			if tica > 0 {
				tice[tica]++
			} else {
				tice[0]++
			}
			break
		}
		c := int(int(result[i]) - int(result[i+1]))
		if c == int(result[i]) {
			if tica > 0 {
				tice[tica]++
			}
			break
		}
		if c < -1 {
			if tica == 0 {
				tice[0]++
			} else {
				tice[tica]++
				tica = 0
			}
		} else {
			tica++
		}
	}
	return tice
}
