package psl

import "bytes"

type Tica []byte

func (t Tica) String() string {
	return bytesToString(t)
}

func (t0 Tica) Equal(t Tica) bool {
	return bytes.Equal(t0, t)
}
