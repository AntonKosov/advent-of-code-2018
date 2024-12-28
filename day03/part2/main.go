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

func read() []Rectangle {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) Rectangle {
		parts := transform.StrToInts(line)
		lt := math.NewVector2(parts[1], parts[2])
		return Rectangle{
			id:          parts[0],
			leftTop:     lt,
			rightBottom: math.NewVector2(lt.X+parts[3]-1, lt.Y+parts[4]-1),
		}
	})
}

func process(rectangles []Rectangle) int {
	intersects := make([]bool, len(rectangles)+1)
	for i, ri := range rectangles {
		for j := i + 1; j < len(rectangles); j++ {
			rj := rectangles[j]
			if ri.Intersects(rj) {
				intersects[ri.id] = true
				intersects[rj.id] = true
			}
		}

		if !intersects[ri.id] {
			return ri.id
		}
	}

	panic("not found")
}

type Rectangle struct {
	id          int
	leftTop     math.Vector2[int]
	rightBottom math.Vector2[int]
}

func (r1 Rectangle) Intersects(r2 Rectangle) bool {
	minX := max(r1.leftTop.X, r2.leftTop.X)
	maxX := min(r1.rightBottom.X, r2.rightBottom.X)
	minY := max(r1.leftTop.Y, r2.leftTop.Y)
	maxY := min(r1.rightBottom.Y, r2.rightBottom.Y)

	return minX <= maxX && minY <= maxY
}
