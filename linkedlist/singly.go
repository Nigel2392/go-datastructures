package linkedlist

import (
	"encoding/json"
	"fmt"
	"strings"
)

// A singly linked list.
//
// This is a generic type, so you can use it with any type.
//
// For example, Singly[int] is a linked list of integers.
//
// If using an uncomparable type, the Remove(value) method will not work.
//
// You can only use RemoveIndex(index).
type Singly[T any] struct {
	head *Node[T]
	len  int
}

// Returns the length of the list.
func (l *Singly[T]) Len() int {
	return l.len
}

// Returns the list as a string.
func (l *Singly[T]) String() string {
	var b strings.Builder
	b.WriteString("[")
	for n := l.head; n != nil; n = n.next {
		fmt.Fprintf(&b, "%v", n.value)
		if n.next != nil {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}

// Returns the head of the list.
func (l *Singly[T]) Head() *Node[T] {
	return l.head
}

// Prepend a value to the beginning of the list.
func (l *Singly[T]) Prepend(v T) {
	var n = &Node[T]{value: v}
	l.prepend(n)
}

// Shift a value from the beginning of the list.
//
// Returns the value that was shifted.
func (l *Singly[T]) Shift() T {
	if l.len == 0 {
		panic("cannot shift from an empty list")
	}
	var v = l.head.value
	l.RemoveIndex(0)
	return v
}

// Reset the list.
func (l *Singly[T]) Reset() {
	l.head = nil
	l.len = 0
}

// Remove a value from the list.
//
// This function will panic with uncomparable types.
//
// Use RemoveIndex(index) instead.
func (l *Singly[T]) Remove(v T) bool {
	if l.len == 0 {
		return false
	}
	var i = 0
	for n := l.head; n != nil; n = n.next {
		if any(n.value) == any(v) { // panic on comparison of uncomparable types.
			l.remove(i, n)
			return true
		}
		i++
	}
	return false
}

// Remove a value from the list at a given index.
func (l *Singly[T]) RemoveIndex(i int) bool {
	if i < 0 || i >= l.len {
		return false
	}

	if i == 0 {
		l.remove(i, l.head)
		return true
	}

	var n *Node[T]
	n = l.head
	for j := 0; j < i; j++ {
		n = n.next
	}
	if n != nil {
		l.remove(i, n)
		return true
	}
	return false
}

// Returns the list as a slice.
func (l *Singly[T]) ToSlice() []T {
	if l.len == 0 {
		return nil
	}
	var slice = make([]T, l.len)
	var i = 0
	for n := l.head; n != nil; n = n.next {
		slice[i] = n.value
		i++
	}
	return slice
}

// MarshalJSON implements the json.Marshaler interface.
func (l *Singly[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToSlice())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *Singly[T]) UnmarshalJSON(data []byte) error {
	var slice []T
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}
	for _, v := range slice {
		l.Prepend(v)
	}
	return nil
}

func (l *Singly[T]) prepend(n *Node[T]) {
	if l.head == nil {
		l.head = n
	} else {
		var oldHead = l.head
		l.head = n
		l.head.next = oldHead
	}
	l.len++
}

// remove removes a node from the list.
func (l *Singly[T]) remove(i int, n *Node[T]) {
	if i == 0 {
		l.head = n.next
	} else {
		if n.next != nil {
			*n = *n.next
		}
	}
	l.len--

}
