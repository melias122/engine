package num

import "psl2/math"

type ph struct {
	p int
	h float64
	f func()
}

func newph(x, y, n, m int) *ph {
	var ph ph
	ph.f = func() {
		ph.p++
		ph.h = math.Value(ph.p, x, y, n, m)
	}
	return &ph
}

func (p *ph) inc() {
	p.f()
}

func (p *ph) reset() {
	p.p = 0
	p.h = 0
}

func (p *ph) plus(r *ph) *ph {
	p.h += r.h
	return p
}

func (p *ph) minus(r *ph) *ph {
	p.h -= r.h
	return p
}
