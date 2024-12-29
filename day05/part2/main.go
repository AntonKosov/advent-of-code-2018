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
	minLength := len(units)
	for skipUnit := 'A'; skipUnit <= 'Z'; skipUnit++ {
		minLength = min(minLength, product(units, skipUnit))
	}

	return minLength
}

func product(units []rune, skip rune) int {
	const diff = 'a' - 'A'
	remaining := stack.New[rune](stack.WithCapacity[rune](len(units)))
	for _, r := range units {
		if r == skip || r == skip+diff {
			continue
		}
		if !remaining.Empty() && math.Abs(remaining.Peek()-r) == diff {
			remaining.Pop()
			continue
		}
		remaining.Push(r)
	}

	return remaining.Size()
}
