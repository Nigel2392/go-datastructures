package stack

import "github.com/Nigel2392/go-datastructures/linkedlist"

// Stack is a stack data structure
//
// It is implemeneted using a singly linked list
//
// This is to save space, as the stack only needs to know about the top element
//
// This does mean that the data gets prepended at the start, and then shifted off,
// as opposed to appended and then popped off.
type Stack[T any] linkedlist.Singly[T]

// Push adds a value to the top of the stack
//
// This is the same as Prepend in a linked list
func (s *Stack[T]) Push(value T) {
	(*linkedlist.Singly[T])(s).Prepend(value)
}

// Pop removes a value from the top of the stack
//
// This is the same as Shift in a linked list
func (s *Stack[T]) Pop() T {
	return (*linkedlist.Singly[T])(s).Shift()
}
