package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/stack"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []rune {
	return []rune(input.Lines()[0])
}

func process(units []rune) int {
	const diff = 'a' - 'A'
	remaining := stack.New[rune](stack.WithCapacity[rune](len(units)))
	for _, r := range units {
		if !remaining.Empty() && math.Abs(remaining.Peek()-r) == diff {
			remaining.Pop()
			continue
		}
		remaining.Push(r)
	}

	return remaining.Size()
}
