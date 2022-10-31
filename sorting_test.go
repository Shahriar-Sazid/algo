package algo

import (
	"testing"

	. "github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func isSorted[T constraints.Ordered](t *testing.T, s []T, sc []T) {
	Len(t, s, len(sc))
	for _, v := range s {
		Contains(t, sc, v)
	}
	for _, v := range sc {
		Contains(t, s, v)
	}
	for i := 0; i < len(s)-1; i++ {
		LessOrEqual(t, s[i], s[i+1])
	}
}
func TestSorting(t *testing.T) {
	s := []int{1, 3, 42, 22, 3, 3, 1, 2}
	sc := []int{1, 3, 42, 22, 3, 3, 1, 2}
	BubbleSort(s)
	isSorted(t, s, sc)

	ss := []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	ssc := []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	BubbleSort(ss)
	isSorted(t, ss, ssc)

	s = []int{}
	sc = []int{}
	BubbleSort(s)
	isSorted(t, s, sc)

	var n []int
	var nc []int
	BubbleSort(n)
	isSorted(t, n, nc)

	s = []int{1}
	sc = []int{1}
	BubbleSort(s)
	isSorted(t, s, sc)

	s = []int{1, 3, 42, 22, 3, 3, 1, 2}
	sc = []int{1, 3, 42, 22, 3, 3, 1, 2}
	InsertionSort(s)
	isSorted(t, s, sc)

	ss = []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	ssc = []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	InsertionSort(ss)
	isSorted(t, ss, ssc)

	s = []int{}
	sc = []int{}
	InsertionSort(s)
	isSorted(t, s, sc)

	var n1 []int
	var n1c []int
	InsertionSort(n1)
	isSorted(t, n1, n1c)

	s = []int{1}
	sc = []int{1}
	InsertionSort(s)
	isSorted(t, s, sc)

	s = []int{1, 3, 42, 22, 3, 3, 1, 2}
	sc = []int{1, 3, 42, 22, 3, 3, 1, 2}
	MergeSort(s)
	isSorted(t, s, sc)

	ss = []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	ssc = []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	MergeSort(ss)
	isSorted(t, ss, ssc)

	s = []int{}
	sc = []int{}
	MergeSort(s)
	isSorted(t, s, sc)

	var n2 []int
	var n2c []int
	MergeSort(n2)
	isSorted(t, n2, n2c)

	s = []int{1}
	sc = []int{1}
	MergeSort(s)
	isSorted(t, s, sc)

	s = []int{1, 3, 42, 22, 3, 3, 1, 2}
	sc = []int{1, 3, 42, 22, 3, 3, 1, 2}
	QuickSort(s)
	isSorted(t, s, sc)

	s = []int{1, 3, 42, 303, 22, 3, 3, 1, 2}
	sc = []int{1, 3, 42, 303, 22, 3, 3, 1, 2}
	QuickSort(s)
	isSorted(t, s, sc)

	ss = []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	ssc = []string{"af", "sdj", "sdf", "23", "ter", "aa", "b"}
	QuickSort(ss)
	isSorted(t, ss, ssc)

	s = []int{}
	sc = []int{}
	QuickSort(s)
	isSorted(t, s, sc)

	var n3 []int
	var n3c []int
	QuickSort(n3)
	isSorted(t, n3, n3c)

	s = []int{1}
	sc = []int{1}
	QuickSort(s)
	isSorted(t, s, sc)
}
