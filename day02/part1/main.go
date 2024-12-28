package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []string {
	lines := input.Lines()
	return lines[:len(lines)-1]
}

func process(ids []string) int {
	two, three := 0, 0
	for _, id := range ids {
		hasTwo, hasThree := analyze(id)
		if hasTwo {
			two++
		}
		if hasThree {
			three++
		}
	}

	return two * three
}

func analyze(id string) (two, three bool) {
	counts := make(map[rune]int, len(id))
	for _, r := range id {
		counts[r]++
	}

	for _, c := range counts {
		switch c {
		case 2:
			two = true
		case 3:
			three = true
		}
		if two && three {
			break
		}
	}

	return
}
