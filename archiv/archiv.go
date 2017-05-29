package archiv

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/melias122/engine/csv"
	"github.com/melias122/engine/engine"
	"github.com/melias122/engine/hrx"
	"github.com/pkg/errors"
)

// Archiv
type Archiv struct {
	engine.Riadok

	Hrx  *hrx.H
	HHrx *hrx.H

	Csvpath string
	Workdir string
	Suffix  string

	riadky []engine.Riadok
}

// NewArchiv funkcia vytvori archiv aj z vystupmi. V pripade ze outdir
// je "-" archiv nevytvori vystupy a len sa interne nacita.
func NewArchiv(csvpath, outdir string, n, m int) (*Archiv, error) {

	if n < 2 || n >= m {
		return nil, fmt.Errorf("Archiv: Nesprávny rozmer databázy: %d/%d", n, m)
	}

	if _, err := os.Stat(csvpath); os.IsNotExist(err) {
		return nil, err
	}

	if _, err := os.Stat(outdir); os.IsNotExist(err) {
		return nil, err
	}

	basename := filepath.Base(csvpath)
	suffix := strings.TrimSuffix(basename, filepath.Ext(basename))
	workdir := filepath.Join(outdir, suffix)

	archiv := &Archiv{
		Riadok: engine.Riadok{
		//n: n,
		//m: m,
		},
		Hrx:  hrx.NewHrx(n, m),
		HHrx: hrx.NewHHrx(n, m),

		Csvpath: filepath.Join(workdir, basename),
		Workdir: workdir,
		Suffix:  suffix,
	}

	// Vytvorenie suboru
	if err := os.MkdirAll(archiv.Workdir, 0755); err != nil {
		log.Printf("Archiv: %s\n", err)
	}

	// Skopirovanie originalnej databazy (.csv)
	if err := copyFile(archiv.Csvpath, csvpath); err != nil {
		log.Printf("Archiv: %s\n", err)
		return nil, fmt.Errorf("Archiv: Nepodarilo sa skopirovat subor %s", csvpath)
	}

	// vytvorenie archivu
	file, err := os.Open(archiv.Csvpath)
	if err != nil {
		return nil, errors.Wrap(err, "could not open file")
	}
	defer file.Close()

	parser := csv.NewParser(file, n, m)
	kombinacie, err := parser.Parse()
	if err != nil {
		return nil, errors.Wrap(err, "could not parse")
	}

	if err := archiv.create([][]int(kombinacie), n, m); err != nil {
		return nil, err
	}

	return archiv, nil
}

func (a *Archiv) create(kombinacie [][]int, n, m int) (e error) {

	writter := csv.NewCsvMaxWriter("Archiv", a.Workdir, csv.SetHeader(engine.ArchivRiadokHeader))

	defer func() {
		e = writter.Close()
	}()

	for r, k := range kombinacie {

		// Prechadzame kombinaciu na riadku
		// a inkrementujeme pocetnost cisla R1-Do
		for y, x := range k {
			a.HHrx.Add(int(x), y)
		}

		// Ak sa v hhrx vyskytli vsetky cisla 1..m
		// nastala udalost 101
		if a.HHrx.Cisla.Is101() {
			// Ked narazime na Uc cislo na riadku Roddo
			// potrebuje spatne dohladat Uc cislo
			var reverse bool

			// Uc cislo je 0 len raz po udalosti 101
			if a.Uc.Cislo == 0 {
				reverse = true
			} else {
				// Incrementovanie cisla Roddo, resp. Hrx
				for y, x := range k {
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
				a.Hrx = hrx.NewHrx(n, m)
				uc := engine.Uc{Riadok: r}
				// Spatne nacitava kombinacie a incremtuje Roddo
				// a Hrx az pokial nenastane udalost 101
				// udalost 101 nastava ked sa kazde cislo vyskytne aspon 1
				for !a.Hrx.Cisla.Is101() {
					for y, x := range kombinacie[uc.Riadok] {
						a.Hrx.Add(int(x), y)
						if a.Hrx.Cisla.Is101() && uc.Cislo == 0 {
							uc.Cislo = x
						}
					}
					uc.Riadok--
				}
				// Nastavenie noveho Uc cisla a riadku pre archiv
				a.Uc = uc
			}
		}

		// Zostavenie Riadkov pre Archiv
		// pre 2 a viac riadkov robime rozdiel(diff) vybranych hodnot
		if a.Pc > 1 {
			a.Add(k, a.HHrx.Cisla, a.Hrx.Cisla, a.Hrx.Value(nil), a.HHrx.Value(nil))
		} else {
			a.Uc.Riadok++                                 // TODO: nespravne ukazovanie uc predtym nez nastane 101
			a.Add(k, a.HHrx.Cisla, a.Hrx.Cisla, 100, 100) // V prvom riadku hodnoty hrx a hhrx su nastavene natvrdo 100
		}
		a.riadky = append(a.riadky, a.Riadok)

		if err := writter.Write(a.Riadok.Record()); err != nil {
			log.Println(err)
			return err
		}
	}
	return
}
