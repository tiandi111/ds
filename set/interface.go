package set

type UnionFindSet interface {
	Find(int) int
	Union(a, b int)
	Connected(a, b int) bool
	Count() int
}
