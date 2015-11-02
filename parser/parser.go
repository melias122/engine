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
	Zhoda []byte
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		s: NewScanner(r),
		// n: m,
		// m: m,
	}
}

func (p *Parser) ParseInt() (int, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != DIGIT {
		return 0, fmt.Errorf("found %q, expected digit", lit)
	}

	// Parse number
	return strconv.Atoi(lit)
	// number, err := strconv.Atoi(lit)
	// if err != nil {
	// 	return 0, err
	// }
	// if number < 1 || number > p.m {
	// 	return 0, fmt.Errorf("expected number in range 1..%d", p.m)
	// }
	// return number, nil
}

func (p *Parser) ParseInts() (Ints, error) {
	var s Ints

	// loop over colon delimited fields
	// expect digit or P,..,Zh on first place
	for {
		tok, _ := p.scanIgnoreWhitespace()
		switch tok {
		case DIGIT:
			// unscan first
			p.unscan()
			// parse digit
			i, err := p.ParseInt()
			if err != nil {
				return nil, err
			}
			s = append(s, i)

			// case DIGIT1-DIGIT2
			if tok, _ := p.scanIgnoreWhitespace(); tok == DASH {
				// parse DIGIT2
				j, err := p.ParseInt()
				if err != nil {
					return nil, err
				}
				if i > j {
					return nil, fmt.Errorf("%d > %d", i, j)
				}
				// fill range
				for i := i + 1; i <= j; i++ {
					s = append(s, i)
				}
			} else {
				// case DIGIT,
				p.unscan()
			}

		case P, N, Pr, Mc, Vc, C19, C0, CA, CB, CC:
			s = append(s, cislovacky[tok]...)
			// for _, num := range cislovacky[tok] {
			// 	if num > p.M {
			// 		break
			// 	}
			// s = append(s, num)
			// }

		case Zh:
			for _, num := range p.Zhoda {
				num := int(num)
				s = append(s, num)
			}
		default:
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
		// if i < 1 || i > p.n {
		// 	return nil, fmt.Errorf("STL expected to be in range 1..%d", p.n)
		// }
		m[i] = []int{}
		// next should be colon(:)
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

// lookup table for cislovacky
var cislovacky = [...][]int{
	P:   {2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98},
	N:   {1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57, 59, 61, 63, 65, 67, 69, 71, 73, 75, 77, 79, 81, 83, 85, 87, 89, 91, 93, 95, 97, 99},
	Pr:  {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
	Mc:  {1, 2, 3, 4, 5, 11, 12, 13, 14, 15, 21, 22, 23, 24, 25, 31, 32, 33, 34, 35, 41, 42, 43, 44, 45, 51, 52, 53, 54, 55, 61, 62, 63, 64, 65, 71, 72, 73, 74, 75, 81, 82, 83, 84, 85, 91, 92, 93, 94, 95},
	Vc:  {6, 7, 8, 9, 10, 16, 17, 18, 19, 20, 26, 27, 28, 29, 30, 36, 37, 38, 39, 40, 46, 47, 48, 49, 50, 56, 57, 58, 59, 60, 66, 67, 68, 69, 70, 76, 77, 78, 79, 80, 86, 87, 88, 89, 90, 96, 97, 98, 99},
	C19: {1, 2, 3, 4, 5, 6, 7, 8, 9},
	C0:  {10, 20, 30, 40, 50, 60, 70, 80, 90},
	CA:  {12, 13, 14, 15, 16, 17, 18, 19, 23, 24, 25, 26, 27, 28, 29, 34, 35, 36, 37, 38, 39, 45, 46, 47, 48, 49, 56, 57, 58, 59, 67, 68, 69, 78, 79, 89},
	CB:  {21, 31, 32, 41, 42, 43, 51, 52, 53, 54, 61, 62, 63, 64, 65, 71, 72, 73, 74, 75, 76, 81, 82, 83, 84, 85, 86, 87, 91, 92, 93, 94, 95, 96, 97, 98},
	CC:  {11, 22, 33, 44, 55, 66, 77, 88, 99},
}
