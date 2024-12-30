package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() []int {
	return transform.StrToInts(input.Lines()[0])
}

func process(tree []int) int {
	return metadata(&tree)
}

func metadata(tree *[]int) int {
	nodesCount, metadataCount := (*tree)[0], (*tree)[1]
	*tree = (*tree)[2:]

	sum := 0
	for range nodesCount {
		sum += metadata(tree)
	}

	for i := range metadataCount {
		sum += (*tree)[i]
	}

	*tree = (*tree)[metadataCount:]

	return sum
}
