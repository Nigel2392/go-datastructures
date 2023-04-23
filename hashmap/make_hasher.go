package hashmap

import (
	"fmt"
	"unsafe"

	"github.com/Nigel2392/go-datastructures"
)

type comparable interface {
	~int | ~uint | ~string | ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~complex64 | ~complex128 | ~bool
}

func Key[T comparable](v T) *MapKey[T] {
	return makeHasher(v, getHashFunc[T]())
}

type MapKey[T comparable] struct {
	_hash func(T) uint64
	v     T
}

func (h *MapKey[T]) Hash() uint64 {
	return h._hash(h.v)
}

func (h *MapKey[T]) Equals(other *MapKey[T]) bool {
	return h.v == other.v
}

func (h *MapKey[T]) Value() T {
	return h.v
}

func (h *MapKey[T]) String() string {
	return fmt.Sprintf("%v", h.v)
}

func (h *MapKey[T]) GoString() string {
	return fmt.Sprintf("%#v", h.v)
}

func makeHasher[T comparable](v T, hashfunc func(T) uint64) *MapKey[T] {
	switch any(*new(T)).(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128:
		return &MapKey[T]{_hash: hashfunc, v: v}
	case string:
		return &MapKey[T]{_hash: hashfunc, v: v}
	case bool:
		return &MapKey[T]{_hash: hashfunc, v: v}
	case nil:
		return &MapKey[T]{_hash: hashfunc, v: v}
	}
	panic("type is not comparable!")
}

func getHashFunc[T comparable]() func(T) uint64 {
	return genericHashFunc[T]
}

func genericHashFunc[T comparable](v T) uint64 {
	switch any(*new(T)).(type) {
	case string:
		return _hash_string(*(*string)(unsafe.Pointer(&v)))
	case int:
		return _hash_int(*(*int)(unsafe.Pointer(&v)))
	case int8:
		return _hash_int(*(*int8)(unsafe.Pointer(&v)))
	case int16:
		return _hash_int(*(*int16)(unsafe.Pointer(&v)))
	case int32:
		return _hash_int(*(*int32)(unsafe.Pointer(&v)))
	case int64:
		return _hash_int(*(*int64)(unsafe.Pointer(&v)))
	case uint:
		return _hash_int(*(*uint)(unsafe.Pointer(&v)))
	case uint8:
		return _hash_int(*(*uint8)(unsafe.Pointer(&v)))
	case uint16:
		return _hash_int(*(*uint16)(unsafe.Pointer(&v)))
	case uint32:
		return _hash_int(*(*uint32)(unsafe.Pointer(&v)))
	case uint64:
		return _hash_int(*(*uint64)(unsafe.Pointer(&v)))
	case uintptr:
		return _hash_int(*(*uintptr)(unsafe.Pointer(&v)))
	case float32:
		return _hash_float(*(*float32)(unsafe.Pointer(&v)))
	case float64:
		return _hash_float(*(*float64)(unsafe.Pointer(&v)))
	case complex64:
		return _hash_float(*(*float32)(unsafe.Pointer(&v)))
	case complex128:
		return _hash_float(*(*float64)(unsafe.Pointer(&v)))
	case bool:
		return _bool_hash(*(*bool)(unsafe.Pointer(&v)))
	case nil:
		return _hash_int(0)
	}
	panic("type is not comparable!")
}

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type uinteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type number interface {
	integer | uinteger
}

func _hash_int[T number](v T) uint64 {
	var x = uint64(v)
	x = (x ^ (x >> 30)) * uint64(0xbf58476d1ce4e5b9)
	x = (x ^ (x >> 27)) * uint64(0x94d049bb133111eb)
	x = x ^ (x >> 31)
	return x
}

type decimal interface {
	~float32 | ~float64
}

func _hash_float[T decimal](v T) uint64 {
	return _hash_int(uint64(v))
}

func _hash_string(v string) uint64 {
	return datastructures.FastStrHash(v)
}

func _bool_hash(v bool) uint64 {
	if v {
		return 1
	}
	return 0
}
