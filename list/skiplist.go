package list

import (
	"fmt"
	"github.com/tiandi111/ds"
	"log"
	"math/rand"
)

type GenericSkipList struct {
	// the head node of the top-most level
	head *slNode
	// average times to up a node's level
	// avgUpLevelTries = 1 / p
	// p is the probability of upgrading a node's level
	avgUpLevelTries int
	len             int
}

// default avgUpLevelTries = 4
func NewGenericSkipList() *GenericSkipList {
	return &GenericSkipList{nil, 4, 0}
}

func NewGenericSkipListWithN(n int) *GenericSkipList {
	if n < 0 {
		panic(errInvalidSkipListP())
	}
	sl := NewGenericSkipList()
	sl.avgUpLevelTries = n
	return sl
}

func (l *GenericSkipList) Add(v ds.Comparable) {
	// random generated level
	// P(level == 1) = 1-p
	// P(level == 2) = p*(1-p)
	// p(level == 3) = p*p*(1-p)
	// ...
	level := l.level()
	if l.head == nil || level > l.head.maxLevel {
		// construct higher levels on the old head node
		l.head = growHeadSlNode(level, l.head)
	}
	l.head.insert(level, v)
	l.len++
}

// get the first occurrence of v
func (l *GenericSkipList) Get(v ds.Comparable) ds.Comparable {
	if l.head == nil {
		return nil
	}
	node := l.head.findExact(v)
	if node == nil {
		return nil
	}
	return node.value()
}

// delete the first occurrence of v
func (l *GenericSkipList) Del(v ds.Comparable) ds.Comparable {
	if l.head == nil {
		return nil
	}
	node := l.head.del(v)
	// no such node
	if node == nil {
		return nil
	}
	l.len--
	return node.value()
}

func (l *GenericSkipList) Len() int {
	return l.len
}

func (l *GenericSkipList) MaxLevel() int {
	return l.head.maxLevel
}

func (l *GenericSkipList) Info() {
	fmt.Printf("Size: %d MaxLevel: %d \n", l.Len(), l.MaxLevel())
	for i := l.MaxLevel(); i >= 0; i-- {
		n := l.head
		cnt := 0
		for n != nil {
			cnt++
			n = n.next[i]
		}
		fmt.Printf("Number of nodes at level %dth: %d\n", i, cnt)
	}
}

// generate a level by p = 1/avgUpLevelTries
func (l *GenericSkipList) level() int {
	level := 0
	for {
		if rand.Intn(l.avgUpLevelTries) > 0 {
			break
		}
		level++
	}
	return level
}

type slNode struct {
	// pointers array
	next     []*slNode
	val      ds.Comparable
	isHead   bool
	maxLevel int
}

// newHeadSlNode construct higher levels on the oldHead
func growHeadSlNode(level int, oldHead *slNode) *slNode {
	if level < 0 {
		panic(errInvaldSkipListHeadNodeLevel())
	}
	if oldHead == nil {
		return &slNode{make([]*slNode, 2*(level+1)), nil, true, level}
	}
	if cap(oldHead.next) <= level {
		newNext := make([]*slNode, 2*(level+1))
		copy(newNext, oldHead.next)
		oldHead.next = newNext
		oldHead.maxLevel = level
	}
	return oldHead
}

// insert the given value under the given level
func (n *slNode) insert(level int, v ds.Comparable) {
	node := &slNode{make([]*slNode, level+1), v, false, level}
	this := n
	last := n
	for i := this.maxLevel; i >= 0; i-- {
		for this != nil && (this.isHead || this.val.CompareTo(v) <= 0) {
			last = this
			this = this.next[i]
		}
		if i <= level {
			tmp := last.next[i]
			last.next[i] = node
			node.next[i] = tmp
		}
		this = last
	}
}

// find the insert position at the given level
// the insert position is the largest smaller node
func (n *slNode) find(level int, v ds.Comparable) *slNode {
	node := n
	last := n
	for i := node.maxLevel; i >= level; i-- {
		for node != nil && (node.isHead || node.val.CompareTo(v) < 0) {
			last = node
			node = node.next[i]
		}
		node = last
	}
	return last
}

// find the exact node
// if no such node, return nil
func (n *slNode) findExact(v ds.Comparable) *slNode {
	node := n.find(0, v)
	if node.next[0] != nil && node.next[0].value().CompareTo(v) == 0 {
		return node.next[0]
	}
	return nil
}

func (n *slNode) del(v ds.Comparable) *slNode {
	node := n
	last := n
	var delNode *slNode
	for i := node.maxLevel; i >= 0; i-- {
		for node != nil && (node.isHead || node.val.CompareTo(v) < 0) {
			last = node
			node = node.next[i]
		}
		if delNode == nil && node != nil && !node.isHead && node.value().CompareTo(v) == 0 {
			delNode = node
			last.next[i] = node.next[i]
		} else if delNode != nil && delNode == node {
			last.next[i] = node.next[i]
		}
		node = last
	}
	return delNode
}

func (n *slNode) value() ds.Comparable {
	return n.val
}

// validate node sequence
func (n *slNode) validate(enableLog bool) bool {
	for i := n.maxLevel; i >= 0; i-- {
		var last *slNode
		node := n
		msg := ""
		for node != nil {
			if enableLog {
				msg += fmt.Sprintf("%v", node.val)
			}
			if last != nil && !last.isHead && last.value().CompareTo(node.val) > 0 {
				return false
			}
			last = node
			node = node.next[i]
		}
		if enableLog {
			log.Printf("the %dth level: %s\n", i, msg)
		}
	}
	return true
}

func (l *GenericSkipList) NewIterator() ds.Iterator {
	return nil
}

type SkipListIterator struct {
}
