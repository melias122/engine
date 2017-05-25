package csv

import (
	"fmt"
	"io"
	"strconv"

	"github.com/pkg/errors"
)

type Parser struct {
	r          *Reader
	n, m       int
	SkipHeader bool
}

func NewParser(r io.Reader, n, m int) *Parser {

	re := NewReader(r)
	re.FieldsPerRecord = -1

	return &Parser{
		r:          re,
		n:          n,
		m:          m,
		SkipHeader: true,
	}
}

func (p *Parser) Parse() ([][]int, error) {

	var (
		combs [][]int
		line  = 1
	)

	for ; ; line++ {

		row, err := p.r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, errors.Wrap(err, "read failed")
		}

		if p.SkipHeader && line == 1 {
			continue
		}

		if len(row) < p.n+3 {
			return nil, fmt.Errorf("not enough fields on line %d", line)
		}

		comb := make([]int, p.n)
		for i, field := range row[3 : p.n+3] {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, errors.Wrapf(err, "on line %v", line)
			}
			comb[i] = num
		}
		combs = append(combs, comb)
	}

	if len(combs) > 0 {
		return combs, nil
	}

	return nil, fmt.Errorf("could not parse file")
}
