package linkedlist

type lliter[T any] struct {
	ptr *Node[T]
}

type Iter[T any] interface {
	Done() bool
	Next() T
}

func NewLLIter[T any](ll *LinkedList[T]) Iter[T] {
	return &lliter[T]{ptr: ll.head}
}

func (i *lliter[T]) Done() bool {
	return i.ptr == nil
}

func (i *lliter[T]) Next() T {
	val := i.ptr.Value()
	i.ptr = i.ptr.next
	return val
}
