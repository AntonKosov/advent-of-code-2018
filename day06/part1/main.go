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

func read() []math.Vector2[int] {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) math.Vector2[int] {
		parts := transform.StrToInts(line)
		return math.NewVector2(parts[0], parts[1])
	})
}

func process(points []math.Vector2[int]) int {
	width, height := moveToOrigin(points)
	closest := slice.New2D[int](height, width)
	findClosest(closest, points)
	infinite := findInfinite(closest)
	largestArea := 0
	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			if v := closest[y][x]; v < 0 || infinite[v] {
				continue
			}
			largestArea = max(largestArea, measureArea(closest, math.NewVector2(x, y)))
		}
	}

	return largestArea
}

func measureArea(closest [][]int, pos math.Vector2[int]) int {
	idx := closest[pos.Y][pos.X]
	if idx < 0 {
		return 0
	}

	closest[pos.Y][pos.X] = -1
	count := 1
	for _, dir := range dirs {
		nextPos := pos.Add(dir)
		if closest[nextPos.Y][nextPos.X] == idx {
			count += measureArea(closest, nextPos)
		}
	}

	return count
}

func findInfinite(closest [][]int) map[int]bool {
	height, width := len(closest), len(closest[0])
	infinite := map[int]bool{}
	for r := range height {
		infinite[closest[r][0]] = true
		infinite[closest[r][width-1]] = true
	}

	for c := range width {
		infinite[closest[0][c]] = true
		infinite[closest[height-1][c]] = true
	}

	return infinite
}

func findClosest(closest [][]int, points []math.Vector2[int]) {
	for r, row := range closest {
		for c := range row {
			pos := math.NewVector2(c, r)
			bestIdx := 0
			count := 1
			for i := 1; i < len(points); i++ {
				bestDst := points[bestIdx].ManhattanDst(pos)
				dst := points[i].ManhattanDst(pos)
				if dst < bestDst {
					bestIdx = i
					count = 1
					continue
				}
				if dst == bestDst {
					count++
				}
			}

			if count > 1 {
				bestIdx = -1
			}

			closest[r][c] = bestIdx
		}
	}
}

func moveToOrigin(points []math.Vector2[int]) (width, height int) {
	leftTop, rightBottom := points[0], points[0]
	for _, point := range points {
		leftTop.X = min(leftTop.X, point.X)
		leftTop.Y = min(leftTop.Y, point.Y)
		rightBottom.X = max(rightBottom.X, point.X)
		rightBottom.Y = max(rightBottom.Y, point.Y)
	}

	for i := range points {
		points[i] = points[i].Sub(leftTop)
	}

	return rightBottom.X - leftTop.X + 1, rightBottom.Y - leftTop.Y + 1
}

var dirs = []math.Vector2[int]{{X: 1}, {X: -1}, {Y: 1}, {Y: -1}}
