package datastructures

import "unsafe"

// Taken from https://github.com/golang/go/blob/master/src/runtime/map_faststr.go
//
//go:linkname memequal runtime.memequal
func memequal(a, b unsafe.Pointer, n uintptr) bool

// Taken from https://github.com/golang/go/blob/master/src/runtime/map_faststr.go
type _string struct {
	str unsafe.Pointer
	len int
}

type stringType interface {
	~string
}

// A fast string comparison function.
//
// Can compare any type which is a direct or indirect subtype of string.
//
// This function is based on the map_faststr.go implementation in the Go runtime.
//
// Returns true if the strings are equal, false otherwise.
func FastStrCmp[T stringType](a, b T) bool {
	var (
		aStr = (*_string)(unsafe.Pointer(&a))
		bStr = (*_string)(unsafe.Pointer(&b))
	)
	if aStr.len != bStr.len {
		return false
	}
	// if len < 32 a lot of comparisons wont matter, or so the documentation says.
	if aStr.len < 32 {
		return aStr.str == bStr.str || memequal(aStr.str, bStr.str, uintptr(aStr.len))
	}

	// Check if the first 4 bytes are equal
	if *((*[4]byte)(unsafe.Pointer(aStr.str))) != *((*[4]byte)(unsafe.Pointer(bStr.str))) {
		return false
	}
	// Check if the last 4 bytes are equal
	if *((*[4]byte)(unsafe.Pointer(uintptr(aStr.str) + uintptr(aStr.len) - 4))) != *((*[4]byte)(unsafe.Pointer(uintptr(bStr.str) + uintptr(bStr.len) - 4))) {
		return false
	}

	// compare the whole string
	return memequal(aStr.str, bStr.str, uintptr(aStr.len))
}

func FastStrHash[T stringType](s T) uint64 {
	// return murmurHash3_64(s, 0)
	var (
		str = (*_string)(unsafe.Pointer(&s))
		h   = uint64(0)
	)
	for i := 0; i < str.len; i++ {
		h = 63*h + uint64(*(*byte)(unsafe.Pointer(uintptr(str.str) + uintptr(i))))
	}
	return h
}
