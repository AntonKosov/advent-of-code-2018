package main

import (
	"fmt"
	"strings"

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

func process(ids []string) string {
	for i := 0; i < len(ids)-1; i++ {
		for j := i + 1; j < len(ids); j++ {
			if common, ok := match(ids[i], ids[j]); ok {
				return common
			}
		}
	}

	panic("not found")
}

func match(id1, id2 string) (common string, ok bool) {
	id2runes := []rune(id2)
	var sb strings.Builder
	for i, r := range id1 {
		if r == id2runes[i] {
			sb.WriteRune(r)
		}
	}

	return sb.String(), sb.Len() == len(id1)-1
}
