package linkedlist

// A node in a doubly linked list.
type DoublyNode[T any] struct {
	value T
	next  *DoublyNode[T]
	prev  *DoublyNode[T]
}

// Returns the node's next pointer.
func (n *DoublyNode[T]) Next() *DoublyNode[T] {
	return n.next
}

// Returns the node's previous pointer.
func (n *DoublyNode[T]) Prev() *DoublyNode[T] {
	return n.prev
}

// Returns the node's value.
func (n *DoublyNode[T]) Value() T {
	return n.value
}
