package list

import (
	"fmt"
)

// The default LinkedList implementation
type GenericLinkedList struct {
	len  int
	head *node
	tail *node
}

func NewGenericLinkedList() *GenericLinkedList {
	return new(GenericLinkedList)
}

// Return the length of the linkedlist
func (l *GenericLinkedList) Len() int {
	return l.len
}

// Add an element to the list
func (l *GenericLinkedList) Add(v interface{}) {
	newTail := newNode(v)
	if l.len == 0 {
		l.head = newTail
		l.tail = newTail
	} else {
		l.tail.next = newTail
		l.tail = newTail
	}
	l.len++
	return
}

// Get the value of the node at n, start from 0
func (l *GenericLinkedList) Get(n int) interface{} {
	return l.get(n).val
}

// Get the node at index n, starts from 0
func (l *GenericLinkedList) get(n int) *node {
	if n < 0 || n >= l.len {
		panic(fmt.Sprintf("index out of bound: %d", n))
	}
	cur := l.head
	for i := 0; i < n; i++ {
		cur = cur.next
	}
	return cur
}

// Delete the node at index n and return the val of the node
func (l *GenericLinkedList) Del(n int) interface{} {
	return l.del(n).val
}

// Delete the node at index n and return it
func (l *GenericLinkedList) del(n int) *node {
	if n < 0 || n >= l.len {
		panic(fmt.Sprintf("index out of bound: %d", n))
	}
	var delNode *node
	if n == 0 {
		delNode = l.head
		l.head = l.head.next
		if l.len == 1 {
			l.tail = nil
		}
	} else {
		lastOne := l.get(n - 1)
		delNode = lastOne.next
		lastOne.next = delNode.next
		if n == l.len-1 {
			l.tail = lastOne
		}
	}
	delNode.next = nil // avoid memory leak
	l.len--
	return delNode
}

// List node
type node struct {
	next *node
	val  interface{}
}

func newNode(v interface{}) *node {
	return &node{
		val: v,
	}
}

func (n *node) value() interface{} {
	return n.val
}

type GenericLinkedListIterator struct {
	index int
	cur   *node
}

func (l *GenericLinkedList) NewIterator() *GenericLinkedListIterator {
	return &GenericLinkedListIterator{-1, l.head}
}

func (i *GenericLinkedListIterator) HasNext() bool {
	return i.cur != nil && (i.index == -1 || i.cur.next != nil)
}

func (i *GenericLinkedListIterator) Next() *GenericLinkedListIterator {
	if i.index >= 0 {
		if !i.HasNext() {
			panic(fmt.Sprintf("index out of bound: %d", i.index))
		}
		i.cur = i.cur.next
	}
	i.index++
	return i
}

func (i *GenericLinkedListIterator) Index() int {
	return i.index
}

func (i *GenericLinkedListIterator) GetValue() interface{} {
	if i.index < 0 {
		panic("call Next() to get the first element")
	}
	if i.cur == nil {
		panic(fmt.Sprintf("index out of bound: %d", i.index))
	}
	return i.cur.value()
}
