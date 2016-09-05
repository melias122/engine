package csv

import (
	"encoding/csv"
	"io"
)

type Reader struct {
	*csv.Reader
}

func NewReader(r io.Reader) *Reader {
	re := csv.NewReader(r)
	re.Comma = ';'
	return &Reader{
		Reader: re,
	}
}
