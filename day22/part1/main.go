package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (depth int, target math.Vector2[int]) {
	lines := input.Lines()
	depth = transform.StrToInts(lines[0])[0]
	targetValues := transform.StrToInts(lines[1])
	target = math.NewVector2(targetValues[0], targetValues[1])

	return
}

func process(depth int, target math.Vector2[int]) int {
	cave := generateCave(depth, target)
	riskLevel := 0
	for _, row := range cave {
		for _, region := range row {
			riskLevel += int(region)
		}
	}

	return riskLevel
}

type Region uint8

const (
	RockyRegion  Region = 0
	WetRegion    Region = 1
	NarrowRegion Region = 2
)

func generateCave(depth int, target math.Vector2[int]) (cave [][]Region) {
	cave = slice.New2D[Region](target.Y+1, target.X+1)
	erosionLevels := slice.New2D[int](target.Y+1, target.X+1)
	for y, row := range cave {
		for x := range row {
			idx := geologicalIndex(math.NewVector2(x, y), target, erosionLevels)
			erosionLevel := (idx + depth) % 20183
			erosionLevels[y][x] = erosionLevel
			cave[y][x] = Region(erosionLevel % 3)
		}
	}

	return
}

func geologicalIndex(pos, target math.Vector2[int], erosionLevels [][]int) int {
	if (pos.X == 0 && pos.Y == 0) || pos == target {
		return 0
	}

	if pos.Y == 0 {
		return pos.X * 16807
	}

	if pos.X == 0 {
		return pos.Y * 48271
	}

	return erosionLevels[pos.Y-1][pos.X] * erosionLevels[pos.Y][pos.X-1]
}
