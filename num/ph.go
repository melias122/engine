package num

type IncFn func(int) float64

type ph struct {
	p  int
	h  float64
	fn IncFn
}

func newph(x, y, n, m int) *ph {
	return &ph{
		fn: func(p int) float64 { return Value(p, x, y, n, m) },
	}
}

func (p *ph) inc() {
	p.p++
	p.h = p.fn(p.p)
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
