package linkedlist

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

func (n *Node[T]) Value() T {
	return n.data
}

type LinkedList[T any] struct {
	length int
	head   *Node[T]
	end    *Node[T]
}

func NewLinkedList[T any](value T) *LinkedList[T] {
	node := Node[T]{
		data: value,
		next: nil,
		prev: nil,
	}
	return &LinkedList[T]{
		head:   &node,
		end:    &node,
		length: 1,
	}
}

func LinkedListFromSlice[T any](slice []T) *LinkedList[T] {
	ll := NewLinkedList[T](slice[0])
	for i := 1; i < len(slice); i++ {
		ll.AppendEnd(slice[i])
	}
	return ll
}

func (ll *LinkedList[T]) AppendEnd(value T) {
	node := Node[T]{
		data: value,
		next: nil,
		prev: ll.end,
	}
	ll.end.next = &node
	ll.end = &node
	ll.length++
}

func (ll *LinkedList[T]) AppendHead(value T) {
	node := Node[T]{
		data: value,
		next: ll.head,
		prev: nil,
	}
	ll.head.prev = &node
	ll.head = &node
	ll.length++
}

func (ll *LinkedList[T]) Len() int {
	return ll.length
}

func (ll *LinkedList[T]) Contains(value T, compFn func(T, T) bool) bool {
	node := ll.head
	for node != nil {
		if compFn(node.data, value) {
			return true
		}
		node = node.next
	}
	return false
}

// returns node at certain index.
// positive integers to start at head
// negative to start at end
func (ll *LinkedList[T]) at(ix int) *Node[T] {

	if ix >= ll.length {
		return nil
	}

	var pos int
	var ptr *Node[T]
	if ix >= 0 {
		pos = 0
		ptr = ll.head
		for pos != ix && ptr != nil {
			ptr = ptr.next
			pos++
		}
	}
	if ix < 0 {
		ix = -ix
		pos = 1
		ptr = ll.end
		for pos != ix && ptr != nil {
			ptr = ptr.prev
			pos++
		}
	}
	return ptr
}

// returns node at certain index.
// positive integers to start at head
// negative to start at end
func (ll *LinkedList[T]) At(ix int) *Node[T] {
	return ll.at(ix)
}

// as in At(..) function - positive index means starting from head
// and negative - from end
func (ll *LinkedList[T]) DeleteAt(ix int) error {
	node := ll.at(ix)
	if node == nil {
		return fmt.Errorf("no node at index %d", ix)
	}

	if node == ll.head {
		ll.head = node.next
	} else if node == ll.end {
		ll.end = node.prev
	}

	if node.next != nil {
		node.next.prev = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	ll.length--
	return nil
}

func (ll *LinkedList[T]) Slice() []T {
	slice := make([]T, ll.length)
	for i := 0; i < ll.length; i++ {
		slice[i] = ll.at(i).data
	}
	return slice
}
