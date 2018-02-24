package engine // import "github.com/melias122/engine/engine"

// Kombinacia reprezentuje kombinaciu cisiel.
// Cisla musia byt vacsie ako 1 a mensie ako 'm'. Velkost kombinacie musi byt 'n'
// a nesmie sa menit.
type Kombinacia []int

// Rc je cislo c. c moze mat hodnotu od 1 do n.
type Rc interface {
	// Rp vrati pocetnost cisla c.
	Rp(c int) int

	// Rh vrati hodnotu cisla c vypocitanu pomocou vzorca H(c, 1, Rp(c), n, m).
	Rh(c int) float64
}

// STLc je cislo c v stlpci s.
type STLc interface {
	// STLp vrati pocetnost cisla c.
	STLp(c int, s int) int

	// STLh vypocita hodnotu cisla pomocou vzorca H(c, s, STLp(c), n, m).
	STLh(c int, s int) float64
}

// Cislo je cislo v riadku a slpci s hodnotami R a STL.
type Cislo interface {
	Rc
	STLc
}

// Rk definuje sucet hodnot Rp.
type Rk interface {
	// R vrati sucet hodnot Rp Kombinacie.
	R(Kombinacia) float64
}

// STLk definuje sucet hodnot STLp.
type STLk interface {
	// STL vrati sucet hodnot STLp Kombinacie.
	STL(Kombinacia) float64
}

// RSTLk kombinuje sucty R a STL.
type RSTLk interface {
	Rk
	STLk
}

// Xk vypocita hodnotu Hrx, pripadne HHrx.
type Xk interface {
	// Ak je zadana Kombinacia k vypocita hodnotu po presune cisiel.
	// V pripade k == nil, vrati aktualnu hodnotu hrx, hhrx.
	X(k Kombinacia) float64
}

// Hrx je rozhranie ktore zdruzuje hodnoty R a STL vypocitane pre 1-DO a OD-DO.
type Hrx interface {
	Cislo
	RSTLk
	Xk

	Uc() (riadok, cislo int)
}

type HHrx interface {
	Cislo
	RSTLk
	Xk
}
