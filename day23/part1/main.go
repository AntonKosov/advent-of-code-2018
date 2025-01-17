package main

import (
	"fmt"
	"slices"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []Nanobot {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) Nanobot {
		parts := transform.StrToInts(line)
		return Nanobot{
			position: math.NewVector3(parts[0], parts[1], parts[2]),
			radius:   parts[3],
		}
	})
}

func process(nanobots []Nanobot) int {
	bot := slices.MaxFunc(nanobots, func(a, b Nanobot) int { return a.radius - b.radius })
	count := 0
	for _, bot2 := range nanobots {
		if bot.position.ManhattanDst(bot2.position) <= bot.radius {
			count++
		}
	}

	return count
}

type Nanobot struct {
	position math.Vector3[int]
	radius   int
}
