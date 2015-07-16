package filter

import "github.com/melias122/psl/num"

func NewP(n, min, max int) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: num.IsP,
	}
}

func NewN(n, min, max int) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: num.IsN,
	}
}

func NewPr(n, min, max int) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: num.IsPr,
	}
}

func NewMc(n, min, max int) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: num.IsMc,
	}
}

func NewVc(n, min, max int) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: num.IsVc,
	}
}

func NewC19(n, min, max int) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: num.IsC19,
	}
}

func NewC0(n, min, max int) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: num.IsC0,
	}
}

func NewcC(n, min, max int) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: num.IscC,
	}
}
func NewCc(n, min, max int) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: num.IsCc,
	}
}

func NewCC(n, min, max int) Filter {
	return cislovacky{
		n:   n,
		min: min,
		max: max,
		fun: num.IsCC,
	}
}

type cislovacky struct {
	n        int
	min, max int
	fun      func(int) bool
}

func (c cislovacky) Check(combination []int) bool {
	var count int
	for _, number := range combination {
		if c.fun(number) {
			count++
		}
	}
	if len(combination) == c.n {
		if !(count >= c.min && count <= c.max) {
			return false
		}
	} else {
		if count > c.max {
			return false
		}
	}
	return true
}
