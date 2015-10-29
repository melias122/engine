package parser

import (
	"fmt"
	"io"
	"strconv"
)

type Ints []int
type MapInts map[int]Ints

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) ParseInt() (int, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != DIGIT {
		return 0, fmt.Errorf("found %q, expected digit", lit)
	}
	// Parse number
	return strconv.Atoi(lit)
}

func (p *Parser) ParseInts() (Ints, error) {
	var s Ints

	// loop over colon delimited fields
	for {
		i, err := p.ParseInt()
		if err != nil {
			return nil, err
		}
		s = append(s, i)

		if tok, _ := p.scanIgnoreWhitespace(); tok == DASH {
			j, err := p.ParseInt()
			if err != nil {
				return nil, err
			}
			if i > j {
				return nil, fmt.Errorf("%d > %d", i, j)
			}
			for i := i + 1; i <= j; i++ {
				s = append(s, i)
			}
		} else {
			p.unscan()
		}
		if tok, _ := p.scanIgnoreWhitespace(); tok != COMMA {
			p.unscan()
			break
		}
	}
	return s, nil
}

func (p *Parser) ParseMapInts() (MapInts, error) {
	var m = make(MapInts)

	// loop over semicolon delimited fields
	for {
		// first we want digit
		i, err := p.ParseInt()
		if err != nil {
			return nil, err
		}
		m[i] = []int{}
		// next should be colon
		if tok, lit := p.scanIgnoreWhitespace(); tok != COLON {
			return nil, fmt.Errorf("found %q, expected :", lit)
		}
		// next we want comma separated list
		if s, err := p.ParseInts(); err != nil {
			return nil, err
		} else {
			m[i] = append(m[i], s...)
		}
		if tok, _ := p.scanIgnoreWhitespace(); tok != SEMICOLON {
			p.unscan()
			break
		}
	}
	return m, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}
