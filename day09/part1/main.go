package main

import (
	"fmt"
	"slices"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/list"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (players, lastMarble int) {
	nums := transform.StrToInts(input.Lines()[0])

	return nums[0], nums[1]
}

func process(players, lastMarble int) int {
	scores := make([]int, players)
	var marbles list.Circular[int]
	current := marbles.Add(0)
	for i := 1; i <= lastMarble; i++ {
		if i%23 != 0 {
			current = marbles.InsertAfter(i, current.Next())
			continue
		}
		for j := 0; j < 6; j++ {
			current = current.Prev()
		}
		scores[(i-1)%players] += i + marbles.Remove(current.Prev())
	}

	return slices.Max(scores)
}
