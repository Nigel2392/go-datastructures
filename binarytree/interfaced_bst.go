package binarytree

import (
	"fmt"
	"math"
	"strings"

	"github.com/Nigel2392/go-datastructures"
)

type InterfacedBST[T datastructures.Comparable[T]] struct {
	root   *IF_BSTNode[T]
	len    int
	height int
}

func (t *InterfacedBST[T]) String() string {
	if t.root == nil {
		return ""
	}

	height := t.root.getHeight()
	IF_BSTNodes := make([][]string, height)

	fillIF_BSTNodes(IF_BSTNodes, t.root, 0)

	var b strings.Builder
	padding := int(math.Pow(2, float64(height)) - 1)

	for i, level := range IF_BSTNodes {
		if i == 0 {
			paddingStr := strings.Repeat(" ", (padding/2)+1)
			b.WriteString(paddingStr)
		} else {
			paddingStr := strings.Repeat(" ", padding/2)
			b.WriteString(paddingStr)
		}

		for j, IF_BSTNode := range level {
			b.WriteString(IF_BSTNode)
			if j != len(level)-1 {
				b.WriteString(strings.Repeat(" ", padding))
			}
		}

		padding /= 2
		b.WriteString("\n")
	}

	return b.String()
}

func NewInterfaced[T datastructures.Comparable[T]](initial T) *InterfacedBST[T] {
	return &InterfacedBST[T]{
		root: &IF_BSTNode[T]{value: initial}}
}

func (t *InterfacedBST[T]) Insert(value T) (inserted bool) {
	if t.root == nil {
		t.root = &IF_BSTNode[T]{value: value}
		t.len++
		return true
	}
	inserted = t.root.insert(value)
	if inserted {
		t.len++
	}
	return inserted
}

func (t *InterfacedBST[T]) Search(value T) (v T, ok bool) {
	if t.root == nil {
		return
	}
	return t.root.search(value)
}

func (t *InterfacedBST[T]) Delete(value T) (deleted bool) {
	if t.root == nil {
		return false
	}
	t.root, deleted = t.root.delete(value)
	if deleted {
		t.len--
	}
	return deleted
}

func (t *InterfacedBST[T]) DeleteIf(predicate func(T) bool) (deleted int) {
	if t.root == nil {
		return 0
	}
	t.root, deleted = t.root.deleteIf(predicate)
	t.len -= int(deleted)
	return deleted
}

func (t *InterfacedBST[T]) Traverse(f func(T)) {
	if t.root == nil {
		return
	}
	t.root.traverse(f)
}

func (t *InterfacedBST[T]) Len() int {
	return t.len
}

func (t *InterfacedBST[T]) Height() int {
	if t.root == nil {
		return 0
	}
	return t.root.getHeight()
}

func (t *InterfacedBST[T]) Clear() {
	t.root = nil
	t.len = 0
}

func fillIF_BSTNodes[T datastructures.Comparable[T]](IF_BSTNodes [][]string, n *IF_BSTNode[T], depth int) {
	if n == nil {
		return
	}

	IF_BSTNodes[depth] = append(IF_BSTNodes[depth], fmt.Sprintf("%v", n.value))
	fillIF_BSTNodes(IF_BSTNodes, n.left, depth+1)
	fillIF_BSTNodes(IF_BSTNodes, n.right, depth+1)
}
