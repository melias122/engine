package filter

type Filter interface {
	Check([]int) bool
}
