package binarysearch

import (
	"math"

	"github.com/Nigel2392/go-datastructures"
)

type lt[T datastructures.Ordered] struct {
	Val T
}

func (lt lt[T]) Lt(val T) bool {
	return lt.Val < val
}

func MakeComparableSlice[T datastructures.Ordered](arr []T) []datastructures.Comparable[T] {
	var newArr = make([]datastructures.Comparable[T], len(arr))
	for i, v := range arr {
		newArr[i] = lt[T]{Val: v}
	}
	return newArr
}

func Search[T datastructures.Comparable[T]](arr []T, val T) int {
	return findTarget(arr, val, 0, len(arr)-1)
}

func findTarget[T datastructures.Comparable[T]](arr []T, val T, start, end int) int {
	if start > end {
		return -1
	}

	var mid = int(math.Floor(float64(start+end) / 2))
	if !arr[mid].Lt(val) && !val.Lt(arr[mid]) {
		return mid
	}
	if val.Lt(arr[mid]) {
		return findTarget(arr, val, start, mid-1)
	}
	if arr[mid].Lt(val) {
		return findTarget(arr, val, mid+1, end)
	}
	return -1
}
