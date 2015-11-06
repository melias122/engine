package psl

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
)

const (
	// Standardne pouzivame ako oddelovat ";"
	Comma = ';'

	// Maximum Excel rows..
	MaxWrite = (5 * 10e4) - 1
)

func IntSuffix() func() string {
	var nFile int
	return func() string {
		nFile++
		return "_" + strconv.Itoa(nFile)
	}
}

func EmptySuffix() string {
	return ""
}

type CsvMaxWriter struct {
	count, max    int
	dir, fileName string
	header        [][]string
	file          *os.File
	writer        *csv.Writer
	Suffix        func() string
	SubDir        string
	NWrites       uint64
}

func NewCsvMaxWriter(dir, fileName string, header [][]string) *CsvMaxWriter {
	return &CsvMaxWriter{
		max:      MaxWrite,
		dir:      dir,
		fileName: fileName,
		header:   header,
		Suffix: func() string {
			return "_" + filepath.Base(dir)
		},
	}
}

func (w *CsvMaxWriter) Reset() error {
	if w.file != nil {
		if err := w.Close(); err != nil {
			return err
		}
	}
	w.count = 0
	if w.SubDir != "" {
		os.Mkdir(filepath.Join(w.dir, w.SubDir), 0755)
	}
	f, err := os.Create(filepath.Join(w.dir, w.SubDir, w.fileName+w.Suffix()+".csv"))
	if err != nil {
		return err
	}
	w.file = f
	w.file.Write([]byte{0xEF, 0xBB, 0xBF}) // TODO: UTF-8 BOM... :( remove this...
	w.writer = csv.NewWriter(w.file)
	w.writer.Comma = Comma
	return w.writer.WriteAll(w.header)
}

func (w *CsvMaxWriter) Close() error {
	if w.file != nil {
		defer w.file.Close()
		w.writer.Flush()
		return w.writer.Error()
	}
	return nil
}

func (w *CsvMaxWriter) Write(record []string) error {
	if w.file == nil || w.writer == nil || w.count > w.max {
		if err := w.Reset(); err != nil {
			return err
		}
	}
	w.count++
	w.NWrites++

	return w.writer.Write(record)
}
