package parser

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
	P
	N
	Pr
	Mc
	Vc
	C19
	C0
	CA // cC
	CB // Cc
	CC // CC
	Zh

	// Misc
	DASH      // "-"
	COMMA     // ","
	COLON     // ":"
	SEMICOLON // ";"
)
