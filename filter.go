package engine

type Filter interface {
	String() string
	Check(Kombinacia) bool
	CheckSkupina(Skupina) bool

	// used for sorting by filter priority
	// because some filters are faster
	// and have to be exectuted sooner than others
	priority() filterPriority
}

type Filters []Filter

func (f Filters) Check(k Kombinacia) bool {
	for _, filter := range f {
		if !filter.Check(k) {
			return false
		}
	}
	return true
}

func (f Filters) CheckSkupina(s Skupina) bool {
	for _, filter := range f {
		if !filter.CheckSkupina(s) {
			return false
		}
	}
	return true
}

func (f Filters) String() string {
	var s string
	for _, filter := range f {
		s += filter.String() + "\n"
	}
	return s
}

type filterPriority int

func (f filterPriority) priority() filterPriority {
	return f
}

const (
	P0 filterPriority = iota
	P1
	P2
	P3
	P4
)

type byPriority Filters

func (b byPriority) Len() int           { return b.Len() }
func (b byPriority) Less(i, j int) bool { return b[i].priority() < b[j].priority() }
func (b byPriority) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
