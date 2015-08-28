package filter

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/melias122/psl/hrx"
	"github.com/melias122/psl/komb"
)

type Filter interface {
	fmt.Stringer
	Check(komb.Kombinacia) bool
	CheckSkupina(hrx.Skupina) bool
}

type Filters []Filter

func (f Filters) Check(k komb.Kombinacia) bool {
	for _, filter := range f {
		if !filter.Check(k) {
			return false
		}
	}
	return true
}

func (f Filters) CheckSkupina(skupina hrx.Skupina) bool {
	for _, filter := range f {
		if !filter.CheckSkupina(skupina) {
			return false
		}
	}
	return true
}

func (f Filters) String() string {
	var s string
	for _, filter := range f {
		s += filter.String() + "\n"
	}
	return s
}

func ParseFloat(s string) (f float64, e error) {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, ",", ".", 1)
	f, e = strconv.ParseFloat(s, 64)
	return
}

func ParseBytes(s string) (c []byte, e error) {
	for _, str := range strings.Split(s, ",") {
		var i int
		i, e = strconv.Atoi(str)
		if e != nil {
			return
		}
		c = append(c, byte(i))
	}
	return
}

func ParseNBytes(n int, s string) ([][]byte, error) {
	var cisla [][]byte = make([][]byte, n)
	for _, strs := range strings.Split(s, ";") {
		stlCisla := strings.Split(strs, ":")
		if len(stlCisla) != 2 {
			return nil, errors.New("Parse error")
		} else {
			stl, err := strconv.Atoi(stlCisla[0])
			if err != nil {
				return nil, err
			}
			if stl < 1 || stl > n {
				return nil, errors.New("Nesprávne zadaný stlpec")
			}
			cisla[stl-1], err = ParseBytes(stlCisla[1])
			if err != nil {
				return nil, err
			}
		}
	}
	return cisla, nil
}

func parseTica(s string) (t komb.Tica, e error) {
	s = strings.TrimSpace(s)
	for _, str := range strings.Split(s, " ") {
		var i int
		i, e = strconv.Atoi(str)
		if e != nil {
			return
		}
		t = append(t, byte(i))
	}
	return
}

func ParseNtica(n int, s string) (komb.Tica, error) {
	ntica, err := parseTica(s)
	if err != nil {
		return nil, err
	}
	if len(ntica) > n {
		return nil, fmt.Errorf("Dĺžka ntice musí byť <= %d", n)
	}
	var sum int
	for i, n := range ntica {
		sum += int(n) * (i + 1)
	}
	if sum != n {
		return nil, fmt.Errorf("Súčet ntice != %d", n)
	}
	for len(ntica) < n {
		ntica = append(ntica, 0)
	}
	return ntica, nil
}

func ParseXtica(n, m int, s string) (komb.Tica, error) {
	xtica, err := parseTica(s)
	if err != nil {
		return nil, err
	}
	lenXtica := (m + 9) / 10
	if len(xtica) > lenXtica {
		return nil, fmt.Errorf("Dĺžka xtice musí byť <= %d", lenXtica)
	}
	var sum int
	for _, n := range xtica {
		sum += int(n)
	}
	if sum != n {
		return nil, fmt.Errorf("Súčet xtice != %d", n)
	}
	for len(xtica) < lenXtica {
		xtica = append(xtica, 0)
	}
	return xtica, nil
}

func dt(f float64) float64 {
	var (
		s  = strconv.FormatFloat(f, 'f', -1, 64)
		dt = 1.0
	)
	if strings.Contains(s, ".") {
		s = strings.Split(s, ".")[1]
		for i := 0; i < len(s); i++ {
			dt /= 10
		}
	}
	return dt
}

func nextGRT(f float64) float64 {
	return f + dt(f)
}

func nextLSS(f float64) float64 {
	return f - dt(f)
}
