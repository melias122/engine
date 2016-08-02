package engine

type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS // whitespace

	// Literals
	IDENT
	DIGIT

	// cislovacky
	p
	n
	pr
	mc
	vc
	c19
	c0
	ca // cC
	cb // Cc
	cc // CC
	zh

	// Misc
	DASH      // "-"
	COMMA     // ","
	COLON     // ":"
	SEMICOLON // ";"
)
