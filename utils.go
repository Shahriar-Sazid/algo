package algo

import "golang.org/x/exp/constraints"

func swap[T constraints.Ordered](a, b *T) {
	temp := *a
	*a = *b
	*b = temp
}

func min[T constraints.Ordered](a T, b ...T) T {
	x := a
	for _, v := range b {
		if x > v {
			x = v
		}
	}
	return x
}

func max[T constraints.Ordered](a T, b ...T) T {
	x := a
	for _, v := range b {
		if x < v {
			x = v
		}
	}
	return x
}
