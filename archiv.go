package engine

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gitlab.com/melias122/engine/csv"
)

type Dimension struct {
	N, M int
}

// Archiv
type Archiv struct {
	Riadok
	Dimension

	Hrx  *H
	HHrx *H

	Csvpath string
	Workdir string
	Suffix  string

	origHeader []string
	riadky     []Riadok

	Skupiny Skupiny

	Predict1DO  Prediction
	PredictODDO Prediction
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
		Riadok: Riadok{
			n: n,
			m: m,
		},
		Dimension: Dimension{N: n, M: m},
		Hrx:       NewHrx(n, m),
		HHrx:      NewHHrx(n, m),

		Csvpath: filepath.Join(workdir, basename),
		Workdir: workdir,
		Suffix:  suffix,
	}

	// Vytvorenie suboru
	if err := os.MkdirAll(archiv.Workdir, 0755); err != nil {
		log.Printf("Archiv: %s\n", err)
	}

	// Skopirovanie originalnej databazy (.csv)
	if err := CopyFile(archiv.Csvpath, csvpath); err != nil {
		log.Printf("Archiv: %s\n", err)
		return nil, fmt.Errorf("Archiv: Nepodarilo sa skopirovat subor %s", csvpath)
	}

	// vytvorenie archivu
	if err := archiv.loadCsv(Parse(archiv.Csvpath, n, m)); err != nil {
		return nil, err
	}

	// vystupy
	if err := archiv.makeFiles(); err != nil {
		return nil, err
	}

	if err := archiv.skupiny(); err != nil {
		return nil, err
	}

	normalizePrediction(&archiv.Predict1DO, archiv.Skupiny)
	normalizePrediction(&archiv.PredictODDO, archiv.Skupiny)
	if err := savePredictions(archiv.Workdir, archiv.Predict1DO, archiv.PredictODDO); err != nil {
		return nil, err
	}

	return archiv, nil
}

func (a *Archiv) skupiny() (err error) {
	// vytvorenie HrxHHrx vystupu a skupiny na filtrovanie/generovanie
	a.Skupiny, err = makeSkupiny(a)
	return
}

func (a *Archiv) loadCsv(chanErrKomb chan ErrKomb) (e error) {

	writter := csv.NewCsvMaxWriter("Archiv", a.Workdir, csv.SetHeader(archivRiadokHeader))
	defer func() {
		e = writter.Close()
	}()

	var (
		kombinacie = make([][]int, 0, 64)
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
		if a.HHrx.Cisla.Is101() {
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
				for !a.Hrx.Cisla.Is101() {
					uc.Riadok--
					for y, x := range kombinacie[uc.Riadok] {
						a.Hrx.Add(int(x), y)
						if a.Hrx.Cisla.Is101() && uc.Cislo == 0 {
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
		a.riadky = append(a.riadky, a.Riadok)

		if err := writter.Write(a.Riadok.record()); err != nil {
			log.Println(err)
			return err
		}
	}
	return
}

func (a *Archiv) rPlus1() error {
	writter := csv.NewCsvMaxWriter("ArchivR+1", a.Workdir, csv.SetHeader(archivRiadokHeader))
	defer func() {
		writter.Close()
	}()
	var (
		hrxNums  = a.Hrx.Cisla.rplus1()
		hhrxNums = a.HHrx.Cisla.rplus1()
		r0       = Riadok{n: a.n, m: a.m}
	)
	for _, r := range a.riadky {
		r0.Add(r.K,
			hhrxNums,
			hrxNums,
			a.Hrx.ValueKombinacia(r.K),
			a.HHrx.ValueKombinacia(r.K),
		)
		if err := writter.Write(r0.record()); err != nil {
			return err
		}
	}
	return nil
}

func (a *Archiv) makeFiles() (err error) {
	if a.Workdir == "-" {
		return
	}
	for _, f := range []func() error{
		a.PocetnostR,
		a.PocetnostS,
		a.mapaXtice,
		a.mapaXtice2,
		a.mapaZhoda2,
		a.statistikaZhoda,
		a.mapaNtice,
		a.statistikaNtice2,
		a.statistikaCifrovacky,
		a.statistikaCislovacky,
		a.predikcia,
		a.rPlus1,
	} {
		e := f()
		if e != nil {
			return e
		}
	}
	return
}

type ErrKomb struct {
	Header []string
	Komb   Kombinacia
	Orig   []string
	Err    error
}

func parse(record []string, n int) (Kombinacia, error) {
	var (
		komb  = make([]int, n)
		err   error
		cislo int
	)
	for i, field := range record[3 : n+3] {
		if strings.ContainsAny(field, ".,") {
			field = strings.Replace(field, ",", ".", -1)
			f, err := strconv.ParseFloat(field, 64)
			if err != nil {
				return nil, err
			}
			cislo = int(f)
		} else {
			cislo, err = strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
		}
		komb[i] = cislo
	}
	return komb, nil
}

func Parse(path string, n, m int) chan ErrKomb {

	ch := make(chan ErrKomb, 1)
	go func() {
		defer close(ch)

		file, err := os.Open(path)
		if err != nil {
			ch <- ErrKomb{Err: err}
			return
		}
		defer file.Close()

		r := csv.NewReader(file)

		// Skip Header
		header, _ := r.Read()

		var nline int
		for {
			nline++
			record, err := r.Read()
			if err == io.EOF {
				return
			}
			if err != nil {
				ch <- ErrKomb{Err: err}
				return
			}
			if len(record) < n+3 {
				ch <- ErrKomb{Err: fmt.Errorf("%s: riadku %d musi byt dlhsi ako %d", file.Name(), nline, n+3)}
				return
			}
			komb, err := parse(record, n)
			if err != nil {
				ch <- ErrKomb{Err: err}
			} else {
				recordCopy := make([]string, len(record))
				copy(recordCopy, record)
				ch <- ErrKomb{Komb: komb, Orig: recordCopy, Header: header}
			}
		}
	}()
	return ch
}
