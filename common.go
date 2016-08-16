package engine

import (
	"strconv"
	"unicode/utf8"
)

func bytesToString(b []byte) string {
	var (
		space []byte
		buf   = make([]byte, 0, 128)
	)
	for _, u := range b {
		buf = append(buf, space...)
		buf = strconv.AppendUint(buf, uint64(u), 10)
		space = []byte(" ")
	}
	return string(buf[0:len(buf)])
}

func itoa(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func Ftoa(f float64) string {
	return ftoa(f)
}

func ftoa(f float64) string {
	buf := make([]byte, 0, 64)
	buf = strconv.AppendFloat(buf, f, 'g', -1, 64)
	for i, w := 0, 0; i < len(buf); i += w {
		runeValue, width := utf8.DecodeRune(buf[i:])
		if runeValue == '.' {
			buf[i] = ','
			break
		}
		w = width
	}
	return string(buf[:len(buf)])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
