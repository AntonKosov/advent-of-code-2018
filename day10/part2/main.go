package main

import (
	"fmt"
	smath "math"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []Point {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) Point {
		parts := transform.StrToInts(line)
		return Point{
			position: math.NewVector2(parts[0], parts[1]),
			velocity: math.NewVector2(parts[2], parts[3]),
		}
	})
}

func process(points []Point) int {
	for size, steps := smath.MaxInt, 0; ; steps++ {
		move(points, 1)
		size2 := dst(points)
		if size2 > size {
			return steps
		}
		size = size2
	}
}

func dst(points []Point) int {
	leftTop := points[0].position
	bottomRight := leftTop
	for _, point := range points {
		pos := point.position
		leftTop.X = min(leftTop.X, pos.X)
		leftTop.Y = min(leftTop.Y, pos.Y)
		bottomRight.X = max(bottomRight.X, pos.X)
		bottomRight.Y = max(bottomRight.Y, pos.Y)
	}

	return leftTop.ManhattanDst(bottomRight)
}

func move(points []Point, steps int) {
	for i := range points {
		points[i].Move(steps)
	}
}

type Point struct {
	position, velocity math.Vector2[int]
}

func (p *Point) Move(steps int) {
	p.position = p.position.Add(p.velocity.Mul(steps))
}
