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

// PopOK removes a value from the top of the stack
//
// This is the same as Shift in a linked list
//
// It returns the value and a boolean indicating whether the value was removed
func (s *Stack[T]) PopOK() (value T, ok bool) {
	if s.Len() == 0 {
		return
	}
	return (*linkedlist.Singly[T])(s).Shift(), true
}

// Peek returns the value at the top of the stack
//
// This is the same as Head in a linked list
func (s *Stack[T]) Peek() T {
	return (*linkedlist.Singly[T])(s).Head().Value()
}

func (s *Stack[T]) Len() int {
	return (*linkedlist.Singly[T])(s).Len()
}
