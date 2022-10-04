package algo

import "golang.org/x/exp/constraints"

func swap[T constraints.Ordered](a, b *T) {
	temp := *a
	*a = *b
	*b = temp
}
