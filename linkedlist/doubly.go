package linkedlist

import (
	"encoding/json"
	"fmt"
	"strings"
)

// A doubly linked list.
//
// This is a generic type, so you can use it with any type.
//
// For example, Doubly[int] is a linked list of integers.
//
// If using an uncomparable type, the Remove(value) method will not work.
//
// You can only use RemoveIndex(index).
type Doubly[T any] struct {
	head *DoublyNode[T]
	tail *DoublyNode[T]
	len  int
}

// Returns the length of the list.
func (l *Doubly[T]) Len() int {
	return l.len
}

// Returns the list as a string.
func (l *Doubly[T]) String() string {
	if l.len == 0 {
		return "[]"
	}
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
func (l *Doubly[T]) Head() *DoublyNode[T] {
	return l.head
}

// Returns the tail of the list.
func (l *Doubly[T]) Tail() *DoublyNode[T] {
	return l.tail
}

// Append a value to the end of the list.
func (l *Doubly[T]) Append(v T) {
	l.append(&DoublyNode[T]{value: v})
}

// Prepend a value to the beginning of the list.
func (l *Doubly[T]) Prepend(v T) {
	l.prepend(&DoublyNode[T]{value: v})
}

// Pop a value from the end of the list.
//
// Returns the value that was popped.
//
// Returns the zero value of the type if the list is empty.
func (l *Doubly[T]) Pop() T {
	if l.len == 0 {
		panic("cannot Pop() from an empty list")
	}
	var v = l.tail.value
	l.RemoveIndex(l.len - 1)
	return v
}

// Shift a value from the beginning of the list.
//
// Returns the value that was shifted.
//
// Returns the zero value of the type if the list is empty.
func (l *Doubly[T]) Shift() T {
	if l.len == 0 {
		panic("cannot Shift() from an empty list")
	}
	var v = l.head.value
	l.RemoveIndex(0)
	return v
}

// Reset the list.
func (l *Doubly[T]) Reset() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

// Remove a value from the list at a given index.
//
// This is O(n/2) time complexity.
func (l *Doubly[T]) RemoveIndex(i int) bool {
	if i < 0 || i >= l.len {
		return false
	}

	if i == 0 {
		l.remove(i, l.head)
		return true
	}

	if i == l.len-1 {
		l.remove(i, l.tail)
		return true
	}

	var n *DoublyNode[T]
	if i < l.len/2 {
		n = l.head
		for j := 0; j < i; j++ {
			n = n.next
		}
	} else {
		n = l.tail
		for j := l.len - 1; j > i; j-- {
			n = n.prev
		}
	}
	if n != nil {
		l.remove(i, n)
		return true
	}
	return false
}

// Returns the list as a slice.
func (l *Doubly[T]) ToSlice() []T {
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
func (l *Doubly[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.ToSlice())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *Doubly[T]) UnmarshalJSON(data []byte) error {
	var slice []T
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}
	for _, v := range slice {
		l.Append(v)
	}
	return nil
}

func (l *Doubly[T]) append(n *DoublyNode[T]) {
	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		var oldTail = l.tail
		l.tail = n
		l.tail.prev = oldTail
		oldTail.next = l.tail
	}
	l.len++
}

func (l *Doubly[T]) prepend(n *DoublyNode[T]) {
	if l.head == nil && l.tail == nil {
		l.head = n
		l.tail = n
	} else {
		var oldHead = l.head
		l.head = n
		l.head.next = oldHead
		oldHead.prev = l.head
	}
	l.len++
}

// remove removes a node from the list.
func (l *Doubly[T]) remove(i int, n *DoublyNode[T]) {
	if i == 0 {
		l.head = n.next
	}
	if i == l.len-1 {
		l.tail = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	l.len--
}
