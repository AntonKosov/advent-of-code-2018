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
			leftTop:     lt,
			rightBottom: math.NewVector2(lt.X+parts[3]-1, lt.Y+parts[4]-1),
		}
	})
}

func process(rectangles []Rectangle) int {
	intersections := map[math.Vector2[int]]bool{}
	for i := 0; i < len(rectangles)-1; i++ {
		ri := rectangles[i]
		for j := i + 1; j < len(rectangles); j++ {
			intersection, ok := ri.Intersection(rectangles[j])
			if !ok {
				continue
			}
			for r := intersection.leftTop.Y; r <= intersection.rightBottom.Y; r++ {
				for c := intersection.leftTop.X; c <= intersection.rightBottom.X; c++ {
					intersections[math.NewVector2(c, r)] = true
				}
			}
		}
	}

	return len(intersections)
}

type Rectangle struct {
	leftTop     math.Vector2[int]
	rightBottom math.Vector2[int]
}

func (r1 Rectangle) Intersection(r2 Rectangle) (intersection Rectangle, ok bool) {
	minX := max(r1.leftTop.X, r2.leftTop.X)
	maxX := min(r1.rightBottom.X, r2.rightBottom.X)
	minY := max(r1.leftTop.Y, r2.leftTop.Y)
	maxY := min(r1.rightBottom.Y, r2.rightBottom.Y)

	return Rectangle{
		leftTop:     math.NewVector2(minX, minY),
		rightBottom: math.NewVector2(maxX, maxY),
	}, minX <= maxX && minY <= maxY
}
