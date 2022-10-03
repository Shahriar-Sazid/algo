package algo

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](arr []T, x T) bool {
	s, e := 0, len(arr)-1
	m := (s + e) / 2

	for {
		if arr[m] == x {
			return true
		}
		if s == m && m == e {
			return false
		}
		if x < arr[m] {
			e = m - 1
			m = (s + e) / 2
		}
		if x > arr[m] {
			s = m + 1
			m = (s + e) / 2
		}
	}
}

func BinarySearchWithCheck[T constraints.Ordered](arr []T, x T) (bool, error) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false, errors.New("array is not sorted in ascending order")
		}
	}
	return BinarySearch(arr, x), nil
}
