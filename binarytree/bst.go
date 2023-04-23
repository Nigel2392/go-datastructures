package binarytree

import (
	"fmt"
	"math"
	"strings"

	"github.com/Nigel2392/go-datastructures"
)

type BST[T datastructures.Ordered] struct {
	root   *BSTNode[T]
	len    int
	height int
}

func (t *BST[T]) String() string {
	if t.root == nil {
		return ""
	}

	height := t.root.getHeight()
	BSTNodes := make([][]string, height)

	fillBSTNodes(BSTNodes, t.root, 0)

	var b strings.Builder
	padding := int(math.Pow(2, float64(height)) - 1)

	for i, level := range BSTNodes {
		if i == 0 {
			paddingStr := strings.Repeat(" ", (padding/2)+1)
			b.WriteString(paddingStr)
		} else {
			paddingStr := strings.Repeat(" ", padding/2)
			b.WriteString(paddingStr)
		}

		for j, BSTNode := range level {
			b.WriteString(BSTNode)
			if j != len(level)-1 {
				b.WriteString(strings.Repeat(" ", padding))
			}
		}

		padding /= 2
		b.WriteString("\n")
	}

	return b.String()
}

func NewBST[T datastructures.Ordered](initial T) *BST[T] {
	return &BST[T]{
		root: &BSTNode[T]{value: initial}}
}

func (t *BST[T]) Insert(value T) (inserted bool) {
	if t.root == nil {
		t.root = &BSTNode[T]{value: value}
		t.len++
		return true
	}
	inserted = t.root.insert(value)
	if inserted {
		t.len++
	}
	return inserted
}

func (t *BST[T]) Search(value T) (v T, ok bool) {
	if t.root == nil {
		return
	}
	return t.root.search(value)
}

func (t *BST[T]) Delete(value T) (deleted bool) {
	if t.root == nil {
		return false
	}
	t.root, deleted = t.root.delete(value)
	if deleted {
		t.len--
	}
	return deleted
}

func (t *BST[T]) DeleteIf(predicate func(T) bool) (deleted int) {
	if t.root == nil {
		return 0
	}
	t.root, deleted = t.root.deleteIf(predicate)
	t.len -= int(deleted)
	return deleted
}

func (t *BST[T]) Traverse(f func(T)) {
	if t.root == nil {
		return
	}
	t.root.traverse(f)
}
func (t *BST[T]) Len() int {
	return t.len
}

func (t *BST[T]) Height() int {
	if t.root == nil {
		return 0
	}
	return t.root.getHeight()
}

func fillBSTNodes[T datastructures.Ordered](BSTNodes [][]string, n *BSTNode[T], depth int) {
	if n == nil {
		return
	}

	BSTNodes[depth] = append(BSTNodes[depth], fmt.Sprintf("%v", n.value))
	fillBSTNodes(BSTNodes, n.left, depth+1)
	fillBSTNodes(BSTNodes, n.right, depth+1)
}
