package archiv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/melias122/psl/komb"
	"github.com/melias122/psl/num"
)

var header = []string{
	"Poradové číslo", "Kombinacie", "P", "N", "PR", "Mc", "Vc",
	"c1-c9", "C0", "cC", "Cc", "CC", "ZH",
	"Sm", "ΔSm", "Kk", "ΔKk", "N-tice", "X-tice",
	"ƩR1-DO", "ΔƩR1-DO", "HHRX", "ΔHHRX",
	"ƩR OD-DO", "ΔƩR OD-DO", "HRX", "ΔHRX",
	"ƩKombinacie", "ΔƩKombinacie", "UC číslo", "UC riadok",
}

type Archiv struct {
	n, m int
	Riadok
	Cisla map[int]*num.N
	hrx   *hrx
	hhrx  *hrx
}

func New(n, m int) *Archiv {
	return &Archiv{
		n:     n,
		m:     m,
		Cisla: make(map[int]*num.N, m),
		hrx:   newhrx(m, func(n *num.N) int { return n.PocR2() }),
		hhrx:  newhrx(m, func(n *num.N) int { return n.PocR1() }),
	}
}

func (a *Archiv) add2Reverse(k [][]int) uc {
	var (
		n, u int
		v    = make(map[int]bool, a.m)
	)
	a.hrx = newhrx(a.m, func(n *num.N) int { return n.PocR2() })
	for _, c := range a.Cisla {
		c.Reset2()
	}
	n = len(k)
	full := false
	for !full {
		for y, x := range k[n-1] {
			c := a.Cisla[x]
			c.Inc2(y)
			a.hrx.add(c)
			v[x] = true
			if len(v) == a.m && !full {
				u = x
				full = true
			}
		}
		n--
	}
	return uc{u, n + 1}
}

func (a *Archiv) add2(k []int) bool {
	if a.uc.Cislo == 0 {
		return true
	}
	for y, x := range k {
		if x == a.uc.Cislo {
			return true
		}
		c := a.Cisla[x]
		c.Inc2(y)
		a.hrx.add(c)
	}
	return false
}

func (a *Archiv) add1(k []int) (*komb.K, bool) {
	var (
		ok bool
		c  *num.N
		ko *komb.K = komb.New(a.n, a.m)
	)
	for y, x := range k {
		if c, ok = a.Cisla[x]; !ok {
			c = num.New(x, a.n, a.m)
			a.Cisla[x] = c

			ko.Push(c)
			c.Inc1(y)
			a.hhrx.add(c)
		} else {
			c.Inc1(y)
			a.hhrx.add(c)
			ko.Push(c)
		}
	}
	return ko, len(a.Cisla) == a.m
}

func (a *Archiv) write(C [][]int) error {
	os.Mkdir(fmt.Sprintf("%d%d", a.n, a.m), 0755)
	o, err := os.Create(fmt.Sprintf("%d%d/Archiv_%d%d.csv", a.n, a.m, a.n, a.m))
	if err != nil {
		return err
	}
	defer o.Close()

	w := csv.NewWriter(o)
	w.Comma = ';'
	if err := w.Write(header); err != nil {
		return err
	}
	for _, c := range C {
		a.Pc++
		k, is101 := a.add1(c)
		if is101 {
			if reverse := a.add2(c); reverse {
				a.uc = a.add2Reverse(C[:a.Pc])
			}
		}
		if a.Pc > 1 {
			r0 := a.Riadok
			a.Riadok.add(k, a.hrx.hrx(), a.hhrx.hrx())
			a.Riadok.diff(r0)
		} else {
			a.Riadok.add(k, a.hrx.hrx(), a.hhrx.hrx())
		}
		if err := w.Write(a.Riadok.record()); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func (a *Archiv) Parse(f *os.File) error {
	k, err := a.parse(f)
	if err != nil {
		return err
	}
	return a.write(k)
}

func (a *Archiv) parse(f *os.File) ([][]int, error) {
	r := csv.NewReader(f)
	r.Comma = rune(';')

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var k [][]int
	for i, r := range records[1:] {
		if len(r) < a.n+3 {
			return nil,
				fmt.Errorf("%s: na riadku %d sa nepodarilo nacitat kombinaciu", f.Name(), i+1)
		}
		c := make([]int, a.n)
		for i, x := range r[3 : a.n+3] {
			cislo, err := strconv.ParseUint(x, 10, 0)
			if err != nil {
				return nil, err
			}
			c[i] = int(cislo)
		}

		k = append(k, c)
	}
	return k, nil
}
