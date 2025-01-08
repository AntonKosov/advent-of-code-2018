package math

func CountBits[T Numbers](value T) int {
	count := 0
	for value != 0 {
		value &= value - 1
		count++
	}

	return count
}
