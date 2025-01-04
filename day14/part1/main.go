package main

import (
	"fmt"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() int {
	return transform.StrToInt(input.Lines()[0])
}

func process(offset int) int {
	const targetRecipies = 10
	recipies := []int{3, 7}
	elf1, elf2 := 0, 1
	move := func(elf *int) { *elf = (*elf + 1 + recipies[*elf]) % len(recipies) }
	for len(recipies) < offset+targetRecipies {
		sum := recipies[elf1] + recipies[elf2]
		recipies = append(recipies, math.NumToDigits[int, int](sum)...)
		move(&elf1)
		move(&elf2)
	}

	result := 0
	for i := offset; i < offset+targetRecipies; i++ {
		result = result*10 + recipies[i]
	}

	return result
}
