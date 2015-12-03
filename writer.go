package psl

import (
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
)

const (
	// Standardne pouzivame ako oddelovat ";"
	Comma = ';'

	// Maximum Excel rows..
	MaxWrite = 500000

	// Ked nechceme zapisat nic
	DiscardCSV = "-"
)

type SuffixFunc func() string

func IntSuffix() SuffixFunc {
	var i int
	return func() string {
		i++
		return "_" + itoa(i)
	}
}

func DefaultSuffix(workingDir string) SuffixFunc {
	var (
		i         int
		suffix    = "_" + filepath.Base(workingDir)
		intSuffix = IntSuffix()
	)
	return func() string {
		if i == 0 {
			i++
			return suffix
		}
		return suffix + intSuffix()
	}
}

func EmptySuffix() SuffixFunc {
	return func() string {
		return ""
	}
}

type CsvMaxWriter struct {

	// aktualne zapisanych riadkov do suboru
	rowsWritten int

	// pocet riadkov zapisanych celkom
	totalRowsWritten int

	// aktualny pracovny priecinok
	workingDir string

	// podpriecikon v pracovnom priecinku
	subdir string

	// nazov suboru
	fileName string

	// suffix suboru
	Suffix SuffixFunc

	// hlavicka csv suboru
	header [][]string

	// aktualny subor kam sa zapisuje
	file *os.File

	// csv wrapper
	writer *csv.Writer

	// prvy zapis
	initialized bool
}

func setSuffixFunc(sf SuffixFunc) func(w *CsvMaxWriter) {
	return func(w *CsvMaxWriter) {
		w.Suffix = sf
	}
}

func setSubdir(subdir string) func(w *CsvMaxWriter) {
	return func(w *CsvMaxWriter) {
		w.subdir = subdir
	}
}

func setHeader(header []string) func(w *CsvMaxWriter) {
	return func(w *CsvMaxWriter) {
		w.header = [][]string{header}
	}
}

func setHeaders(headers [][]string) func(w *CsvMaxWriter) {
	return func(w *CsvMaxWriter) {
		w.header = headers
	}
}

func NewCsvMaxWriter(fileName, workingDir string, options ...func(*CsvMaxWriter)) *CsvMaxWriter {
	w := &CsvMaxWriter{
		workingDir: workingDir,
		fileName:   fileName,
		Suffix:     DefaultSuffix(workingDir),
	}
	for _, option := range options {
		option(w)
	}
	if workingDir != DiscardCSV {
		os.Mkdir(filepath.Join(w.workingDir, w.subdir), 0755)
	}
	return w
}

// TotalRowsWriten vrati pocet zapisanych riadkov
func (w *CsvMaxWriter) TotalRowsWriten() int {
	return w.totalRowsWritten
}

// Write zapise retazce record do suboru
func (w *CsvMaxWriter) Write(record []string) error {
	if !w.initialized {
		if err := w.add(); err != nil {
			return err
		}
		w.initialized = true
	}
	// ak sme dosiali limit pre zapis do suboru
	// aktualny subor zatvorime a otvorime novy
	if w.rowsWritten == MaxWrite {
		if err := w.close(); err != nil {
			return err
		}
		if err := w.add(); err != nil {
			return err
		}
	}

	// zapis csv
	if err := w.writer.Write(record); err != nil {
		return err
	}

	// aktualny pocet zapisanych do suboru
	w.rowsWritten++

	return nil
}

// Close zatvori subor a zapise vsetky data v buffer
func (w *CsvMaxWriter) Close() error {
	if !w.initialized {
		return nil
	}
	return w.close()
}

func (w *CsvMaxWriter) close() error {

	if w.file == nil {
		return errors.New("CsvMaxWriter: close: pokus o zatvorenie nevytvoreneho suboru")
	}

	if w.writer == nil {
		return errors.New("CsvMaxWriter: close: writer nema byt nil")
	}

	writer := w.writer

	// vynulovanie writera
	w.writer = nil

	// flushnutie buffera
	writer.Flush()

	// kontrola chyby
	if err := writer.Error(); err != nil {
		return err
	}

	// celkom zapisanych riadkov
	w.totalRowsWritten += w.rowsWritten

	// resetneme pocet zapisanych riadkov v aktualnom subore
	// hlavicka sa nerata
	w.rowsWritten = 0

	// if err := w.file.Sync(); err != nil {
	// 	return err
	// }

	file := w.file

	// vynulovanie suboru
	w.file = nil

	// zatvorenie suboru
	return file.Close()
}

func (w *CsvMaxWriter) add() (err error) {

	if w.file != nil {
		return errors.New("CsvMaxWriter: add: subor nebol zatvoreny")
	}

	// nazov dalsieho suboru
	filename := w.nextFilename()

	// vytvorime subor
	w.file, err = os.Create(filename)
	if err != nil {
		return
	}

	// zapiseme BOM, kvoli MS Excelu. TODO: UTF-8 BOM... :( remove this...
	_, err = w.file.Write([]byte{0xEF, 0xBB, 0xBF})
	if err != nil {
		return
	}

	if w.writer != nil {
		return errors.New("CsvMaxWriter: add: csv writer ma byt nil")
	}
	// nastavime csvWriter
	w.writer = csv.NewWriter(w.file)

	// nastavime oddelovac
	w.writer.Comma = Comma

	// zapisem hlavicku
	return w.writer.WriteAll(w.header)
}

func (w *CsvMaxWriter) nextFilename() string {

	// specialny pripad ked nezapisujeme nic
	if w.workingDir == DiscardCSV {
		return os.DevNull
	}

	// vrati nazov suboru v tvare workingDir/subDir/fileName_suffix.csv
	// ak je subDir "" tak workingDir/fileName_suffix.csv
	return filepath.Join(w.workingDir, w.subdir, w.fileName+w.Suffix()+".csv")
}
