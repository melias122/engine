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
