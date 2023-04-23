package hashmap

import "github.com/Nigel2392/go-datastructures"

type bucket[T1 datastructures.Hashable[T1], T2 any] struct {
	root *bucketNode[T1, T2]
	_len int
}

func (b *bucket[T1, T2]) insert(hash uint64, k T1, v T2) {

	var key = &keyNode[T1, T2]{
		_hash: hash,
		key:   k,
		value: v,
	}

	if b.root == nil {
		var newNode = &bucketNode[T1, T2]{
			_hash: hash,
			next:  key,
		}

		b._len++
		b.root = newNode
		return
	}

	b.root.insert(key)
	b._len++
}

func (b *bucket[T1, T2]) retrieve(k T1) (v T2, ok bool) {
	var key = &keyNode[T1, T2]{
		_hash: k.Hash(),
		key:   k,
	}

	return b.root.retrieve(key)
}

func (b *bucket[T1, T2]) delete(k T1) (ok bool) {
	if b.root == nil {
		return false
	}

	var key = &keyNode[T1, T2]{
		_hash: k.Hash(),
		key:   k,
	}

	b.root, ok = b.root.delete(key)
	if ok {
		b._len--
	}
	return ok
}

func (b *bucket[T1, T2]) pop(k T1) (v T2, ok bool) {
	if b.root == nil {
		return
	}

	var key = &keyNode[T1, T2]{
		_hash: k.Hash(),
		key:   k,
	}

	b.root, v, ok = b.root.pop(key)
	if ok {
		b._len--
	}
	return
}

func (b *bucket[T1, T2]) traverse(f func(k T1, v T2) bool) (continueLoop bool) {
	return b.root.traverse(f)
}

func (b *bucket[T1, T2]) len() int {
	return b._len
}
