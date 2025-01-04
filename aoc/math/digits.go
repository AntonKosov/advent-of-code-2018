package math

import "slices"

func CountDigits[T Numbers](num T) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}

	return count
}

func NumToDigits[T, R Numbers](num T) []R {
	if num == 0 {
		return []R{0}
	}

	var reversed []R
	for num != 0 {
		reversed = append(reversed, R(num%10))
		num /= 10
	}

	slices.Reverse(reversed)

	return reversed
}
