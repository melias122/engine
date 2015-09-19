package num

import (
	"bytes"
	"log"
	"strconv"
)

type C [10]byte

func NewC(cislo int) C {
	if cislo <= 0 {
		log.Println("NewC cislo == 0")
		return C{}
	}
	var c C
	for i, f := range []func(int) bool{IsP, IsN, IsPr, IsMc, IsVc, IsC19, IsC0, IscC, IsCc, IsCC} {
		if f(cislo) {
			c[i]++
		}
	}
	return c
}

func (c1 *C) Plus(c2 C) {
	for i, c := range c2 {
		c1[i] += c
	}
}

func (c1 *C) Minus(c2 C) {
	for i, c := range c2 {
		c1[i] -= c
	}
}

func (c1 C) String() string {
	var buf bytes.Buffer
	for i, c := range c1 {
		if i > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(strconv.Itoa(int(c)))
	}
	return buf.String()
}

func (c C) Strings() []string {
	s := make([]string, len(c))
	for i, cislo := range c {
		s[i] = strconv.Itoa(int(cislo))
	}
	return s
}

func IsP(c int) bool {
	return c%2 == 0
}

func IsN(c int) bool {
	return c%2 == 1
}

func IsPr(c int) bool {
	if c != 2 && IsP(c) {
		return false
	}
	switch c {
	case 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97:
		return true
	default:
		return false
	}
}

func IsMc(c int) bool {
	c %= 10
	return c < 6 && c > 0
}

func IsVc(c int) bool {
	return !IsMc(c)
}

// 1..9
func IsC19(c int) bool {
	return c < 10
}

// 10, 20, 30, 40 ...
func IsC0(c int) bool {
	return c%10 == 0
}

// 12, 13, 14, 15, 23, 24 ...
func IscC(c int) bool {
	return c/10 < c%10 && c > 10
}

// 21, 31, 32, 41, 42, 43 ...
func IsCc(c int) bool {
	x := c % 10
	return c/10 > x && x != 0
}

// 11, 22, 33, 44 ...
func IsCC(c int) bool {
	return c/10 == c%10
}
