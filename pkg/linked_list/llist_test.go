package linkedlist_test

import (
	"slices"
	"testing"

	list "github.com/eonianmonk/snake-game/pkg/linked_list"
)

func TestLLis(t *testing.T) {
	t.Run("test-llist", func(t *testing.T) {
		entries := []int{0, 1, 2, 3, 4}
		ll := list.LinkedListFromSlice(entries)
		// indexing
		for i := 0; i < len(entries); i++ {
			if ll.At(i).Value() != entries[i] {
				t.Fatal("failed at llist indexing")
			}
		}
		// inverse indexing
		for i := len(entries); i > 0; i-- {
			if ll.At(-i).Value() != entries[len(entries)-i] {
				t.Fatalf("failed at inverse llist indexing: %d, expected %d", ll.At(-i).Value(), entries[len(entries)-i])
			}
		}
		//delete 0
		ll.DeleteAt(0)
		sample := ll.Slice()
		if slices.Compare(sample, entries[1:]) != 0 {
			t.Fatalf("failed to delete at 0: got %v, wanted %v", sample, entries[1:])
		}
		// delete last
		ll = list.LinkedListFromSlice(entries)
		ll.DeleteAt(-1)
		sample = ll.Slice()
		if slices.Compare(sample, entries[:len(entries)-1]) != 0 {
			t.Fatalf("failed to delete at -1: got %v, wanted %v", sample, entries[:len(entries)-1])
		}
	})
	t.Run("llist-iter", func(t *testing.T) {
		entries := []int{0, 1, 2, 3, 4}
		ll := list.LinkedListFromSlice(entries)
		i := 0
		it := list.NewLLIter(ll)
		for i < len(entries) && !it.Done() {
			val := it.Next()
			if entries[i] != val {
				t.Fatalf("failed to iterate though values: got %d, expected %d", val, entries[i])
			}
			i++
		}
		if i != len(entries) {
			t.Fatalf("failed to iterate through all values")
		}
	})
}
