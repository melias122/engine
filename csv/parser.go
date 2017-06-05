package csv

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/melias122/engine/engine"
	"github.com/pkg/errors"
)

func ParsePath(path string, n, m int) ([]engine.Kombinacia, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewParser(f, n, m).Parse()
}

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

func (p *Parser) Parse() ([]engine.Kombinacia, error) {

	var (
		k    []engine.Kombinacia
		line = 1
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

		new := make(engine.Kombinacia, p.n)
		for i, field := range row[3 : p.n+3] {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, errors.Wrapf(err, "on line %v", line)
			}
			if num < 1 || num > p.m {
				return nil, fmt.Errorf("on line %v: %v is not valid number", line, num)
			}
			new[i] = num
		}
		k = append(k, new)
	}

	if len(k) > 0 {
		return k, nil
	}

	return nil, fmt.Errorf("could not parse file")
}
