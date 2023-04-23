package stack_test

import (
	"testing"

	"github.com/Nigel2392/go-datastructures/stack"
)

func TestStack(t *testing.T) {
	var s = stack.Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)
	s.Push(9)
	s.Push(10)

	if s.Pop() != 10 {
		t.Errorf("Expected 10, got %d", s.Pop())
	}
	if s.Pop() != 9 {
		t.Errorf("Expected 9, got %d", s.Pop())
	}
	if s.Pop() != 8 {
		t.Errorf("Expected 8, got %d", s.Pop())
	}
	if s.Pop() != 7 {
		t.Errorf("Expected 7, got %d", s.Pop())
	}
	if s.Pop() != 6 {
		t.Errorf("Expected 6, got %d", s.Pop())
	}
	if s.Pop() != 5 {
		t.Errorf("Expected 5, got %d", s.Pop())
	}
	if s.Pop() != 4 {
		t.Errorf("Expected 4, got %d", s.Pop())
	}
	if s.Pop() != 3 {
		t.Errorf("Expected 3, got %d", s.Pop())
	}
	if s.Pop() != 2 {
		t.Errorf("Expected 2, got %d", s.Pop())
	}
	if s.Pop() != 1 {
		t.Errorf("Expected 1, got %d", s.Pop())
	}
}
