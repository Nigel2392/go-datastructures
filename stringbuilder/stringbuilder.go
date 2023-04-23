package stringbuilder

import (
	"io"
	"strings"

	"github.com/Nigel2392/go-datastructures/linkedlist"
)

// A string builder implemented with a linked list.
//
// This allows for O(1) prepend and append operations.
//
// This is a wrapper around a linkedlist.Doubly[string].
type StringBuilder struct {
	linkedlist.Doubly[string]
	len int
}

// Len returns the length of the string builder.
func (sb *StringBuilder) Len() int {
	return sb.len
}

// Append a string to the end of the string builder.
func (sb *StringBuilder) Append(s string) {
	sb.len += len(s)
	sb.Doubly.Append(s)
}

// Prepend a string to the beginning of the string builder.
func (sb *StringBuilder) Prepend(s string) {
	sb.len += len(s)
	sb.Doubly.Prepend(s)
}

// Reset the string builder.
//
// This will clear the string builders internal nodes.
func (sb *StringBuilder) Reset() {
	sb.len = 0
	sb.Doubly.Reset()
}

// Returns the string builder as a string.
func (sb *StringBuilder) String() string {
	var (
		i int
		n *linkedlist.DoublyNode[string]
	)
	var rs = make([]byte, sb.len)
	for n = sb.Doubly.Head(); n != nil; n = n.Next() {
		i += copy(rs[i:], n.Value())
	}
	return string(rs)
}

// Checks if the stringbuilder contains the given string.
//
// It will allocate a new string to check if the stringbuilder contains the given string.
func (sb *StringBuilder) Contains(s string) bool {
	return strings.Contains(sb.String(), s)
}

// Functionality for the Stringer interface.
func (sb *StringBuilder) GoString() string {
	return sb.String()
}

// Implement io.Writer.
//
// Appends to the end of the stringbuilder.
func (sb *StringBuilder) Write(p []byte) (n int, err error) {
	sb.Append(string(p))
	return len(p), nil
}

// Implement io.ByteWriter.
//
// Appends to the end of the stringbuilder.
func (sb *StringBuilder) WriteByte(c byte) error {
	sb.Append(string(c))
	return nil
}

// Implement io.RuneWriter.
//
// Appends to the end of the stringbuilder.
func (sb *StringBuilder) WriteRune(r rune) (n int, err error) {
	sb.Append(string(r))
	return 1, nil
}

// Implement io.StringWriter.
//
// Appends to the end of the stringbuilder.
func (sb *StringBuilder) WriteString(s string) (n int, err error) {
	sb.Append(s)
	return len(s), nil
}

type Grower interface {
	Grow(n int)
}

// Implement io.WriterTo.
//
// Writes the stringbuilder to the io.Writer.
func (sb *StringBuilder) WriteTo(w io.Writer) (n int64, err error) {
	var (
		node *linkedlist.DoublyNode[string]
	)

	if g, ok := w.(Grower); ok {
		g.Grow(sb.len)
	}

	for node = sb.Doubly.Head(); node != nil; node = node.Next() {
		var nn int
		nn, err = w.Write([]byte(node.Value()))
		n += int64(nn)
		if err != nil {
			return
		}
	}
	return
}
