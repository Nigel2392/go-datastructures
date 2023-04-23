package hashmap

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Nigel2392/go-datastructures"
)

const defaultBucketLen = 16

// Not to be used directly. Use the Map() function instead.
//
// A hashmap implementation with a specified number of buckets.
//
// These buckets are implemented as binary search trees.
//
// Inside the binary search trees, the keys are stored as linked lists.
//
// It is up to the user to ensure that the key type implements the datastructures.Hashable[T] interface,
//
// and that the key hashing function is secure, fast and collision-free.
type HashMap[T1 datastructures.Hashable[T1], T2 any] struct {
	buckets   []*bucket[T1, T2]
	len       int
	bucketLen uint64
}

// Returns a new HashMap[T1, T2].
//
// If no argument is given, the default number of buckets is used (16).
func Map[T1 datastructures.Hashable[T1], T2 any](amount ...int) *HashMap[T1, T2] {
	if len(amount) == 0 {
		return newMap[T1, T2](defaultBucketLen)
	} else if len(amount) > 1 {
		panic(fmt.Sprintf("Map[T1, T2] takes at most 1 argument, %d given", len(amount)))
	}
	var k = amount[0]
	if k < 0 {
		panic(fmt.Sprintf("Map[T1, T2] takes a positive integer, %d given", amount))
	}
	if k == 0 {
		return newMap[T1, T2](defaultBucketLen)
	}

	return newMap[T1, T2](calcBuckets(uint64(k)))
}

// Calculates the number of buckets to use for the given number of items.
//
// The number of buckets is a power of 2.
func calcBuckets(items uint64) uint64 {
	if items > (1 << 31) {
		return uint64(float64(items) * 1.5)
	}

	var buckets uint64 = 1
	for buckets < items {
		buckets <<= 1
	}

	if buckets > (1 << 31) {
		return uint64(float64(buckets>>1) * 1.5)
	}

	return buckets
}

// Instantiates a new HashMap[T1, T2] with the given number of buckets.
func newMap[T1 datastructures.Hashable[T1], T2 any](buckets uint64) *HashMap[T1, T2] {
	var table = HashMap[T1, T2]{
		bucketLen: buckets,
		buckets:   make([]*bucket[T1, T2], buckets, buckets),
	}
	for i := range table.buckets {
		table.buckets[i] = &bucket[T1, T2]{}
	}
	return &table
}

func indexOf(hash uint64, buckets uint64) uint64 {
	return (hash ^ (hash >> 16)) & (buckets - 1)
}

func (t *HashMap[T1, T2]) Set(k T1, v T2) {
	var hash uint64 = k.Hash()
	t.buckets[indexOf(hash, t.bucketLen)].insert(hash, k, v)
	t.len++
}

func (t *HashMap[T1, T2]) Get(k T1) (v T2, ok bool) {
	var hash uint64 = k.Hash()
	return t.buckets[indexOf(hash, t.bucketLen)].retrieve(k)
}

func (t *HashMap[T1, T2]) Delete(k T1) (ok bool) {
	var hash uint64 = k.Hash()
	ok = t.buckets[indexOf(hash, t.bucketLen)].delete(k)
	if ok {
		t.len--
	}
	return
}

func (t *HashMap[T1, T2]) Len() int {
	return t.len
}

func (t *HashMap[T1, T2]) Keys() []T1 {
	var keys = make([]T1, t.len, t.len)
	var i int
	for _, bucket := range t.buckets {
		bucket.traverse(func(k T1, v T2) bool {
			keys[i] = k
			i++
			return true
		})
	}

	return keys
}

func (t *HashMap[T1, T2]) Values() []T2 {
	var values = make([]T2, t.len, t.len)
	var i int
	for _, bucket := range t.buckets {
		bucket.traverse(func(k T1, v T2) bool {
			values[i] = v
			i++
			return true
		})
	}
	return values
}

func (t *HashMap[T1, T2]) Clear() {
	for i := range t.buckets {
		t.buckets[i] = &bucket[T1, T2]{}
	}
	t.len = 0
}

func (t *HashMap[T1, T2]) Range(f func(k T1, v T2) (continueLoop bool)) {
	for _, bucket := range t.buckets {
		if !bucket.traverse(f) {
			return
		}
	}
}

func (t *HashMap[T1, T2]) Pop(k T1) (v T2, ok bool) {
	var hash uint64 = k.Hash()
	v, ok = t.buckets[indexOf(hash, t.bucketLen)].pop(k)
	if ok {
		t.len--
	}
	return
}

func (t *HashMap[T1, T2]) String() string {
	var b strings.Builder
	b.WriteString("{")
	var i int
	t.Range(func(k T1, v T2) (continueLoop bool) {
		b.WriteString(fmt.Sprintf("%v:%v\n", k, v))
		if i < t.len-1 {
			b.WriteString(", ")
		}
		i++
		return true
	})
	b.WriteString("}")
	return b.String()
}

func (t *HashMap[T1, T2]) GoString() string {
	var b strings.Builder
	b.WriteString("Map[T1, T2]{")
	for j, bucket := range t.buckets {
		var i int
		b.WriteString("\n\tBucket{")
		b.WriteString(fmt.Sprintf("index: %d", j))
		b.WriteString(", bucketLen: ")
		b.WriteString(strconv.FormatInt(int64(bucket._len), 10))
		b.WriteString(", items: ")
		b.WriteString("[")
		bucket.traverse(func(k T1, v T2) bool {
			b.WriteString(fmt.Sprintf("%v:%v", k, v))
			if i < bucket._len-1 {
				b.WriteString(", ")
			}
			i++
			return true
		})
		b.WriteString("]")
		b.WriteString("}")
		if j < len(t.buckets)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("\n}")
	return b.String()
}
