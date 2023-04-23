package datastructures

type Ordered interface {
	~int | ~uint | ~string |
		~int8 | ~int16 | ~int32 | ~int64 |
		~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type Hashable[T any] interface {
	Hash() uint64
	Equalizer[T]
}

type Equalizer[T any] interface {
	Equals(T) bool
}

type Comparable[T any] interface {
	Equalizer[T]
	Lt(T) bool
	Gt(T) bool
}
