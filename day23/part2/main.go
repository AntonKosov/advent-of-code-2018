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
	edges := make([]Edge, 0, len(nanobots)*2)
	origin := math.NewVector3(0, 0, 0)
	for _, nanobot := range nanobots {
		dst := nanobot.position.ManhattanDst(origin)
		r := nanobot.radius
		edges = append(
			edges,
			Edge{dst: dst - r, depth: 1},
			Edge{dst: dst + r + 1, depth: -1},
		)
	}

	slices.SortFunc(edges, func(a, b Edge) int { return a.dst - b.dst })

	// assuming the target point is positive
	maxDepth, depth, dst := 0, 0, 0
	for _, edge := range edges {
		depth += edge.depth
		if depth > maxDepth {
			maxDepth = depth
			dst = edge.dst
		}
	}

	return dst
}

type Edge struct {
	dst   int
	depth int
}

type Nanobot struct {
	position math.Vector3[int]
	radius   int
}
