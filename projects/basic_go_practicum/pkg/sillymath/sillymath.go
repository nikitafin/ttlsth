package sillymath

import "golang.org/x/exp/constraints"

func SliceSum[T constraints.Ordered](input []T) T {
	var accum T
	for _, v := range input {
		accum += v
	}

	return accum
}

func MapSlice[T constraints.Ordered](input []T, op func(T) T) {
	for i := range input {
		input[i] = op(input[i])
	}
}

func FoldSlice[T constraints.Ordered](input []T, op func(T, T) T, init T) T {
	accum := op(init, input[0])
	for _, v := range input[1:] {
		accum = op(accum, v)
	}
	return accum
}
