package linkedlist

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

func (ll *LinkedList[T]) AppendEnd(value T) {
	node := Node[T]{
		data: value,
		next: nil,
		prev: ll.end,
	}
	ll.end.next = &node
	ll.end = &node
}

func (ll *LinkedList[T]) AppendHead(value T) {
	node := Node[T]{
		data: value,
		next: ll.head,
		prev: nil,
	}
	ll.head.prev = &node
	ll.head = &node
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
func (ll *LinkedList[T]) At(ix int) *Node[T] {
	pos := 0
	var ptr *Node[T]
	if ix >= 0 {
		ptr = ll.head
		for pos != ix && ptr != nil {
			ptr = ptr.prev
		}
	}
	if ix < 0 {
		ix = -ix
		ptr = ll.end
		for pos != ix && ptr != nil {
			ptr = ptr.next
		}
	}
	return ptr
}
