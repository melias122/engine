package komb

import (
	"bytes"
	"strconv"
)

type Kombinacia []byte

func (k Kombinacia) String() string {
	var buf bytes.Buffer
	for i, n := range k {
		if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.Itoa(int(n)))
	}
	return buf.String()
}

func (k Kombinacia) Contains(cislo byte) bool {
	switch len(k) {
	case 0:
		return false
	case 1:
		return k[0] == cislo
	default:
		return bytes.Contains(k, []byte{cislo})
	}

}
