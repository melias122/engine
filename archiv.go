package psl

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var archivHeader = []string{
	"Poradové číslo", "Kombinacie", "P", "N", "PR", "Mc", "Vc", "c1-c9", "C0", "cC", "Cc",
	"CC", "ZH", "Sm", "\u0394Sm", "Kk", "\u0394Kk", "N-tice", "X-tice", "ƩR1-DO", "ΔƩR1-DO",
	"ƩR1-DO \"r+1\"", "ƩSTL1-DO", "ΔƩSTL1-DO", "ƩSTL1-DO \"r+1\"", "Δ(ƩR1-DO-ƩSTL1-DO)",
	"HHRX", "ΔHHRX", "ƩR OD-DO", "ΔƩR OD-DO", "ƩSTL OD-DO", "ΔƩSTL OD-DO",
	"Δ(ƩROD-DO-ƩSTLOD-DO)", "HRX", "ΔHRX", "ƩKombinacie", "ΔƩKombinacie",
	"UC číslo", "UC riadok",
	"Cifra 1", "Cifra 2", "Cifra 3", "Cifra 4", "Cifra 5", "Cifra 6", "Cifra 7", "Cifra 8", "Cifra 9", "Cifra 0",
}

type Archiv struct {
	n, m int
	Riadok
	Hrx  *H
	HHrx *H

	WorkingDir string
	Suffix     string

	origHeader []string
	riadky     []Riadok

	Skupiny Skupiny
}

func NewArchiv(workingDir string, n, m int) *Archiv {
	archiv := &Archiv{
		n: n,
		m: m,
		Riadok: Riadok{
			n: n,
			m: m,
		},
		Hrx:  NewHrx(n, m),
		HHrx: NewHHrx(n, m),

		WorkingDir: workingDir,
		Suffix:     filepath.Base(workingDir),
	}
	return archiv
}

func (a *Archiv) write(chanErrKomb chan ErrKomb) error {

	w := NewCsvMaxWriter(a.WorkingDir, "Archiv", [][]string{archivHeader})
	defer w.Close()

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
				a.Hrx = NewHrx(a.n, a.m)
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

		// Zostavenie Riadkov pre Archiv
		// pre 2 a viac riadkov robime rozdiel(diff) vybranych hodnot
		if a.Pc > 1 {
			a.Add(errKomb.Komb, a.HHrx.Cisla, a.Hrx.Cisla, a.Hrx.Value(), a.HHrx.Value())
		} else {
			a.Uc.Riadok++                                            // TODO: nespravne ukazovanie uc predtym nez nastane 101
			a.Add(errKomb.Komb, a.HHrx.Cisla, a.Hrx.Cisla, 100, 100) // V prvom riadku hodnoty hrx a hhrx su nastavene natvrdo 100
		}
		a.origHeader = errKomb.Header
		a.Riadok.origStrings = errKomb.Orig
		// fmt.Println(a.Riadok.origStrings)
		a.riadky = append(a.riadky, a.Riadok)
		// fmt.Println(a.Riadok.origStrings)

		if err := w.Write(a.Riadok.record()); err != nil {
			return err
		}
	}
	return nil
}

func Make(path, workingDir string, n, m int) (*Archiv, error) {

	if path == "" {
		return nil, fmt.Errorf("Nebola zadana cesta k súboru!")
	}
	if n < 2 || n >= m || m > 99 {
		return nil, fmt.Errorf("Nesprávny rozmer databázy: %d/%d", n, m)
	}
	dir := strings.Split(filepath.Base(path), ".")[0]
	dir = filepath.Join(workingDir, dir)

	// Vytvorenie suboru
	if err := os.Mkdir(dir, 0755); err != nil {
		log.Printf("Archiv.Make(): %s\n", err)
	}

	archiv := NewArchiv(dir, n, m)
	chanErrKomb := Parse(path, n, m)

	if err := archiv.write(chanErrKomb); err != nil {
		return nil, err
	}

	if err := archiv.PocetnostR(); err != nil {
		return nil, err
	}

	if err := archiv.PocetnostS(); err != nil {
		return nil, err
	}

	if err := archiv.mapaXtice(); err != nil {
		return nil, err
	}

	if err := archiv.mapaZhoda(); err != nil {
		return nil, err
	}

	if err := archiv.statistikaZhoda(); err != nil {
		return nil, err
	}

	if err := archiv.mapaNtice(); err != nil {
		return nil, err
	}
	if err := archiv.statistikaNtice(); err != nil {
		return nil, err
	}

	if err := archiv.statistikaCifrovacky(); err != nil {
		return nil, err
	}

	if err := archiv.statistikaCislovacky(); err != nil {
		return nil, err
	}

	hrxtab := NewHrxTab(archiv.Hrx, archiv.HHrx, n, m)
	hrxSkupiny, err := hrxtab.Make(archiv.WorkingDir)
	if err != nil {
		return nil, err
	}
	archiv.Skupiny = hrxSkupiny

	return archiv, nil
}
