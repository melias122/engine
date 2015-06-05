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

func (a *Archiv) add2Reverse(kombinacie [][]int) Uc {
	var (
		u int
		v = make(map[int]bool, a.m)
	)
	a.Hrx = hrx.New(a.m, func(n *num.N) int { return n.PocR2() })
	for _, N := range a.Cisla {
		N.Reset2()
	}
	var (
		is101 bool
		i     = len(kombinacie)
	)
	for !is101 {
		i--
		for y, x := range kombinacie[i] {
			N := a.Cisla[x-1]
			N.Inc2(y)
			a.Hrx.Add(N)
			v[x] = true
			if len(v) == a.m && !is101 {
				u = x
				is101 = true
			}
		}
	}
	return Uc{u, i}
}

func (a *Archiv) add2(kombinacia []int) bool {
	if a.Uc.Cislo == 0 {
		return true
	}
	for y, x := range kombinacia {
		if x == a.Uc.Cislo {
			return true
		}
		N := a.Cisla[x-1]
		N.Inc2(y)
		a.Hrx.Add(N)
	}
	return false
}

func (a *Archiv) add1(kombinacia []int) (*komb.K, bool) {
	var (
		K *komb.K = komb.New(a.n, a.m)
	)
	for y, x := range kombinacia {
		if a.Cisla[x-1] == nil {
			N := num.New(x, a.n, a.m)
			a.Cisla[x-1] = N

			a.nCisla++
			if a.nCisla == a.m {
				a.Is101 = true
			}

			K.Push(N)
			N.Inc1(y)
			a.HHrx.Add(N)
		} else {
			N := a.Cisla[x-1]
			N.Inc1(y)
			a.HHrx.Add(N)
			K.Push(N)
		}
	}
	return K, a.Is101
}

func (a *Archiv) write(chanErrKomb chan ErrKomb) error {

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
	var (
		kombinacie = make([][]int, 0, 64)
	)
	for errKomb := range chanErrKomb {

		if errKomb.Err != nil {
			return errKomb.Err
		}
		komb := errKomb.Komb
		kombinacie = append(kombinacie, komb)

		a.Pc++
		K, is101 := a.add1(komb)
		if is101 {
			if reverse := a.add2(komb); reverse {
				uc := a.add2Reverse(kombinacie)
				a.Uc = Uc{Cislo: uc.Cislo, Riadok: uc.Riadok + a.Uc.Riadok}
				kombinacie = kombinacie[uc.Riadok:]
			}
		}
		if a.Pc > 1 {
			r0 := a.Riadok
			a.Riadok.add(K, a.Hrx.Get(), a.HHrx.Get())
			a.Riadok.diff(r0)
		} else {
			a.Uc.Riadok++ // TODO: nespravne ukazovanie uc predtym nez nastane 101
			a.Riadok.add(K, 100, 100)
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

	// Vytvorenie suboru
	if err := os.Mkdir(dir, 0755); err != nil {
		log.Printf("%s\n", err)
	}

	archiv := New(dir, n, m)
	if err := archiv.write(Parse(path, n, m)); err != nil {
		return nil, err
	}
	return archiv, nil
}
