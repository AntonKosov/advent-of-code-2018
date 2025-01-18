package main

import (
	"fmt"
	"maps"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/stack"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []math.Vector4[int] {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) math.Vector4[int] {
		nums := transform.StrToInts(line)
		return math.NewVector4(nums[0], nums[1], nums[2], nums[3])
	})
}

func process(points []math.Vector4[int]) int {
	restPoints := make(map[math.Vector4[int]]bool, len(points))
	for _, point := range points {
		restPoints[point] = true
	}

	constellations := 0
	for len(restPoints) > 0 {
		constellations++
		constellation := stack.New[math.Vector4[int]]()
		var startPoint math.Vector4[int]
		maps.Keys(restPoints)(func(v math.Vector4[int]) bool {
			startPoint = v
			return false
		})
		constellation.Push(startPoint)
		delete(restPoints, constellation.Peek())
		for !constellation.Empty() {
			point := constellation.Pop()
			for _, offset := range reachablePositions {
				pos := point.Add(offset)
				if restPoints[pos] {
					delete(restPoints, pos)
					constellation.Push(pos)
				}
			}
		}
	}

	return constellations
}

var reachablePositions []math.Vector4[int]

func init() {
	const distance = 3
	origin := math.NewVector4(0, 0, 0, 0)
	for u0 := -distance; u0 <= distance; u0++ {
		for u1 := -distance; u1 <= distance; u1++ {
			for u2 := -distance; u2 <= distance; u2++ {
				for u3 := -distance; u3 <= distance; u3++ {
					pos := math.NewVector4(u0, u1, u2, u3)
					if pos.ManhattanDst(origin) <= distance {
						reachablePositions = append(reachablePositions, pos)
					}
				}
			}
		}
	}
}
