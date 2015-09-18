package archiv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/melias122/psl/komb"
)

type ErrKomb struct {
	Header []string
	Komb   komb.Kombinacia
	Orig   []string
	Err    error
}

func parse(record []string, n int) (komb.Kombinacia, error) {
	var (
		komb  = make([]byte, n)
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
		komb[i] = byte(cislo)
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
		r.Comma = rune(';')

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
