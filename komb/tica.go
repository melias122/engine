package komb

import (
	"bytes"
	"strconv"
)

type Tica []byte

func (t Tica) String() string {
	var buf bytes.Buffer
	for i, el := range t {
		if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.Itoa(int(el)))
	}
	return buf.String()
}
