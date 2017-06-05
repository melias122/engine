package engine

import (
	"strconv"
)

func (k Kombinacia) String() string {
	buf := make([]byte, 0, len(k)*3)
	space := ""
	for _, n := range k {
		buf = append(buf, space...)
		space = " "
		buf = strconv.AppendInt(buf, int64(n), 10)
	}
	return string(buf)
}

func (k Kombinacia) Sucet() int {
	var sucet int
	for _, cislo := range k {
		sucet += int(cislo)
	}
	return sucet
}
