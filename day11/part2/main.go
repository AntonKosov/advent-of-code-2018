package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

const gridSize = 300

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() int {
	return transform.StrToInt(input.Lines()[0])
}

func process(serialNumber int) string {
	grid := generateGrid(serialNumber)
	cache := slice.New3D[*int](gridSize+1, gridSize+1, gridSize+1)
	maxPowerLevel := 0
	var bestPos math.Vector2[int]
	bestSize := 0
	for r := 1; r <= gridSize; r++ {
		for c := 1; c <= gridSize; c++ {
			for size := 1; r+size-1 <= gridSize && c+size-1 <= gridSize; size++ {
				pos := math.NewVector2(c, r)
				powerLevel := calcPower(grid, pos, size, cache)
				if powerLevel > maxPowerLevel {
					maxPowerLevel = powerLevel
					bestPos = pos
					bestSize = size
				}
			}
		}
	}

	return fmt.Sprintf("%v,%v,%v", bestPos.X, bestPos.Y, bestSize)
}

func calcPower(grid [][]int, pos math.Vector2[int], size int, cache [][][]*int) (sum int) {
	if size == 0 {
		return 0
	}

	if v := cache[pos.Y][pos.X][size]; v != nil {
		return *v
	}

	defer func() { cache[pos.Y][pos.X][size] = &sum }()

	sum = calcPower(grid, pos.Add(math.NewVector2(1, 1)), size-1, cache)
	sum += grid[pos.Y][pos.X]
	for i := 1; i < size; i++ {
		sum += grid[pos.Y+i][pos.X] + grid[pos.Y][pos.X+i]
	}

	return sum
}

func generateGrid(serialNumber int) [][]int {
	grid := slice.New2D[int](gridSize+1, gridSize+1)
	for r := 1; r < gridSize+1; r++ {
		for c := 1; c < gridSize+1; c++ {
			rackID := c + 10
			powerLevel := (rackID*r + serialNumber) * rackID
			powerLevel = (powerLevel%1000)/100 - 5
			grid[r][c] = powerLevel
		}
	}

	return grid
}
