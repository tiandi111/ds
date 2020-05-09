package list

import "fmt"

// TODO: implement list with a sentinel node
// Circular Doubly Linked List implementation
type GenericDoublyLinkedList struct {
	head *DllNode
	tail *DllNode
	len  int
}

func NewGenericDoublyLinkedList() *GenericDoublyLinkedList {
	return new(GenericDoublyLinkedList)
}

func (dl *GenericDoublyLinkedList) Add(val interface{}) {
	node := NewDllNode(nil, nil, val)
	if dl.tail == nil {
		node.next = node
		node.prev = node
		dl.head = node
		dl.tail = node
	} else {
		dl.tail.next = node
		node.prev = dl.tail
		dl.head.prev = node
		node.next = dl.head
	}
	dl.tail = node
	dl.len++
}

func (dl *GenericDoublyLinkedList) InsertAfter(node, after *DllNode) {
	node.next = after.next
	node.prev = after
	after.next.prev = node
	after.next = node
	if after == dl.tail {
		dl.tail = node
	}
	dl.len++
}

func (dl *GenericDoublyLinkedList) Get(n int) interface{} {
	return dl.GetNode(n).Value()
}

func (dl *GenericDoublyLinkedList) GetNode(n int) *DllNode {
	if n < 0 || n >= dl.len {
		panic(fmt.Sprintf("index out of bound: %d", n))
	}
	node := dl.head
	for i := 0; i < n; i++ {
		node = node.Next()
	}
	return node
}

func (dl *GenericDoublyLinkedList) Del(n int) interface{} {
	delNode := dl.GetNode(n)
	dl.DelNode(delNode)
	return delNode.Value()
}

// TODO: check if n belongs to dl
func (dl *GenericDoublyLinkedList) DelNode(n *DllNode) {
	if dl.head == n && dl.tail == n {
		dl.head = nil
		dl.tail = nil
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
		if dl.head == n {
			dl.head = n.next
		}
		if dl.tail == n {
			dl.tail = n.prev
		}
	}
	// avoid memory leak!
	n.next = nil
	n.prev = nil
	n.val = nil
	dl.len--
}

func (dl *GenericDoublyLinkedList) Len() int {
	return dl.len
}

type DllNode struct {
	prev *DllNode
	next *DllNode
	val  interface{}
}

func NewDllNode(prev, next *DllNode, val interface{}) *DllNode {
	return &DllNode{prev, next, val}
}

func (n *DllNode) Prev() *DllNode {
	return n.prev
}

func (n *DllNode) Next() *DllNode {
	return n.next
}

func (n *DllNode) Value() interface{} {
	return n.val
}
