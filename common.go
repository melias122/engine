package psl

import (
	"math"
	"strconv"
	"unicode/utf8"
)

func bytesToString(b []byte) string {
	buf := make([]byte, 0, 128)
	space := []byte(" ")
	for i, u := range b {
		if i > 0 {
			buf = append(buf, space...)
		}
		buf = strconv.AppendUint(buf, uint64(u), 10)
	}
	return string(buf[:len(buf)])
}

func itoa(i int) string {
	return strconv.FormatInt(int64(i), 10)
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

func lessFloat64(a, b float64) bool {
	return math.Float64bits(a) < math.Float64bits(b)
}

func greaterFloat64(a, b float64) bool {
	return math.Float64bits(a) > math.Float64bits(b)
}

func outOfRangeFloats64(x1, x2, y1, y2 float64) bool {
	if greaterFloat64(x1, y2) || lessFloat64(x2, y1) {
		return true
	}
	return false
}
