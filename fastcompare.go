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

const (
	c1_64 uint64 = 0x87c37b91114253d5
	c2_64 uint64 = 0x4cf5ad432745937f
)

// Implementation of murmurHash3_64
func UniqueStrHash[T stringType](s T, seed uint32) uint64 {
	data := []byte(s)
	length := len(data)
	h1 := uint64(seed)
	h2 := uint64(seed)
	nBlocks := length / 16

	for i := 0; i < nBlocks; i++ {
		k1 := getblock64(data, i*16+0)
		k2 := getblock64(data, i*16+8)

		k1 *= c1_64
		k1 = (k1 << 31) | (k1 >> (64 - 31))
		k1 *= c2_64
		h1 ^= k1

		h1 = (h1 << 27) | (h1 >> (64 - 27))
		h1 += h2
		h1 = h1*5 + 0x52dce729

		k2 *= c2_64
		k2 = (k2 << 33) | (k2 >> (64 - 33))
		k2 *= c1_64
		h2 ^= k2

		h2 = (h2 << 31) | (h2 >> (64 - 31))
		h2 += h1
		h2 = h2*5 + 0x38495ab5
	}

	tail := data[nBlocks*16:]

	k1 := uint64(0)
	k2 := uint64(0)

	switch length & 15 {
	case 15:
		k2 ^= uint64(tail[14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(tail[13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(tail[12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(tail[11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(tail[10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(tail[9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(tail[8])
		k2 *= c2_64
		k2 = (k2 << 33) | (k2 >> (64 - 33))
		k2 *= c1_64
		h2 ^= k2

		fallthrough
	case 8:
		k1 ^= uint64(tail[7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(tail[6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(tail[5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(tail[4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(tail[3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(tail[2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(tail[0])
		k1 *= c1_64
		k1 = (k1 << 31) | (k1 >> (64 - 31))
		k1 *= c2_64
		h1 ^= k1
	}

	h1 ^= uint64(length)
	h2 ^= uint64(length)

	h1 += h2
	h2 += h1

	h1 = fmix64(h1)
	h2 = fmix64(h2)

	h1 += h2
	// h2 += h1

	return h1
}

func getblock64(data []byte, offset int) uint64 {
	return uint64(data[offset+0]) |
		(uint64(data[offset+1]) << 8) |
		(uint64(data[offset+2]) << 16) |
		(uint64(data[offset+3]) << 24) |
		(uint64(data[offset+4]) << 32) |
		(uint64(data[offset+5]) << 40) |
		(uint64(data[offset+6]) << 48) |
		(uint64(data[offset+7]) << 56)
}

func fmix64(k uint64) uint64 {
	k ^= k >> 33
	k *= 0xff51afd7ed558ccd
	k ^= k >> 33
	k *= 0xc4ceb9fe1a85ec53
	k ^= k >> 33
	return k
}
