package archiv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/melias122/psl/hrx"
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
	Is101  bool
	nCisla int
	Cisla  []*num.N
	Hrx    *hrx.H
	HHrx   *hrx.H

	Dir  string
	Path string
}

func New(dir string, n, m int) *Archiv {
	return &Archiv{
		n:     n,
		m:     m,
		Cisla: make([]*num.N, m),
		Hrx:   hrx.New(m, func(n *num.N) int { return n.PocR2() }),
		HHrx:  hrx.New(m, func(n *num.N) int { return n.PocR1() }),

		Dir:  dir,
		Path: filepath.Join(dir, "Archiv_"+dir+".csv"),
	}
}

func (a *Archiv) add2Reverse(k [][]byte) uc {
	var (
		u int
		v = make(map[int]bool, a.m)
	)
	a.Hrx = hrx.New(a.m, func(n *num.N) int { return n.PocR2() })
	for _, c := range a.Cisla {
		c.Reset2()
	}
	i := len(k)
	full := false
	for !full {
		i--
		for y := range k[i] {
			x := int(k[i][y])
			c := a.Cisla[x-1]
			c.Inc2(y)
			a.Hrx.Add(c)
			v[x] = true
			if len(v) == a.m && !full {
				u = x
				full = true
			}
		}
	}
	return uc{u, i + 1}
}

func (a *Archiv) add2(k []byte) bool {
	if a.uc.Cislo == 0 {
		return true
	}
	for y := range k {
		x := int(k[y])
		if x == a.uc.Cislo {
			return true
		}
		c := a.Cisla[x-1]
		c.Inc2(y)
		a.Hrx.Add(c)
	}
	return false
}

func (a *Archiv) add1(k []byte) (*komb.K, bool) {
	var (
		// c  *num.N
		ko *komb.K = komb.New(a.n, a.m)
	)
	for y := range k {
		x := int(k[y])
		if a.Cisla[x-1] == nil {
			c := num.New(x, a.n, a.m)
			a.Cisla[x-1] = c

			a.nCisla++
			if a.nCisla == a.m {
				a.Is101 = true
			}

			ko.Push(c)
			c.Inc1(y)
			a.HHrx.Add(c)
		} else {
			c := a.Cisla[x-1]
			c.Inc1(y)
			a.HHrx.Add(c)
			ko.Push(c)
		}
	}
	return ko, a.Is101
}

func (a *Archiv) write(Kombinacie [][]byte) error {

	file, err := os.Create(a.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.Comma = ';'
	if err := w.Write(header); err != nil {
		return err
	}

	for _, c := range Kombinacie {
		a.Pc++
		k, is101 := a.add1(c)
		if is101 {
			if reverse := a.add2(c); reverse {
				a.uc = a.add2Reverse(Kombinacie[:a.Pc])
			}
		}
		if a.Pc > 1 {
			r0 := a.Riadok
			a.Riadok.add(k, a.Hrx.Get(), a.HHrx.Get())
			a.Riadok.diff(r0)
		} else {
			a.Riadok.add(k, 100, 100)
		}
		if err := w.Write(a.Riadok.record()); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func Make(path string, n, m int) (*Archiv, error) {

	if path == "" {
		return nil, fmt.Errorf("Nebola zadana cesta k súboru!")
	}
	if n < 2 || n >= m || m > 99 {
		return nil, fmt.Errorf("Nesprávny rozmer databázy: %d/%d", n, m)
	}
	base := filepath.Base(path)
	dir := strings.Split(base, ".")[0]

	// parse csv
	Kombinacie, err := Parse(path, n, m)
	if err != nil {
		return nil, err
	}

	// Vytvorenie suboru
	err = os.Mkdir(dir, 0755)
	if err != nil {
		log.Printf("%s\n", err)
	}

	a := New(dir, n, m)
	err = a.write(Kombinacie)
	if err != nil {
		return nil, err
	}
	return a, nil
}
