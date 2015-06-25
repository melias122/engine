package num

type ph struct {
	key
	hodnota float64
}

func newph(x, y, n, m int) ph {
	// Kvoli symetrickosti binomickych cisiel
	// maju cisla na poziciach rovnaku pocetnost
	// Priklad db n=5, m=35..
	// cislo 1 == 35, 2 == 34 ...
	// stlpec 1 == 5, 2 == 4 ...
	if x > (m/2)+m%2 {
		x = (x - (m + 1)) * (-1)
		y = n - y + 1
	}
	return ph{
		key: key{x: byte(x), y: byte(y), n: byte(n), m: byte(m)},
	}
}

func (p *ph) inc() {
	p.pocet++
	p.hodnota = value(p.key)
}

func (p *ph) plus(r ph) {
	p.hodnota += r.hodnota
}

func (p *ph) minus(r ph) {
	p.hodnota -= r.hodnota
}

func (p ph) Pocet() int {
	return int(p.pocet)
}

func (p ph) Hodnota() float64 {
	return p.hodnota
}
