package hashmap

import (
	"fmt"

	"github.com/Nigel2392/go-datastructures"
)

type keyNode[T1 datastructures.Hashable[T1], T2 any] struct {
	_hash uint64
	key   T1
	value T2
	next  *keyNode[T1, T2]
}

func (n *keyNode[T1, T2]) insert(v *keyNode[T1, T2]) {
	if n.key.Equals(v.key) {
		n.value = v.value
		return
	}

	if n.next == nil {
		n.next = v
		return
	}

	n.next.insert(v)
}

func (n *keyNode[T1, T2]) retrieve(other *keyNode[T1, T2]) (value T2, ok bool) {
	if n == nil {
		return
	}

	if n.key.Equals(other.key) {
		return n.value, true
	}

	return n.next.retrieve(other)
}

func (n *keyNode[T1, T2]) delete(other *keyNode[T1, T2]) (newRoot *keyNode[T1, T2], deleted bool) {
	if n == nil {
		return nil, false
	}

	if n.key.Equals(other.key) {
		return n.next, true
	}

	n.next, deleted = n.next.delete(other)
	return n, deleted
}

func (n *keyNode[T1, T2]) pop(other *keyNode[T1, T2]) (newRoot *keyNode[T1, T2], value T2, ok bool) {
	if n == nil {
		return
	}

	if n.key.Equals(other.key) {
		return n.next, n.value, true
	}

	n.next, value, ok = n.next.pop(other)
	return n, value, ok
}

type bucketNode[T1 datastructures.Hashable[T1], T2 any] struct {
	_hash uint64
	next  *keyNode[T1, T2]
	left  *bucketNode[T1, T2]
	right *bucketNode[T1, T2]
}

func (n *bucketNode[T1, T2]) insert(v *keyNode[T1, T2]) {
	if n._hash == v._hash {
		if n.next == nil {
			n.next = v
			return
		}
		n.next.insert(v)
		return
	} else if n._hash < v._hash {
		if n.right == nil {
			n.right = &bucketNode[T1, T2]{
				_hash: v._hash,
				next:  v,
			}
		} else {
			n.right.insert(v)
		}
		return
	} else if n._hash > v._hash {
		if n.left == nil {
			n.left = &bucketNode[T1, T2]{
				_hash: v._hash,
				next:  v,
			}
		} else {
			n.left.insert(v)
		}
		return
	}
}

func (n *bucketNode[T1, T2]) retrieve(k *keyNode[T1, T2]) (value T2, ok bool) {
	if n == nil {
		return
	}

	if n._hash < k._hash {
		return n.right.retrieve(k)
	} else if n._hash > k._hash {
		return n.left.retrieve(k)
	}

	return n.next.retrieve(k)
}

func (n *bucketNode[T1, T2]) delete(other *keyNode[T1, T2]) (newRoot *bucketNode[T1, T2], deleted bool) {
	if n == nil {
		return nil, false
	}

	if other._hash < n._hash {
		n.left, deleted = n.left.delete(other)
	} else if other._hash > n._hash {
		n.right, deleted = n.right.delete(other)
	} else {
		n.next, deleted = n.next.delete(other)
	}
	if n.next == nil {
		fmt.Println("found, deleting node")
		if n.left == nil {
			return n.right, true
		} else if n.right == nil {
			return n.left, true
		}

		// find the min node in the right subtree
		var minNode = n.right.findMin()
		n._hash = minNode._hash
		n.next = minNode.next
		// delete the min node from the right subtree
		n.right, deleted = n.right.deleteNode(minNode)
	}
	return n, deleted
}

// deleteNode deletes the node with the given hash
func (n *bucketNode[T1, T2]) deleteNode(other *bucketNode[T1, T2]) (newRoot *bucketNode[T1, T2], deleted bool) {
	if n == nil {
		return nil, false
	}

	if other._hash < n._hash {
		n.left, deleted = n.left.deleteNode(other)
	} else if other._hash > n._hash {
		n.right, deleted = n.right.deleteNode(other)
	} else {
		if n.left == nil {
			return n.right, true
		} else if n.right == nil {
			return n.left, true
		}

		// find the min node in the right subtree
		var minNode = n.right.findMin()
		n._hash = minNode._hash
		n.next = minNode.next
		// delete the min node from the right subtree
		n.right, deleted = n.right.deleteNode(minNode)
	}
	return n, deleted
}

func (n *bucketNode[T1, T2]) pop(k *keyNode[T1, T2]) (newRoot *bucketNode[T1, T2], value T2, ok bool) {
	if n == nil {
		return nil, value, false
	}

	if k._hash < n._hash {
		n.left, value, ok = n.left.pop(k)
	} else if k._hash > n._hash {
		n.right, value, ok = n.right.pop(k)
	} else {
		n.next, value, ok = n.next.pop(k)
		if n.next == nil {
			if n.left == nil {
				return n.right, value, ok
			} else if n.right == nil {
				return n.left, value, ok
			}

			// find the min node in the right subtree
			var minNode = n.right.findMin()
			n._hash = minNode._hash
			n.next = minNode.next
			// delete the min node from the right subtree
			n.right, _ = n.right.deleteNode(minNode)
		}
	}

	return n, value, ok
}

func (n *bucketNode[T1, T2]) findMin() *bucketNode[T1, T2] {
	if n.left == nil {
		return n
	}

	return n.left.findMin()
}

func (n *bucketNode[T1, T2]) traverse(f func(k T1, v T2) bool) (continueLoop bool) {
	if n == nil {
		return true
	}
	if !n.left.traverse(f) {
		return true
	}
	for n.next != nil {
		if !f(n.next.key, n.next.value) {
			return false
		}
		n.next = n.next.next
	}
	if !n.right.traverse(f) {
		return true
	}
	return true
}
