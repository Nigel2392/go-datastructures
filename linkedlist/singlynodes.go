package linkedlist

// A node in a singly linked list.
type Node[T any] struct {
	value T
	next  *Node[T]
}

// Returns the node's next pointer.
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Returns the node's value.
func (n *Node[T]) Value() T {
	return n.value
}
