package psl

import (
	"bytes"
	"strconv"
	"strings"
)

type Tica []byte

func (t Tica) String() string {
	s := make([]string, len(t))
	for i, n := range t {
		s[i] = strconv.Itoa(int(n))
	}
	return strings.Join(s, " ")
}

func (t0 Tica) Equal(t Tica) bool {
	return bytes.Equal(t0, t)
}
