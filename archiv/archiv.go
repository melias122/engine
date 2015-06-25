package archiv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/melias122/psl/hrx"
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
	Hrx  *hrx.H
	HHrx *hrx.H

	Dir  string
	Path string

	riadky []Riadok
}

func New(dir string, n, m int) *Archiv {
	archiv := &Archiv{
		n: n,
		m: m,
		Riadok: Riadok{
			n: n,
			m: m,
		},
		Hrx:  hrx.New(n, m),
		HHrx: hrx.New(n, m),

		Dir:  dir,
		Path: filepath.Join(dir, "Archiv_"+dir+".csv"),
	}
	return archiv
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
		kombinacie = make([][]byte, 0, 64)
	)
	for errKomb := range chanErrKomb {
		a.Pc++

		// Skontrolujem ci nenastala chyba pri parsovani
		if errKomb.Err != nil {
			return errKomb.Err
		}
		// Treba odkladat kombinacie, kvoli spatnemu
		// dohladaniu Uc
		kombinacie = append(kombinacie, errKomb.Komb)

		// Prechadzame kombinaciu na riadku
		// a inkrementujeme pocetnost cisla R1-Do
		for y, x := range errKomb.Komb {
			a.HHrx.Add(int(x), y)
		}
		// Ak sa v hhrx vyskytli vsetky cisla 1..m
		// nastala udalost 101
		if a.HHrx.Is101() {
			// Ked narazime na Uc cislo na riadku Roddo
			// potrebuje spatne dohladat Uc cislo
			var reverse bool

			// Uc cislo je 0 len raz po udalosti 101
			if a.Uc.Cislo == 0 {
				reverse = true
			} else {
				// Incrementovanie cisla Roddo, resp. Hrx
				for y, x := range errKomb.Komb {
					// Ak na riadku narazime na Uc Cislo
					// porebujeme ho spatne dohladat
					if x == a.Uc.Cislo {
						reverse = true
					}
					a.Hrx.Add(int(x), y)
				}
			}
			// Spatne dohladanie Uc cisla a riadku a inrementovanie cisiel Roddo
			if reverse {
				// Nova hrx zostava a resetovanie cisiel Roddo
				a.Hrx = hrx.New(a.n, a.m)
				uc := Uc{Riadok: len(kombinacie)}
				// Spatne nacitava kombinacie a incremtuje Roddo
				// a Hrx az pokial nenastane udalost 101
				// udalost 101 nastava ked sa kazde cislo vyskytne aspon 1
				for !a.Hrx.Is101() {
					uc.Riadok--
					for y, x := range kombinacie[uc.Riadok] {
						a.Hrx.Add(int(x), y)
						if a.Hrx.Is101() && uc.Cislo == 0 {
							uc.Cislo = x
						}
					}
				}
				// Nastavenie noveho Uc cisla a riadku pre archiv
				a.Uc.Cislo = uc.Cislo
				a.Uc.Riadok += uc.Riadok

				// Mozeme predchadzajuce kombinacie zahodit
				// a zaujimaju nas uz len od posledneho Uc riadku
				kombinacie = kombinacie[uc.Riadok:]
			}
		}

		// Vytvorenie kombinacie
		var (
			n1 = num.Zero(a.n, a.m)
			n2 = num.Zero(a.n, a.m)
		)
		if a.HHrx.Is101() {
			for _, x := range errKomb.Komb {
				n1.Plus(a.HHrx.GetN(int(x)))
				n2.Plus(a.Hrx.GetN(int(x)))
			}
		} else {
			// Ak databaza este nie je naplnena
			// cisla s pocetnostou 1 sa tvaria ako 0
			// kvoli tomu aby hodnoty R1-DO v archive zacinali od 0,0
			for _, x := range errKomb.Komb {
				N := a.HHrx.GetN(int(x))
				if N.PocetR() > 1 {
					n1.Plus(a.HHrx.GetN(int(x)))
				} else {
					n1.Plus(num.New(int(x), a.n, a.m))
				}
			}
		}

		// Zostavenie Riadkov pre Archiv
		// pre 2 a viac riadkov robime rozdiel(diff) vybranych hodnot
		if a.Pc > 1 {
			a.Add(errKomb.Komb, n1, n2, a.Hrx.Value(), a.HHrx.Value())
		} else {
			a.Uc.Riadok++                         // TODO: nespravne ukazovanie uc predtym nez nastane 101
			a.Add(errKomb.Komb, n1, n2, 100, 100) // V prvom riadku hodnoty hrx a hhrx su nastavene natvrdo 100
		}
		a.riadky = append(a.riadky, a.Riadok)

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
	chanErrKomb := Parse(path, n, m)

	// Archiv.csv
	if err := archiv.write(chanErrKomb); err != nil {
		return nil, err
	}

	// PocetnostR.csv
	// if err := archiv.PocetnostR(); err != nil {
	// return nil, err
	// }

	// PocetnostS.csv
	// if err := archiv.PocetnostS(); err != nil {
	// 	return nil, err
	// }

	// Mapa Ntice
	// if err := archiv.MapaNtice(); err != nil {
	// 	return nil, err
	// }

	return archiv, nil
}
