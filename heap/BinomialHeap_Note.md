# Binomial Heap Note

## Binomial Tree
The binomial Tree B[k] consists of two binomial tree B[k-1] that are linked together: the root of one is the leftmost child of the root of the other<br/>

## Structure
```go
type node struct {
	parent  *node
	sibling *node
	child   *node // left most child
	degree  int   // the number of children
	val     ds.Comparable
}
```

## Properties
1. A set of binomial trees<br/>
2. Each binomial tree in H obeys the min-heap property<br/>
3. For any non-negative integer, there is at most one binomial tree in H whose root has degree k<br/> 