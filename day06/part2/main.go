package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

const maxDst = 10_000

// const maxDst = 32

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
	count := 0
	for y := range height {
	nextPoint:
		for x := range width {
			pos := math.NewVector2(x, y)
			totalDst := 0
			for _, point := range points {
				totalDst += point.ManhattanDst(pos)
				if totalDst >= maxDst {
					continue nextPoint
				}
			}
			count++
		}
	}

	return count
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
