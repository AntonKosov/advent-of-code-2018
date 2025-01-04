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

func process(targetSequence int) int {
	targetRecipies := math.NumToDigits[int, int](targetSequence)
	recipies := []int{3, 7}
	elf1, elf2 := 0, 1
	move := func(elf *int) { *elf = (*elf + 1 + recipies[*elf]) % len(recipies) }
	for {
		sum := recipies[elf1] + recipies[elf2]
		newRecepies := math.NumToDigits[int, int](sum)
		recipies = append(recipies, newRecepies...)
		if len(recipies) > len(targetRecipies) {
		next:
			for j := 0; j < len(newRecepies); j++ {
				for k, d := range targetRecipies {
					if d != recipies[len(recipies)-len(targetRecipies)+k-j] {
						continue next
					}
				}
				return len(recipies) - len(targetRecipies) - j
			}
		}
		move(&elf1)
		move(&elf2)
	}
}
