package linkedlist_test

import (
	"testing"

	"github.com/Nigel2392/go-datastructures/linkedlist"
)

func TestSingly(t *testing.T) {
	var s = new(linkedlist.Singly[string])
	var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, l := range letters {
		s.Prepend(l)
	}
	var i = 0
	for n := s.Head(); n != nil; n = n.Next() {
		if n.Value() != letters[len(letters)-1-i] {
			t.Errorf("Expected %s, got %s", letters[len(letters)-1-i], n.Value())
		}
		i++
	}

	for i := 0; i < 13; i++ {
		var letter = s.Shift()
		if letter != letters[len(letters)-1-i] {
			t.Errorf("Expected %s, got %s", letters[len(letters)-1-i], letter)
		}
	}

	if s.Len() != len(letters)/2 {
		t.Errorf("Expected length %d, got %d", len(letters)/2, s.Len())
	}

	t.Log(s)
}

func TestDoubly(t *testing.T) {
	var d = new(linkedlist.Doubly[string])
	var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, l := range letters {
		d.Append(l)
	}
	var i = 0
	for n := d.Head(); n != nil; n = n.Next() {
		if n.Value() != letters[i] {
			t.Errorf("Expected %s, got %s", letters[i], n.Value())
		}
		i++
	}

	for i := 0; i < 13; i++ {
		var letter = d.Pop()
		if letter != letters[len(letters)-1-i] {
			t.Errorf("Expected %s, got %s", letters[len(letters)-1-i], letter)
		}
	}

	if d.Len() != len(letters)/2 {
		t.Errorf("Expected length %d, got %d", len(letters)/2, d.Len())
	}

	t.Log(d)
}
