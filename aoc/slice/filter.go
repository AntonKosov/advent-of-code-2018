package slice

func Filter[T any](items []T, condition func(T) bool) (result []T) {
	for _, item := range items {
		if condition(item) {
			result = append(result, item)
		}
	}

	return
}