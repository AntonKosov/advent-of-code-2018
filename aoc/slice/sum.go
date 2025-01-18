package slice

import "github.com/AntonKosov/advent-of-code-2018/aoc/math"

func Sum[I any, T math.Numbers](items []I, value func(I) T) (sum T) {
	for _, item := range items {
		sum += value(item)
	}

	return
}
