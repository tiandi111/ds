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

func NewGenericSkipListWithP(p float32) *GenericSkipList {
	if p < 0 || p > 1 {
		panic(errInvalidSkipListP())
	}
	sl := NewGenericSkipList()
	sl.avgUpLevelTries = int(1 / p)
	return sl
}

func (l *GenericSkipList) Add(v ds.Comparable) {
	level := l.level()
	if l.head == nil || level > l.head.maxLevel {
		// construct more level on the head node
		l.head = growHeadSlNode(level, l.head)
	}
	i := l.head.find(level, v)
	i.insert(level, v)
	l.len++
}

func (l *GenericSkipList) Get(v ds.Comparable) ds.Comparable {
	return l.head.find(0, v).value()
}

// delete the first occurrence of v
func (l *GenericSkipList) Del(v ds.Comparable) ds.Comparable {
	if l.head == nil {
		return nil
	}
	return l.head.del(v).value()
}

func (l *GenericSkipList) Len() int {
	return l.len
}

func (l *GenericSkipList) MaxLevel() int {
	return l.head.maxLevel
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

// newHeadSlNode construct more level on the oldHead
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

// first find the insert position, then insert the value node all the way to the bottom
func (n *slNode) insert(level int, v ds.Comparable) {
	node := &slNode{make([]*slNode, len(n.next)), v, false, n.maxLevel}
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
// node will be inserted at the largest smaller node or head node
func (n *slNode) find(level int, v ds.Comparable) *slNode {
	node := n
	last := n
	for i := node.maxLevel; i >= level; i-- {
		for node != nil && (node.isHead || node.val.CompareTo(v) <= 0) {
			last = node
			node = node.next[i]
		}
	}
	return last
}

func (n *slNode) del(v ds.Comparable) *slNode {
	node := n
	last := n
	var delNode *slNode
	for i := node.maxLevel; i >= 0; i-- {
		for node != nil && (node.isHead || node.val.CompareTo(v) < 0) {
			node = node.next[i]
			last = node
		}
		if node != nil && node.val.CompareTo(v) == 0 {
			last.next[i] = node.next[i]
			delNode = node
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
