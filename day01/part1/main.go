package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []int {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	return slice.Map(lines, func(line string) int { return transform.StrToInt(line) })
}

func process(numbers []int) int {
	freq := 0
	for _, num := range numbers {
		freq += num
	}

	return freq
}
