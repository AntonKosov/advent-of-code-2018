package math

func SolveSystemOfLinearEquations[T Numbers](matrix [][]T) {
	gaussElimination(matrix)
	backSubstitution(matrix)
}

func gaussElimination[T Numbers](matrix [][]T) {
	vars := len(matrix)
	for varIdx := 0; varIdx < vars-1; varIdx++ {
		validLine(matrix, varIdx)
		for row := varIdx + 1; row < vars; row++ {
			if matrix[row][varIdx] == 0 {
				continue
			}
			for col := varIdx + 1; col <= vars; col++ {
				matrix[row][col] = matrix[row][col]*matrix[varIdx][varIdx] - matrix[varIdx][col]*matrix[row][varIdx]
			}
			matrix[row][varIdx] = 0
		}
	}
}

func validLine[T Numbers](matrix [][]T, varIdx int) {
	if matrix[varIdx][varIdx] != 0 {
		return
	}

	for r := varIdx + 1; r < len(matrix); r++ {
		if matrix[r][varIdx] != 0 {
			matrix[r], matrix[varIdx] = matrix[varIdx], matrix[r]
			return
		}
	}

	panic("multiple solution")
}

func backSubstitution[T Numbers](matrix [][]T) {
	vars := len(matrix)
	for i := vars - 1; i >= 0; i-- {
		for j := i + 1; j < vars; j++ {
			matrix[i][vars] -= matrix[i][j] * matrix[j][vars]
			matrix[i][j] = 0
		}
		matrix[i][vars] = matrix[i][vars] / matrix[i][i]
		matrix[i][i] = 1
	}
}
