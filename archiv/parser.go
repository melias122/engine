package archiv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func Parse(path string, n, m int) ([][]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return parse(f, n, m)
}

func parse(f *os.File, n, m int) ([][]byte, error) {
	r := csv.NewReader(f)
	r.Comma = rune(';')

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	var k [][]byte
	for i, rec := range records[1:] {
		if len(rec) < n+3 {
			return nil,
				fmt.Errorf("%s: na riadku %d sa nepodarilo nacitat kombinaciu", f.Name(), i+1)
		}
		c := make([]byte, n)
		for i, x := range rec[3 : n+3] {
			cislo, err := strconv.ParseUint(x, 10, 0)
			if err != nil {
				return nil, err
			}
			c[i] = byte(cislo)
		}

		k = append(k, c)
	}
	return k, nil
}
