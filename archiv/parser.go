package archiv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type ErrKomb struct {
	Komb []int
	Err  error
}

func Parse(path string, n, m int) chan ErrKomb {

	ch := make(chan ErrKomb, 8)
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
		r.Read()

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
			komb := make([]int, n)
			for i, field := range record[3 : n+3] {
				cislo, err := strconv.Atoi(field)
				if err != nil {
					ch <- ErrKomb{Err: err}
					return
				}
				komb[i] = cislo
			}
			ch <- ErrKomb{Komb: komb}
		}
	}()
	return ch
}
