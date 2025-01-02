package slice

func New2D[T any](d1, d2 int) [][]T {
	arr := make([][]T, d1)
	for i := range arr {
		arr[i] = make([]T, d2)
	}

	return arr
}

func New3D[T any](d1, d2, d3 int) [][][]T {
	arr := make([][][]T, d1)
	for i := range arr {
		arr[i] = New2D[T](d2, d3)
	}

	return arr
}

func Fill[T any](arr []T, value T) {
	for i := range arr {
		arr[i] = value
	}
}

func Fill2D[T any](arr [][]T, value T) {
	for _, row := range arr {
		Fill(row, value)
	}
}
