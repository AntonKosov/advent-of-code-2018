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
	return value(&tree)
}

func value(tree *[]int) int {
	nodesCount, metadataCount := (*tree)[0], (*tree)[1]
	*tree = (*tree)[2:]

	nodeValues := make(map[int]int, nodesCount)
	for i := range nodesCount {
		nodeValues[i+1] = value(tree)
	}

	sum := 0
	for _, v := range (*tree)[:metadataCount] {
		if nodesCount == 0 {
			sum += v
			continue
		}
		sum += nodeValues[v]
	}

	*tree = (*tree)[metadataCount:]

	return sum
}
