package algo

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](slice []T) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				swap(&slice[j], &slice[j+1])
			}
		}
	}
}

func InsertionSort[T constraints.Ordered](slice []T) {
	for i := 1; i < len(slice); i++ {
		el := slice[i]
		j := i - 1
		for ; j != -1 && slice[j] > el; j-- {
			slice[j+1] = slice[j]
		}
		slice[j+1] = el
	}
}

func MergeSort[T constraints.Ordered](slice []T) {
	m := len(slice) / 2

	if len(slice) == 0 {
		return
	}
	if len(slice) == 1 {
		return
	}
	if len(slice) == 2 {
		if slice[0] > slice[1] {
			swap(&slice[0], &slice[1])
		}
		return
	}

	MergeSort(slice[:m])
	MergeSort(slice[m:])

	left := make([]T, m)
	right := make([]T, len(slice)-m)
	left = append(left, slice[:m]...)
	right = append(right, slice[m:]...)

	j, k := 0, 0
	for i := 0; i < len(slice); i++ {
		if left[j] < right[k] {
			slice[i] = left[j]
			j++
		} else {
			slice[i] = right[k]
			k++
		}
	}
}

func QuickSort[T constraints.Ordered](slice []T) {
	if len(slice) == 0 || len(slice) == 1 {
		return
	}
	pivot := slice[len(slice)-1]

	j := 0
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < pivot {
			swap(&slice[i], &slice[j])
			j++
		}
	}

	for i := len(slice) - 2; i >= j; i-- {
		slice[i+1] = slice[i]
	}

	slice[j] = pivot

	QuickSort(slice[:j])
	QuickSort(slice[j+1:])
}
