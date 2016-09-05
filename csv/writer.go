package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

const (
	// Standardne pouzivame ako oddelovat ";"
	Comma = ';'

	// Maximum Excel rows..
	MaxWrite = 500000
)

type SuffixFunc func() string

func IntSuffix() SuffixFunc {
	var i int
	return func() string {
		i++
		return "_" + strconv.Itoa(i)
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

	error error
}

func SetSuffixFunc(sf SuffixFunc) func(w *CsvMaxWriter) {
	return func(w *CsvMaxWriter) {
		w.Suffix = sf
	}
}

func SetSubdir(subdir string) func(w *CsvMaxWriter) {
	return func(w *CsvMaxWriter) {
		w.subdir = subdir
	}
}

func SetHeader(header []string) func(w *CsvMaxWriter) {
	return func(w *CsvMaxWriter) {
		w.header = [][]string{header}
	}
}

func SetHeaders(headers [][]string) func(w *CsvMaxWriter) {
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
	if err := os.MkdirAll(filepath.Join(w.workingDir, w.subdir), 0755); err != nil {
		w.error = err
	}
	return w
}

// TotalRowsWriten vrati pocet zapisanych riadkov
func (w *CsvMaxWriter) TotalRowsWriten() int {
	return w.totalRowsWritten
}

// Write zapise retazce record do suboru
func (w *CsvMaxWriter) Write(record []string) error {

	// uplne prvy zapis ked este nebol vytvoreny subor
	// je potrebne inicializovat file a writer
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

	if w.writer == nil {
		panic("nil writer")
	}

	// zapis csv
	if err := w.writer.Write(record); err != nil {
		return err
	}
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

func (w *CsvMaxWriter) close() (err error) {

	if w.file == nil {
		return errors.New("CsvMaxWriter: close: pokus o zatvorenie nevytvoreneho suboru")
	}

	if w.writer == nil {
		return errors.New("CsvMaxWriter: close: writer nema byt nil")
	}

	// defer aby sme zakazdym zatvorenim urcite zavreli subor
	defer func() {
		// zatvorenie suboru
		if e := w.file.Close(); e != nil {
			err = e
		}

		// vynulovanie suboru
		w.file = nil
	}()

	// flushnutie buffera
	w.writer.Flush()

	// kontrola chyby
	if err = w.writer.Error(); err != nil {
		return
	}

	// vynulovanie writera
	w.writer = nil

	// celkom zapisanych riadkov
	w.totalRowsWritten += w.rowsWritten

	// resetneme pocet zapisanych riadkov v aktualnom subore
	// hlavicka sa nerata
	w.rowsWritten = 0

	return
}

func (w *CsvMaxWriter) add() (err error) {

	if w.file != nil {
		return errors.New("CsvMaxWriter: add: subor nebol zatvoreny")
	}

	// nazov dalsieho suboru
	filename := w.nextFilename()

	// vytvorime subor
	if w.file, err = os.Create(filename); err != nil {
		return
	}

	// zapiseme BOM, kvoli MS Excelu. TODO: UTF-8 BOM... :( remove this...
	if _, err = w.file.Write([]byte{0xEF, 0xBB, 0xBF}); err != nil {
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
	if w.header != nil {
		err = w.writer.WriteAll(w.header)
	}
	return
}

func (w *CsvMaxWriter) nextFilename() string {
	// vrati nazov suboru v tvare workingDir/subDir/fileName_suffix.csv
	// ak je subDir "" tak workingDir/fileName_suffix.csv
	return filepath.Join(w.workingDir, w.subdir, w.fileName+w.Suffix()+".csv")
}
