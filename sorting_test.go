package algo

import (
	"testing"

	. "github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func isSorted[T constraints.Ordered](t *testing.T, s []T) {
	for i := 0; i < len(s)-1; i++ {
		LessOrEqual(t, s[i], s[i+1])
	}
}
func TestSorting(t *testing.T) {
	s := []int{1, 3, 42, 22, 3, 3, 1, 2}
	BubbleSort(s)
	isSorted(t, s)

	ss := []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	BubbleSort(ss)
	isSorted(t, ss)

	s = []int{}
	BubbleSort(s)
	isSorted(t, s)

	var n []int
	BubbleSort(n)
	isSorted(t, n)

	s = []int{1}
	BubbleSort(s)
	isSorted(t, s)

	s = []int{1, 3, 42, 22, 3, 3, 1, 2}
	InsertionSort(s)
	isSorted(t, s)

	ss = []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	InsertionSort(ss)
	isSorted(t, ss)

	s = []int{}
	InsertionSort(s)
	isSorted(t, s)

	var n1 []int
	BubbleSort(n1)
	isSorted(t, n1)

	s = []int{1}
	InsertionSort(s)
	isSorted(t, s)

	s = []int{1, 3, 42, 22, 3, 3, 1, 2}
	MergeSort(s)
	isSorted(t, s)

	ss = []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	MergeSort(ss)
	isSorted(t, ss)

	s = []int{}
	MergeSort(s)
	isSorted(t, s)

	var n2 []int
	BubbleSort(n2)
	isSorted(t, n2)

	s = []int{1}
	MergeSort(s)
	isSorted(t, s)

}
