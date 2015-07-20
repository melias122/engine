package filter

import "github.com/melias122/psl/komb"

type Filter interface {
	Check(komb.Kombinacia) bool
}
