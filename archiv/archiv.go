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
	Cisla map[int]*num.N
	Hrx   *hrx.H
	HHrx  *hrx.H

	Dir  string
	Path string
}

func New(dir string, n, m int) *Archiv {
	return &Archiv{
		n:     n,
		m:     m,
		Cisla: make(map[int]*num.N, m),
		Hrx:   hrx.NewHrx(m),
		HHrx:  hrx.NewHHrx(m),

		Dir:  dir,
		Path: filepath.Join(dir, "Archiv_"+dir+".csv"),
	}
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
		K          *komb.K
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

		// Vytvorenie kombinacie pre aktualny Riadok
		K = komb.New(a.n, a.m)

		// Prechadzame kombinaciu na riadku
		// a inkrementujeme pocetnost cisla R1-Do
		for y, x := range errKomb.Komb {
			N, ok := a.Cisla[x]
			// Ak cislo neexistuje, treba ho vytvorit
			// Nastava pri kazdom cisle prave 1
			if !ok {
				N = num.New(x, a.n, a.m)
				a.Cisla[x] = N
			}
			N.Inc1(y)
			a.HHrx.Add(N)
		}
		// Ak mame plnu databazu(nastala udalost 101)
		if len(a.Cisla) == a.m {
			// Ked narazime na Uc cislo na riadku Roddo
			// potrebuje spatne dohladat Uc cislo
			var reverse bool

			// Uc cislo je 0 len raz po udalosti 101
			if a.Uc.Cislo == 0 {
				reverse = true
			} else {
				// Incrementovanie cisla Roddo
				for y, x := range errKomb.Komb {
					// Ak na riadku narazime na Uc Cislo
					// porebujeme ho spatne dohladat
					if x == a.Uc.Cislo {
						reverse = true
					}
					N := a.Cisla[x]
					N.Inc2(y)
					a.Hrx.Add(N)
				}
			}
			// Spatne dohladanie Uc cisla a riadku a inrementovanie cisiel Roddo
			if reverse {
				// Nova hrx zostava a resetovanie cisiel Roddo
				a.Hrx = hrx.NewHrx(a.m)
				for _, N := range a.Cisla {
					N.Reset2()
				}
				var (
					is101  = false
					vyskyt = make(map[int]bool, a.m)
					uc     = Uc{Riadok: len(kombinacie)}
				)
				// Spatne nacitava kombinacie a incremtuje Roddo
				// a Hrx az pokial nenastane udalost 101
				// udalost 101 nastava ked sa kazde cislo vyskytne aspon 1
				for !is101 {
					uc.Riadok--
					for y, x := range kombinacie[uc.Riadok] {
						N := a.Cisla[x]
						N.Inc2(y)
						a.Hrx.Add(N)
						vyskyt[x] = true
						if len(vyskyt) == a.m && !is101 {
							uc.Cislo = x
							is101 = true
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
			// Vytvorenie kombinacie
			for _, c := range errKomb.Komb {
				K.Push(a.Cisla[c].Copy(a.n, a.m))
			}
		} else {
			// Ak databaza este nie je naplnena
			// cisla s pocetnostou 1 sa tvaria ako 0
			// kvoli tomu aby hodnoty R1-DO v archive zacinali od 0,0
			for _, c := range errKomb.Komb {
				N := a.Cisla[c]
				if N.PocR1() > 1 {
					K.Push(N.Copy(a.n, a.m))
				} else {
					K.Push(num.New(N.Cislo(), a.n, a.m))
				}
			}
		}

		// Zostavenie Riadkov pre Archiv
		// pre 2 a viac riadkov robime rozdiel(diff) vybranych hodnot
		if a.Pc > 1 {
			r0 := a.Riadok
			a.Riadok.add(K, a.Hrx.Get(), a.HHrx.Get())
			a.Riadok.diff(r0)
		} else {
			a.Uc.Riadok++             // TODO: nespravne ukazovanie uc predtym nez nastane 101
			a.Riadok.add(K, 100, 100) // V prvom riadku hodnoty hrx a hhrx su nastavene natvrdo 100
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
