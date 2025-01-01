package math_test

import (
	"testing"

	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
)

func TestSystemOfLinearEquationsSolver(t *testing.T) {
	matrix := [][]int{
		{1, 3, -2, 5},
		{3, 5, 6, 7},
		{2, 4, 3, 8},
	}
	math.SolveSystemOfLinearEquations(matrix)
	expectedMatrix := [][]int{
		{1, 0, 0, -15},
		{0, 1, 0, 8},
		{0, 0, 1, 2},
	}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] != expectedMatrix[r][c] {
				t.Logf("Unexpected value at [%v][%v]: expected %v, actual %v", r, c, expectedMatrix[r][c], matrix[r][c])
				t.Fail()
			}
		}
	}
}
