package algo

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](slice []T) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
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
