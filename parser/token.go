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

	// Misc
	DASH      // "-"
	COMMA     // ","
	COLON     // ":"
	SEMICOLON // ";"
)
