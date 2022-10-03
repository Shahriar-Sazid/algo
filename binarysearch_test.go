package algo

import (
	"testing"

	. "github.com/stretchr/testify/require"
)

func Test_BinarySearch(t *testing.T) {
	a := []uint{1, 2, 3, 4, 5, 6}

	found := BinarySearch(a, 7)
	False(t, found)

	b := []int{2, 5, 6, 7, 8, 9, 11, 44, 66, 77}

	found = BinarySearch(b, 7)
	True(t, found)

	found = BinarySearch(b, 0)
	False(t, found)

	found = BinarySearch(b, 10)
	False(t, found)

	found, err := BinarySearchWithCheck(b, 10)
	Nil(t, err)
	False(t, found)

	b = []int{2, 5, 6, 7, 8, 9, 11, 44, 88, 77}

	_, err = BinarySearchWithCheck(b, 10)
	NotNil(t, err)

	s := []string{"abc", "def", "ghi", "gri", "ijk"}
	found, err = BinarySearchWithCheck(s, "ijk")
	Nil(t, err)
	True(t, found)

	s = []string{"abc", "def", "aghi", "gri", "ijk"}
	_, err = BinarySearchWithCheck(s, "ijk")
	NotNil(t, err)
}
