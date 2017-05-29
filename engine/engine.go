package engine

type Parser interface {
	Parse() ([]Kombinacia, error)
}
