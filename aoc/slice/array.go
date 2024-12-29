package slice

func New2D[T any](width, height int) [][]T {
	arr := make([][]T, height)
	for i := range arr {
		arr[i] = make([]T, width)
	}

	return arr
}
