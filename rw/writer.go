package rw

import (
	"csv"
	"os"
	"path/filepath"
)

const (
	// Standardne pouzivame ako oddelovat ";"
	Comma = ';'

	// Maximum Excel rows..
	MaxWrite = 5 * 10e4
)

var (
	wdir string
)

func SetWDir(path string) {
	wdir = path
}

type Writer struct {
	filename string
	file     *os.File
	w        *csv.Writer
}

func NewWriter(FileName string) (*Writer, error) {
	file, err := os.Create(filepath.Join(wdir, FileName+".csv"))
	if err != nil {
		return nil, err
	}
	w := csv.Writer(file)
	w.Comma = Comma
	return &Writer{file, w}, nil
}

func (w *Writer) Write(record []string) error {
	if err := w.w.Write(record); err != nil {
		return err
	}
	return nil
}

func (w *Writer) Flush() error {
	return w.w.Flush()
}

func (w *Writer) Close() error {
	return w.file.Close()
}
