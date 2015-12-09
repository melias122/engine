package psl

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Archiv
type Archiv struct {
	Riadok

	Hrx  *H
	HHrx *H

	CsvPath    string
	WorkingDir string
	Suffix     string

	origHeader []string
	riadky     []Riadok

	Skupiny Skupiny

	Predict1DO  Prediction
	PredictODDO Prediction
}

// NewArchiv funkcia vytvori archiv aj z vystupmi. V pripade ze currentWorkingDir
// je "-" archiv nevytvori vystupy a len sa interne nacita.
func NewArchiv(csvPath, currentWorkingDir string, n, m int) (*Archiv, error) {
	// Cesta k suboru musi byt zadana
	if csvPath == "" {
		log.Println("Archiv: csvPath: ", csvPath)
		return nil, fmt.Errorf("Archiv: Nebola zadana cesta k súboru!")
	}

	// Aktualny pracovny priecinok musi byt zadany
	if currentWorkingDir == "" {
		log.Println("Archiv: cwd: ", currentWorkingDir)
		return nil, fmt.Errorf("Archiv: Nebol zadany pracovny priecinok")
	}

	// Skontrolovanie minimalneho a maximalneho rozmeru databazy
	if n < 2 || n >= m || m > 99 {
		log.Println("Archiv: Zadany rozmer: ", n, "/", m)
		return nil, fmt.Errorf("Archiv: Nesprávny rozmer databázy: %d/%d", n, m)
	}

	var filename string
	if s := strings.Split(filepath.Base(csvPath), ".csv"); len(s) != 2 {
		log.Printf("Archiv: %s not a .csv", s)
		return nil, fmt.Errorf("Archiv: Subor %s musi byt typu csv", csvPath)
	} else {
		filename = s[0]
	}

	var (
		WorkingDir string
		CsvPath    string
		Suffix     string
	)
	if currentWorkingDir == "-" {
		WorkingDir = currentWorkingDir
		CsvPath = csvPath
	} else {
		WorkingDir = filepath.Join(currentWorkingDir, filename)
		CsvPath = filepath.Join(WorkingDir, filename+".csv")
		Suffix = filepath.Base(WorkingDir)
	}

	archiv := &Archiv{
		Riadok: Riadok{
			n: n,
			m: m,
		},
		Hrx:  NewHrx(n, m),
		HHrx: NewHHrx(n, m),

		CsvPath:    CsvPath,
		WorkingDir: WorkingDir,
		Suffix:     Suffix,
	}

	// Vytvorenie suboru
	if archiv.WorkingDir != "-" {
		if err := os.Mkdir(archiv.WorkingDir, 0755); err != nil {
			log.Printf("Archiv: %s\n", err)
		}
		// Skopirovanie originalnej databazy (.csv)
		if err := CopyFile(archiv.CsvPath, csvPath); err != nil {
			log.Printf("Archiv: %s\n", err)
			return nil, fmt.Errorf("Archiv: Nepodarilo sa skopirovat subor %s", csvPath)
		}
	}
	// vytvorenie archivu
	if err := archiv.loadCsv(Parse(archiv.CsvPath, n, m)); err != nil {
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
	if err := savePredictions(archiv.WorkingDir, archiv.Predict1DO, archiv.PredictODDO); err != nil {
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

	writter := NewCsvMaxWriter("Archiv", a.WorkingDir, setHeader(archivRiadokHeader))
	defer func() {
		e = writter.Close()
	}()

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
	writter := NewCsvMaxWriter("ArchivR+1", a.WorkingDir, setHeader(archivRiadokHeader))
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
	if a.WorkingDir == "-" {
		return
	}
	for _, f := range []func() error{
		a.PocetnostR,
		a.PocetnostS,
		a.mapaXtice,
		a.mapaZhoda,
		a.statistikaZhoda,
		a.mapaNtice,
		a.statistikaNtice,
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

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(dst, src string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	// Dont want to create hard link..
	// if err = os.Link(src, dst); err == nil {
	// 	return
	// }

	// instead we copy file
	err = copyFileContents(src, dst)
	return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
