package engine

type Parser interface {
	Parse() ([]Kombinacia, error)
}

type (
	Kombinacia []int

	Cislovacka int

	Cifrovacky [10]byte
)
