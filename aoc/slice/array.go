package slice

func New2D[T any](width, height int) [][]T {
	arr := make([][]T, height)
	for i := range arr {
		arr[i] = make([]T, width)
	}

	return arr
}

func Fill2D[T any](arr [][]T, value T) {
	for r, row := range arr {
		for c := range row {
			arr[r][c] = value
		}
	}
}
