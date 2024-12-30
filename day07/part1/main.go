package main

import (
	"fmt"
	"strings"

	"github.com/AntonKosov/advent-of-code-2018/aoc/heap"
	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (instructions map[rune][]rune) {
	lines := input.Lines()
	lines = lines[:len(lines)-1]
	instructions = make(map[rune][]rune, 26)
	for _, line := range lines {
		first, second := rune(line[5]), rune(line[36])
		instructions[first] = append(instructions[first], second)
	}

	return
}

func process(instructions map[rune][]rune) string {
	var order strings.Builder
	rank := make(map[rune]int, len(instructions))
	for _, steps := range instructions {
		for _, r := range steps {
			rank[r]++
		}
	}

	pq := heap.New(func(r1, r2 rune) bool { return r1 < r2 })
	for step := range instructions {
		if rank[step] == 0 {
			pq.Push(step)
		}
	}

	for pq.Len() > 0 {
		currentStep := pq.Pop()
		order.WriteRune(currentStep)
		steps := instructions[currentStep]
		for _, step := range steps {
			rank[step]--
			if rank[step] == 0 {
				pq.Push(step)
			}
		}
	}

	return order.String()
}
