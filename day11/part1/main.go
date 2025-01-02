package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

const (
	gridSize   = 300
	squareSize = 3
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() int {
	return transform.StrToInt(input.Lines()[0])
}

func process(serialNumber int) string {
	grid := generateGrid(serialNumber)
	maxPowerLevel := 0
	var bestPos math.Vector2[int]
	for r := 1; r <= gridSize-squareSize+1; r++ {
		for c := 1; c <= gridSize-squareSize+1; c++ {
			pos := math.NewVector2(c, r)
			powerLevel := calcPower(grid, pos)
			if powerLevel > maxPowerLevel {
				maxPowerLevel = powerLevel
				bestPos = pos
			}
		}
	}

	return fmt.Sprintf("%v,%v", bestPos.X, bestPos.Y)
}

func calcPower(grid [][]int, pos math.Vector2[int]) int {
	sum := 0
	for y := pos.Y; y <= pos.Y+squareSize-1; y++ {
		for x := pos.X; x <= pos.X+squareSize-1; x++ {
			sum += grid[y][x]
		}
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
